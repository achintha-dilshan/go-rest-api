package api

import (
	"database/sql"

	"github.com/achintha-dilshan/go-rest-api/internal/routes"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type router struct {
	db *sql.DB
}

type Router interface {
	Init() *chi.Mux
}

func NewRouter(db *sql.DB) Router {
	return &router{
		db: db,
	}
}

func (r *router) Init() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Auth Routes
	router.Mount("/auth", routes.NewAuthRoutes(r.db).Get())

	return router
}
