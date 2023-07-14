package routes

import (
	"net/http"

	"github.com/UBC-Product-Management-Club/website-server/database"
	"github.com/UBC-Product-Management-Club/website-server/models"
	"github.com/UBC-Product-Management-Club/website-server/utils"
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

// GET request for executive
func getDepartmentExecutive(c *gin.Context) {
	label := c.Param("department")
	execs, err := database.GetAllDepartmentExecutives(label)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, execs)
}

// POST request for executive
func postExecutive(c *gin.Context) {
	var newExec models.Executive

	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	file := form.File["img"][0]
	name := form.Value["name"][0]
	title := form.Value["title"][0]
	department := form.Value["department"][0]

	blobFile, err := file.Open()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	img, err := utils.UploadFile(blobFile, "team_pictures", file.Filename)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	newExec = models.Executive{
		Img:        img,
		Name:       name,
		Title:      title,
		Department: department,
	}

	if err := database.AddExecutive(newExec); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newExec)
}
