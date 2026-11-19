package inits

import (
	"gin-basics/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(err)
	}

	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}
}
