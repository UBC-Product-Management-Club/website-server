package main

import (
	"fmt"
	"os"
	"time"

	routes "github.com/UBC-Product-Management-Club/website-server/api/routes"
	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/UBC-Product-Management-Club/website-server/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[%s] Cannot start the server\n", time.Now().Format(time.RFC3339))
			os.Exit(1)
		}
	}()

	router := gin.Default()
	routes.InitRouter(router)
	database.InitDatabase()
	utils.InitUtils()
	router.Run(":8080")
}
