package database

import (
	"database/sql"
	"log"
	"sync"

	"github.com/achintha-dilshan/go-rest-api/config"
	"github.com/go-sql-driver/mysql"
)

var (
	DB   *sql.DB
	once sync.Once
)

func Init() {
	once.Do(func() {
		dbConfig := mysql.Config{
			Addr:   config.Env.DBHost + ":" + config.Env.DBPort,
			DBName: config.Env.DBName,
			User:   config.Env.DBUser,
			Passwd: config.Env.DBPassword,
		}

		// open the database connection
		var err error
		DB, err = sql.Open("mysql", dbConfig.FormatDSN())
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
