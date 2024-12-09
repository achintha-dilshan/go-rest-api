package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	ServerPort    string

	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

// Init initializes the configuration by reading from environment variables.
func Init() *Config {
	// load env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		ServerPort:    os.Getenv("SERVER_PORT"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBName:        os.Getenv("DB_NAME"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
	}
}

// Global configuration instance
var Env = Init()
