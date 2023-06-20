package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// A middleware for handling errors on the server
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return errorHandler
}

// Handling errors for every request
func errorHandler(c *gin.Context) {
	defer handleFatalError(c)
	c.Next()
	c.JSON(-1, c.Errors)
}

// Handle fatal errors (panics)
func handleFatalError(c *gin.Context) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			c.AbortWithError(http.StatusInternalServerError, err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected fatal error has occured"})
		}
	}
}
