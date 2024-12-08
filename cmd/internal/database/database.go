package database

import (
	"database/sql"
	"log"
	"sync"

	"github.com/achintha-dilshan/go-rest-api/cmd/internal/utils"
	"github.com/go-sql-driver/mysql"
)

var (
	DB   *sql.DB
	once sync.Once
)

func Init() {
	once.Do(func() {
		config := mysql.Config{
			Addr:   utils.GetEnv("DB_HOST", "localhost") + ":3306",
			DBName: utils.GetEnv("DB_NAME", "go_rest_api"),
			User:   utils.GetEnv("DB_USER", "root"),
			Passwd: utils.GetEnv("DB_PASSWORD", ""),
		}

		// open the database connection
		var err error
		DB, err = sql.Open("mysql", config.FormatDSN())
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		// verify the connection
		if err = DB.Ping(); err != nil {
			log.Fatalf("Failed to verify the database connection: %v", err)
		}

		log.Println("Database connection successfully initialized")
	})
}
