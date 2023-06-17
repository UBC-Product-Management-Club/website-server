package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Configure trusted proxy
func ProxyMiddleware() gin.HandlerFunc {
	return verifyProxy
}

// Verify the proxy of the upcoming request
func verifyProxy(c *gin.Context) {
	hostname := "load-balancer-dns.com"
	if isProxy(c) && c.GetHeader("Host") != hostname {
		c.AbortWithStatus(http.StatusForbidden)
	}
	c.Next()
}

// Check if the request comes from a proxy
func isProxy(c *gin.Context) bool {
	return c.GetHeader("X-Forwarded-For") != "" &&
		c.GetHeader("X-Forwarded-Proto") != "" &&
		c.GetHeader("X-Forwarded-Host") != ""
}
