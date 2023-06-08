package routes

import (
	"log"
	"net/http"

	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/UBC-Product-Management-Club/website-server/models"
	"github.com/gin-gonic/gin"
)

func getFellows(c *gin.Context) {
	fellows, err := database.GetAllFellows()
	if err != nil {
		log.Fatal(err.Error())
	}
	c.JSON(http.StatusOK, fellows)
}

func postFellows(c *gin.Context) {
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
