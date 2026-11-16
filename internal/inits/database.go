package inits

import (
	"gin-basics/internal/database"
)

func ConnectDatabase() {
	database.InitDB()
}
