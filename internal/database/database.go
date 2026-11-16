package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gin-basics/internal/models"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	dbPath := os.Getenv("DB_PATH")

	if dbPath == "" {
		dbPath = "go.db"
	}

	connection, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	connection.AutoMigrate(
		&models.Todo{},
	)

	db = connection

	return nil
}

func GetDB() *gorm.DB {
	return db
}
