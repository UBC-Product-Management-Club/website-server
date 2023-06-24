package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var routeTimeout map[string]time.Duration

const _DefaultTimeout = 1000 * time.Millisecond

// Middleware for tracking timeout for each request
// The time out can vary depending on the route/request (specified initially with timeoutMap)
func TimeoutMiddleware(timeoutMap map[string]time.Duration) gin.HandlerFunc {
	routeTimeout = timeoutMap
	return timeoutTracker
}

// Track the timeout for each request
func timeoutTracker(c *gin.Context) {
	newRequestCtx, cancel := configureNewContext(c)
	defer cancel()
	cw, tw := configureNewWriter(c)

	finishChan := make(chan struct{})
	panicChan := make(chan interface{})
	go func() {
		defer close(finishChan)
		defer close(panicChan)
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		c.Next()
		finishChan <- struct{}{}
	}()

	select {
	case <-finishChan:
		c.Next()
	case <-newRequestCtx.Done():
		c.AbortWithError(http.StatusRequestTimeout, fmt.Errorf("%s Request Timeout", c.Request.Method))
		tw.Expire()
	case p := <-panicChan:
		c.Writer = cw
		panic(p)
	}
}

// Configure a new context with a timeout
// A context.CancelFunc is returned to allow caller to release the resources associated with context timeout
func configureNewContext(c *gin.Context) (context.Context, context.CancelFunc) {
	timeout, ok := routeTimeout[c.FullPath()]
	if !ok {
		timeout = _DefaultTimeout
	}
	newRequestCtx, cancel := context.WithTimeout(c.Request.Context(), timeout)
	c.Request = c.Request.WithContext(newRequestCtx)
	return newRequestCtx, cancel
}

// Configure new writer that will only write within the valid time
func configureNewWriter(c *gin.Context) (gin.ResponseWriter, *timeoutWriter) {
	oldWriter := c.Writer
	newWriter := newTimeoutWriter(oldWriter)
	c.Writer = newWriter
	return oldWriter, newWriter
}
