package routes

import (
	"github.com/achintha-dilshan/go-rest-api/cmd/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(router chi.Router) {
	authHandler := &handlers.AuthHandler{}

	router.Get("/login", authHandler.Login)
	router.Get("/register", authHandler.Register)
}
