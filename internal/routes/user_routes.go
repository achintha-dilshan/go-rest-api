package routes

import (
	"database/sql"

	"github.com/achintha-dilshan/go-rest-api/internal/handlers"
	"github.com/achintha-dilshan/go-rest-api/internal/middlewares"
	"github.com/achintha-dilshan/go-rest-api/internal/repositories"
	"github.com/achintha-dilshan/go-rest-api/internal/services"
	"github.com/go-chi/chi/v5"
)

type userRoutes struct {
	db *sql.DB
}

type UserRoutes interface {
	Get() *chi.Mux
}

func NewUserRoutes(db *sql.DB) UserRoutes {
	return &userRoutes{db: db}
}

func (r *userRoutes) Get() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middlewares.AuthMiddleware)

	repo := repositories.NewUserRepository(r.db)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	router.Patch("/password-reset", handler.ResetPassword)
	router.Patch("/update", handler.UpdateUser)
	router.Delete("/delete", handler.DeleteUser)

	return router
}
