package routes

import (
	"database/sql"

	"github.com/achintha-dilshan/go-rest-api/internal/handlers"
	"github.com/achintha-dilshan/go-rest-api/internal/middlewares"
	"github.com/achintha-dilshan/go-rest-api/internal/repositories"
	"github.com/achintha-dilshan/go-rest-api/internal/services"
	"github.com/go-chi/chi/v5"
)

type postRoutes struct {
	db *sql.DB
}

type PostRoutes interface {
	Get() *chi.Mux
}

func NewPostRoutes(db *sql.DB) PostRoutes {
	return &postRoutes{
		db: db,
	}
}

func (r *postRoutes) Get() *chi.Mux {
	router := chi.NewRouter()

	repo := repositories.NewPostRepository(r.db)
	service := services.NewPostService(repo)
	handler := handlers.NewPostHandler(service)

	router.Get("/", handler.GetAllPosts)
	router.Get("/{id}", handler.GetSinglePost)
	router.With(middlewares.AuthMiddleware).Post("/", handler.CreatePost)
	router.With(middlewares.AuthMiddleware).Patch("/{id}", handler.EditPost)
	router.With(middlewares.AuthMiddleware).Delete("/{id}", handler.DeletePost)

	return router
}
