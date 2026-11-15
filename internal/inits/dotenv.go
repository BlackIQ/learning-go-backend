package inits

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("\n-> Error: Loading dotenv")
		log.Fatal(err)
	}
}
