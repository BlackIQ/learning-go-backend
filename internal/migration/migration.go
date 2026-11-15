package main

import (
	"gin-basics/internal/inits"
	"gin-basics/internal/models"
)

func init() {
	inits.LoadDotEnv()
	inits.ConnectDatabase()
}

func main() {
	inits.DB.AutoMigrate(
		&models.Todo{},
		&models.Note{},
	)
}
