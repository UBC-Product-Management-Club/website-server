package main

import (
	"log"

	routes "github.com/UBC-Product-Management-Club/website-server/api/routes"
	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRouter(router)
	if err := database.InitDatabase(); err != nil {
		log.Fatalln(err)
	}
	defer database.CloseDatabase()
	router.Run(":8080")
}
