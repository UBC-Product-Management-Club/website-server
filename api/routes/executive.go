package routes

import (
	"net/http"

	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/UBC-Product-Management-Club/website-server/models"
	"github.com/gin-gonic/gin"
)

// GET request for executive
func getExecutive(c *gin.Context) {
	execs, err := database.GetAllExecutives()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, execs)
}

// POST request for executive
func postExecutive(c *gin.Context) {
	var newExec models.Executive

	if err := c.BindJSON(&newExec); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := database.AddExecutive(newExec); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newExec)
}
