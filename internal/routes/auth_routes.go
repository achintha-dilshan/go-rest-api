package routes

import (
	"database/sql"

	"github.com/achintha-dilshan/go-rest-api/internal/handlers"
	"github.com/achintha-dilshan/go-rest-api/internal/repositories"
	"github.com/achintha-dilshan/go-rest-api/internal/services"
	"github.com/go-chi/chi/v5"
)

type authRoutes struct {
	db *sql.DB
}

type AuthRoutes interface {
	Get() *chi.Mux
}

func NewAuthRoutes(db *sql.DB) AuthRoutes {
	return &authRoutes{db: db}
}

func (r *authRoutes) Get() *chi.Mux {
	router := chi.NewRouter()

	repo := repositories.NewUserRepository(r.db)
	service := services.NewUserService(repo)
	handler := handlers.NewAuthHandler(service)

	router.Post("/login", handler.LoginUser)
	router.Post("/register", handler.RegisterUser)

	return router
}
