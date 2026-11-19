package main

import (
	"embed"
	"fmt"
	"log"

	"gin-basics/internal/config"
	"gin-basics/internal/inits"
	"gin-basics/internal/router"
)

//go:embed all:frontend/dist
var distFS embed.FS

func init() {
	inits.LoadEnv()
	inits.ConnectDatabase()
}

func main() {
	cfg := config.GetConfig()

	r := router.SetupRouter(&distFS)

	addr := fmt.Sprintf("0.0.0.0:%v", cfg.Port)

	err := r.Run(addr)

	if err != nil {
		log.Fatal("\n-> Error: Running Gin")
		log.Fatal(err)
	}
}
