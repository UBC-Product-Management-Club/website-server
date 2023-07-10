package routes

import (
	"time"

	"github.com/UBC-Product-Management-Club/website-server/api/middlewares"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var routeTimeout map[string]time.Duration

// Initialize the router
func InitRouter(r *gin.Engine) {
	if router == nil {
		router = r
	}
	routeTimeout = map[string]time.Duration{
		"/hello":     		 		(500 * time.Millisecond),
		// "/executive":				(1000 * time.Millisecond),
		"/executive/:department": 	(1000 * time.Millisecond),
		"/fellow":    		 		(1000 * time.Millisecond),
		"/project":   		 		(1500 * time.Millisecond),
	}

	initMiddleware()
	initRoutes()
}

// Set up end points
func initRoutes() {
	setGetPostRoutes("/hello", getGreetings, postGreetings)
	setGetPostRoutes("/executive", getExecutive, postExecutive)
	setGetPostRoutes("/executive/:department", getDepartmentExecutive, postExecutive)
	setGetPostRoutes("/fellow", getFellow, postFellow)
	setGetPostRoutes("/project", getProject, postProject)
}

func setGetPostRoutes(routeName string, getFunc gin.HandlerFunc, postFunc gin.HandlerFunc) {
	router.GET(routeName, getFunc)
	router.POST(routeName, postFunc)
}

// Initialize middleware
func initMiddleware() {
	router.Use(middlewares.ErrorHandlerMiddleware())
	router.Use(middlewares.CorsMiddleware())

	router.SetTrustedProxies(nil)
	router.Use(middlewares.ProxyMiddleware())
	router.Use(helmet.Default())

	router.Use(middlewares.TimeoutMiddleware(routeTimeout))
}
