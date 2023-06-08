package routes

import (
	"log"
	"net/http"

	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/UBC-Product-Management-Club/website-server/models"
	"github.com/gin-gonic/gin"
)

// GET request for fellow
func getFellow(c *gin.Context) {
	fellows, err := database.GetAllFellows()
	if err != nil {
		log.Fatal(err.Error())
	}
	c.JSON(http.StatusOK, fellows)
}

// POST request for fellow
func postFellow(c *gin.Context) {
	var newFellow models.Fellow

	if err := c.BindJSON(&newFellow); err != nil {
		log.Fatal(err.Error())
		return
	}

	if err := database.AddFellow(newFellow); err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(http.StatusCreated, newFellow)
}
