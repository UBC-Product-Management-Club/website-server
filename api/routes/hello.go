package routes

import (
	"log"
	"net/http"

	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/UBC-Product-Management-Club/website-server/models"
	"github.com/gin-gonic/gin"
)

func getGreetings(c *gin.Context) {
	greetings, err := database.GetAllGreetings()
	if err != nil {
		log.Fatal(err.Error())
	}
	c.JSON(http.StatusOK, greetings)
}

func postGreetings(c *gin.Context) {
	var newGreeting models.Greeting

	if err := c.BindJSON(&newGreeting); err != nil {
		log.Fatal(err.Error())
		return
	}

	if err := database.AddGreeting(newGreeting); err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(http.StatusCreated, newGreeting)
}
