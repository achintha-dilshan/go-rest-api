package api

import (
	"database/sql"
	"log"
	"net/http"
)

type apiServer struct {
	addr string
	db   *sql.DB
}

type APIServer interface {
	Run() error
}

func NewAPIServer(addr string, db *sql.DB) APIServer {
	return &apiServer{
		addr: addr,
		db:   db,
	}
}

func (s *apiServer) Run() error {
	router := NewRouter(s.db)

	server := &http.Server{
		Addr:    s.addr,
		Handler: router.Init(),
	}

	log.Printf("Server is running on port %v", s.addr)

	return server.ListenAndServe()
}
