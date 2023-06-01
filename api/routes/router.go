package routes

import (
	"github.com/UBC-Product-Management-Club/website-server/api/middlewares"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter(r *gin.Engine) {
	if router == nil {
		router = r
	}

	initRoutes()
	initMiddleware()
}

func initRoutes() {
	router.GET("/hello", getGreetings)
	router.POST("/hello", postGreetings)
}

func initMiddleware() {
	router.Use(middlewares.CorsMiddleware())
}
