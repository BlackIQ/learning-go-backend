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
	inits.ConnectDatabase()
}

func main() {
	r := router.SetupRouter()

	port := os.Getenv("PORT")
	addr := fmt.Sprintf("0.0.0.0:%v", port)

	err := r.Run(addr)

	if err != nil {
		log.Fatal("\n-> Error: Running Gin")
		log.Fatal(err)
	}
}
