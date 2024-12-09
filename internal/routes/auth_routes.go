package routes

import (
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func AuthRoutes() http.Handler {
	router := chi.NewRouter()
	router.Get("/login", handlers.LoginHandler)
	router.Get("/register", handlers.RegisterHandler)

	return router
}
