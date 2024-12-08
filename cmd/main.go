package main

import (
	"log"
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/cmd/internal/routes"
)

func main() {
	router := routes.Init()

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Printf("Server is running on port %v", ":8000")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
