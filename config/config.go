package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string

	ServerHost string
	ServerPort string

	DBDriver   string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	AccessTokenKey  string
	RefreshTokenKey string
}

// Init initializes the configuration by reading from environment variables.
func Init() *Config {
	// load env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		AppEnv:          os.Getenv("APP_ENV"),
		ServerHost:      os.Getenv("SERVER_HOST"),
		ServerPort:      os.Getenv("SERVER_PORT"),
		DBDriver:        os.Getenv("DB_DRIVER"),
		DBHost:          os.Getenv("DB_HOST"),
		DBPort:          os.Getenv("DB_PORT"),
		DBName:          os.Getenv("DB_NAME"),
		DBUser:          os.Getenv("DB_USER"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		AccessTokenKey:  os.Getenv("ACCESS_TOKEN_KEY"),
		RefreshTokenKey: os.Getenv("REFRESH_TOKEN_KEY"),
	}
}

// Global configuration instance
var Env = Init()
