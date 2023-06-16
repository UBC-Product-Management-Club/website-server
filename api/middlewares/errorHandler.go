package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var logger *log.Logger

// A middleware for handling errors on the server
func ErrorHandlerMiddleware() gin.HandlerFunc {
	logger = log.New(os.Stdout, "", log.Flags()&^(log.Ldate|log.Ltime))
	return errorHandler
}

// Handling errors for every request
func errorHandler(c *gin.Context) {
	defer handleFatalError(c)
	c.Next()
	for _, err := range c.Errors {
		logger.Println(err.Error())
	}
}

// Handle fatal errors (panics)
func handleFatalError(c *gin.Context) {
	if err := recover(); err != nil {
		logger.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Unexpected fatal error has occured"})
	}
}
