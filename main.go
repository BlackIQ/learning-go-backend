package main

import (
	"fmt"
	"log"
	"os"

	"gin-basics/internal/inits"
	"gin-basics/internal/router"
)

func init() {
	inits.LoadDotEnv()
}

func main() {
	db := inits.ConnectDatabase()

	r := router.SetupRouter(db)

	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

	addr := fmt.Sprintf("0.0.0.0:%v", port)

	err := r.Run(addr)

	if err != nil {
		log.Fatal("\n-> Error: Running Gin")
		log.Fatal(err)
	}
}
