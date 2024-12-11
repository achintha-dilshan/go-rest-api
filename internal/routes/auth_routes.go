package routes

import (
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func AuthRoutes() http.Handler {
	r := chi.NewRouter()

	authHandler := handlers.NewAuthHandler()

	r.Post("/login", authHandler.LoginUser)
	r.Post("/register", authHandler.RegisterUser)

	return r
}
