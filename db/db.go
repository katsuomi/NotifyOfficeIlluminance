package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/katsuomi/NotifyOfficeIlluminance/models"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open("postgres", "host=db port=5432 user=NotifyOfficeIlluminance dbname=NotifyOfficeIlluminance password=NotifyOfficeIlluminance sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
	illuminance := models.Illuminance{
		ID:    1,
		Illuminance:  2430,
		CreatedAt: "2019-11-26T12:08:19+09:00"
	}
	db.Create(&illuminance)
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&models.Illuminance{})
}
