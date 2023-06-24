package routes

import (
	"net/http"

	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/UBC-Product-Management-Club/website-server/models"
	"github.com/gin-gonic/gin"
)

// GET request for greetings
func getGreetings(c *gin.Context) {
	greetings, err := database.GetAllGreetings()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, greetings)
}

// POST request for greetings
func postGreetings(c *gin.Context) {
	var newGreeting models.Greeting

	if err := c.BindJSON(&newGreeting); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := database.AddGreeting(newGreeting); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newGreeting)
}
