package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/achintha-dilshan/go-rest-api/config"
	"github.com/go-sql-driver/mysql"
)

type database struct {
	db       *sql.DB
	once     sync.Once
	isClosed bool
	mu       sync.Mutex
}

type Database interface {
	Connect() error
	Close() error
	GetDB() (*sql.DB, error)
}

func NewDatabase() *database {
	return &database{}
}

func (h *database) Connect() error {
	var initErr error
	h.once.Do(func() {
		if config.Env.DBHost == "" || config.Env.DBPort == "" || config.Env.DBUser == "" || config.Env.DBName == "" {
			initErr = errors.New("database configuration is incomplete")
			log.Println(initErr)
			return
		}

		dbConfig := mysql.Config{
			Addr:                 fmt.Sprintf("%s:%s", config.Env.DBHost, config.Env.DBPort),
			DBName:               config.Env.DBName,
			User:                 config.Env.DBUser,
			Passwd:               config.Env.DBPassword,
			Net:                  "tcp",
			AllowNativePasswords: true,
		}

		db, err := sql.Open("mysql", dbConfig.FormatDSN())
		if err != nil {
			initErr = fmt.Errorf("failed to connect to the database: %w", err)
			log.Println(initErr)
			return
		}

		// Configure connection pool settings
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(5)
		db.SetConnMaxIdleTime(0)

		if err := db.Ping(); err != nil {
			initErr = fmt.Errorf("failed to verify the database connection: %w", err)
			log.Println(initErr)
			return
		}

		h.db = db
		log.Println("Database connection successfully initialized")
	})

	return initErr
}

func (h *database) Close() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.db != nil && !h.isClosed {
		if err := h.db.Close(); err != nil {
			return fmt.Errorf("error closing database connection: %w", err)
		}
		h.isClosed = true
		log.Println("Database connection successfully closed")
	}

	return nil
}

func (h *database) GetDB() (*sql.DB, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.isClosed {
		return nil, errors.New("attempt to use a closed database connection")
	}

	return h.db, nil
}
