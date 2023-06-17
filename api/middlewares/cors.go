package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Return configuration for CORS
func CorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	config.AllowOriginFunc = checkOrigin
	return cors.New(config)
}

func checkOrigin(org string) bool {
	localAddr := map[string]bool{"http://localhost:3000": true}
	remoteAddr := map[string]bool{"https://www.ubcpm.club/": true}

	return localAddr[org] || remoteAddr[org]
}
