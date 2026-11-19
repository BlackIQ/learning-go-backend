package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port         int
	Secret       string
	DatabasePath string
}

var config *Config

func LoadConfig() error {
	// APP_PORT
	APP_PORT := os.Getenv("APP_PORT")
	if APP_PORT == "" {
		return fmt.Errorf("[Loading Env] APP_PORT is not set")
	}
	port, err := strconv.Atoi(APP_PORT)
	if err != nil {
		return fmt.Errorf("[Loading Env] %v", err.Error())
	}
	if port <= 0 || port > 65535 {
		return fmt.Errorf("[Loading Env] APP_PORT should be between 0 and 65535")
	}

	// JWT_SECRET
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return fmt.Errorf("[Loading Env] JWT_SECRET is not set")
	}
	if len(secret) < 32 {
		return fmt.Errorf("[Loading Env] JWT_SECRET must be at least 32 bytes")
	}

	// DATABASE_PATH
	databasePath := os.Getenv("DATABASE_PATH")
	if databasePath == "" {
		return fmt.Errorf("[Loading Env] DATABASE_PATH is not set")
	}

	config = &Config{
		Port:         port,
		DatabasePath: databasePath,
		Secret:       secret,
	}

	return nil
}

func GetConfig() *Config {
	if config == nil {
		panic("config is not initialized")
	}

	return config
}
