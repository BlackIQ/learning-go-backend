package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"gin-basics/internal/inits"
	"gin-basics/internal/router"
)

//go:embed all:frontend/dist
var distFS embed.FS

func init() {
	inits.LoadDotEnv()
	inits.ConnectDatabase()
}

func main() {
	r := router.SetupRouter(&distFS)

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
