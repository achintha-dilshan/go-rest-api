package routes

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Init() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Auth Routes
	RegisterAuthRoutes(router)

	return router
}
