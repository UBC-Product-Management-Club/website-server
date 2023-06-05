package main

import (
	routes "github.com/UBC-Product-Management-Club/website-server/api/routes"
	database "github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRouter(router)
	database.InitDatabase()

	router.Run(":8080")
}
