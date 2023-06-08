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

	initMiddleware()
	initRoutes()
}

func initRoutes() {
	setGetPostRoutes("/hello", getGreetings, postGreetings)
	setGetPostRoutes("/executive", getExecutive, postExecutive)
	setGetPostRoutes("/fellow", getFellow, postFellow)
	setGetPostRoutes("/project", getProject, postProject)
}

func setGetPostRoutes(routeName string, getFunc gin.HandlerFunc, postFunc gin.HandlerFunc) {
	router.GET(routeName, getFunc)
	router.POST(routeName, postFunc)
}

func initMiddleware() {
	router.Use(middlewares.CorsMiddleware())
}
