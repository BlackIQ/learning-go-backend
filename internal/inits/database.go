package inits

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gin-basics/internal/models"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("go.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Todo{}, &models.Note{})

	return db
}
