package main

import (
	"log"

	"gin-basics/internal/inits"
	"gin-basics/internal/router"
)

func init() {
	inits.LoadDotEnv()
	inits.ConnectDatabase()
}

func main() {
	r := router.SetupRouter()

	err := r.Run("0.0.0.0:8080")

	if err != nil {
		log.Fatal("\n-> Error: Running Gin")
		log.Fatal(err)
	}
}
