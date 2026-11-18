package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gin-basics/internal/models"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	mode := "development"

	dsn := "host=10.1.10.120 user=postgres password=9e866eeed379fac0 dbname=todogo port=5432 sslmode=disable"
	dbPath := "go.db"

	var dialector gorm.Dialector

	if mode == "development" {
		dialector = sqlite.Open(dbPath)
	} else {
		dialector = postgres.Open(dsn)
	}

	connection, err := gorm.Open(dialector, &gorm.Config{})

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
