package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/katsuomi/NotifyOfficeIlluminance/db"
	"github.com/katsuomi/NotifyOfficeIlluminance/models"
)

// Service procides illuminance's behavior
type IlluminanceRepository struct{}

// Illuminance is alias of entity.Illuminance struct
type Illuminance models.Illuminance

// GetAll is get all Illuminance
func (_ IlluminanceRepository) GetAll() ([]Illuminance, error) {
	db := db.GetDB()
	var i []Illuminance
	if err := db.Table("illuminances").Select("id, illuminance").Scan(&i).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// CreateModel is create Illuminance model
func (_ IlluminanceRepository) CreateModel(c *gin.Context) (Illuminance, error) {
	db := db.GetDB()
	var i Illuminance
	if err := c.BindJSON(&i); err != nil {
		return i, err
	}
	if err := db.Create(&i).Error; err != nil {
		return i, err
	}
	return i, nil
}
