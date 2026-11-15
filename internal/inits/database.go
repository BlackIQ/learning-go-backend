package inits

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	DB, err = gorm.Open(sqlite.Open("go.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database")
	}
}
