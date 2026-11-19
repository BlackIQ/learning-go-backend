package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gin-basics/internal/config"
	"gin-basics/internal/models"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	cfg := config.GetConfig()

	connection, err := gorm.Open(sqlite.Open(cfg.DatabasePath), &gorm.Config{})

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
