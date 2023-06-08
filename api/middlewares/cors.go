package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Return configuration for CORS
func CorsMiddleware() gin.HandlerFunc {
	// allow all origins. can be configured later
	return cors.Default()
}
