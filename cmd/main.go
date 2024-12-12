package main

import (
	"log"
	"net/http"
	"os"

	"github.com/achintha-dilshan/go-rest-api/config"
	"github.com/achintha-dilshan/go-rest-api/database"
	"github.com/achintha-dilshan/go-rest-api/internal/routes"
)

func main() {
	// init database
	database.Init()
	defer database.Close()

	// run migrations
	if len(os.Args) > 1 {

		// Path to your migration files
		migrationPath := "./database/migrations"

		// Run the migration
		database.RunMigration(database.DB, migrationPath)
		return
	}

	// init router
	router := routes.Init()

	// init server
	port := ":" + config.Env.ServerPort
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Printf("Server is running on port %v", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
