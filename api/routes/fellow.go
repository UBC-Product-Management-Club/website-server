package routes

import (
	"net/http"

	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/UBC-Product-Management-Club/website-server/models"
	"github.com/gin-gonic/gin"
)

// GET request for fellow
func getFellow(c *gin.Context) {
	fellows, err := database.GetAllFellows()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, fellows)
}

// POST request for fellow
func postFellow(c *gin.Context) {
	var newFellow models.Fellow

	if err := c.BindJSON(&newFellow); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := database.AddFellow(newFellow); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newFellow)
}
