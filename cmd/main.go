package main

import (
	"log"

	"github.com/achintha-dilshan/go-rest-api/cmd/api"
	"github.com/achintha-dilshan/go-rest-api/config"
	"github.com/achintha-dilshan/go-rest-api/database"
)

func main() {
	// init database
	db := database.NewDatabase()
	if err := db.Connect(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// Get the database instance
	sqlDB, err := db.GetDB()
	if err != nil {
		log.Fatalf("Error retrieving database instance: %v", err)
	}

	// init server
	port := ":" + config.Env.ServerPort
	server := api.NewAPIServer(port, sqlDB)

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
