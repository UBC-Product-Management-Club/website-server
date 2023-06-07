package main

import (
	routes "github.com/UBC-Product-Management-Club/website-server/api/routes"
	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRouter(router)
	database.InitDatabase()
	defer database.CloseDatabase()
	router.Run(":8080")
}
