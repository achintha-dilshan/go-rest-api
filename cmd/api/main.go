package main

import (
	"log"
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/config"
	"github.com/achintha-dilshan/go-rest-api/database"
	"github.com/achintha-dilshan/go-rest-api/internal/routes"
)

func main() {
	// init database
	database.Init()

	defer func() {
		if err := database.DB.Close(); err != nil {
			log.Printf("Failed to close the database: %v", err)
		} else {
			log.Println("Database connection closed")
		}
	}()

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
