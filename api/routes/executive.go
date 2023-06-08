package routes

import (
	"log"
	"net/http"

	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/UBC-Product-Management-Club/website-server/models"
	"github.com/gin-gonic/gin"
)

// GET request for executive
func getExecutive(c *gin.Context) {
	execs, err := database.GetAllExecutives()
	if err != nil {
		log.Fatal(err.Error())
	}
	c.JSON(http.StatusOK, execs)
}

// POST request for executive
func postExecutive(c *gin.Context) {
	var newExec models.Executive

	if err := c.BindJSON(&newExec); err != nil {
		log.Fatal(err.Error())
		return
	}

	if err := database.AddExecutive(newExec); err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(http.StatusCreated, newExec)
}
