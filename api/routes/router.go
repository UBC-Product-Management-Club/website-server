package routes

import (
	"github.com/UBC-Product-Management-Club/website-server/api/middlewares"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// Initialize the router
func InitRouter(r *gin.Engine) {
	if router == nil {
		router = r
	}

	initMiddleware()
	initRoutes()
}

// Set up end points
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

// Initialize middleware
func initMiddleware() {
	router.Use(middlewares.CorsMiddleware())

	router.SetTrustedProxies(nil)
	router.Use(middlewares.ProxyMiddleware())
	router.Use(helmet.Default())
}
