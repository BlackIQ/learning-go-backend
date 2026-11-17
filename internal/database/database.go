package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gin-basics/internal/models"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	// dbPath := os.Getenv("DB_PATH")
	dsn := "host=10.1.10.120 user=postgres password=9e866eeed379fac0 dbname=todogo port=5432 sslmode=disable"

	// if dbPath == "" {
	// 	dbPath = "go.db"
	// }

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
