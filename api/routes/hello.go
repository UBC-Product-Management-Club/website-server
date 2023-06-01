package routes

import (
	"net/http"

	"github.com/UBC-Product-Management-Club/website-server/models"
	"github.com/gin-gonic/gin"
)

var greetings = []models.Greeting{
	{Message: "Hello World"},
	{Message: "Call To Earthlings"},
	{Message: "Knock Knck"},
}

func getGreetings(c *gin.Context) {
	c.JSON(http.StatusOK, greetings)
}

func postGreetings(c *gin.Context) {
	var newGreeting models.Greeting

	if err := c.BindJSON(&newGreeting); err != nil {
		return
	}

	greetings = append(greetings, newGreeting)
	c.JSON(http.StatusCreated, newGreeting)
}
