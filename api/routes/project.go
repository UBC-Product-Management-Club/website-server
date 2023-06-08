package routes

import (
	"log"
	"net/http"

	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/UBC-Product-Management-Club/website-server/models"
	"github.com/gin-gonic/gin"
)

// GET request for project
func getProject(c *gin.Context) {
	projects, err := database.GetAllProjects()
	if err != nil {
		log.Fatal(err.Error())
	}
	c.JSON(http.StatusOK, projects)
}

// POST request for project
func postProject(c *gin.Context) {
	var newProject models.Project

	if err := c.BindJSON(&newProject); err != nil {
		log.Fatal(err.Error())
		return
	}

	if err := database.AddProject(newProject); err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(http.StatusCreated, newProject)
}
