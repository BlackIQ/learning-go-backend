package main

import (
	"fmt"

	"github.com/BlackIQ/gin-todo/internal/router"
)

func main() {
	r := router.SetupRouter()

	err := r.Run("0.0.0.0:8080")

	if err != nil {
		fmt.Println("\n-> Error running the app.")
		fmt.Print(err)
	}
}
