package inits

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gin-basics/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	DB, err = gorm.Open(sqlite.Open("go.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(
		&models.Todo{},
		&models.Note{},
	)
}
