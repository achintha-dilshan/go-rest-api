package routes

import (
	"database/sql"

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

	router.Get("/", nil)
	router.Post("/", nil)
	router.Get("/{id}", nil)
	router.Post("/{id}/edit", nil)
	router.Post("/{id}/delete", nil)

	return router
}
