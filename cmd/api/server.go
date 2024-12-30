package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/config"
)

type apiServer struct {
	db *sql.DB
}

type APIServer interface {
	Run() error
}

func NewAPIServer(db *sql.DB) APIServer {
	return &apiServer{
		db: db,
	}
}

func (s *apiServer) Run() error {
	port := ":" + config.Env.ServerPort

	server := &http.Server{
		Addr:    port,
		Handler: NewRouter(s.db).Init(),
	}

	log.Printf("Server is running on port %v", port)

	return server.ListenAndServe()
}
