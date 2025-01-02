package api

import (
	"database/sql"

	"github.com/achintha-dilshan/go-rest-api/internal/routes"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	// Auth Routes
	router.Mount("/auth", routes.NewAuthRoutes(r.db).Get())

	// User Routes
	router.Mount("/user", routes.NewUserRoutes(r.db).Get())

	// Post Routes
	router.Mount("/posts", routes.NewPostRoutes(r.db).Get())

	return router
}
