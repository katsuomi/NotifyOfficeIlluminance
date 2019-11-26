package controllers

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/katsuomi/NotifyOfficeIlluminance/models/repository"
)

// Controller is user controlller
type IlluminanceController struct{}

// Index action: GET /illuminances
func (pc IlluminanceController) Index(c *gin.Context) {
	var i repository.IlluminanceRepository
	p, err := i.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /illuminances
func (pc IlluminanceController) Create(c *gin.Context) {
	fmt.Println(c)
	var i repository.IlluminanceRepository
	p, err := i.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, p)
	}
}
