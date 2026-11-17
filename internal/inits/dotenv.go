package inits

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("No .env file found, using environment variables")
		log.Println(err)
	}
}
