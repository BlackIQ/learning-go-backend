package inits

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gin-basics/internal/models"
)

func ConnectDatabase() *gorm.DB {
	dbPath := os.Getenv("DB_PATH")

	if dbPath == "" {
		dbPath = "go.db"
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Todo{}, &models.Note{})

	return db
}
