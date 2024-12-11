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
		// validate required environment variables
		if config.Env.DBHost == "" || config.Env.DBPort == "" || config.Env.DBUser == "" || config.Env.DBName == "" {
			log.Fatal("Database configuration is incomplete. Check environment variables.")
		}

		dbConfig := mysql.Config{
			Addr:                 config.Env.DBHost + ":" + config.Env.DBPort,
			DBName:               config.Env.DBName,
			User:                 config.Env.DBUser,
			Passwd:               config.Env.DBPassword,
			Net:                  "tcp",
			AllowNativePasswords: true, // ensures compatibility with MySQL native password authentication
		}

		// open the database connection
		var err error
		DB, err = sql.Open("mysql", dbConfig.FormatDSN())
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		// Configure connection pool settings
		DB.SetMaxOpenConns(10)   // Maximum number of open connections to the database
		DB.SetMaxIdleConns(5)    // Maximum number of idle connections in the pool
		DB.SetConnMaxIdleTime(0) // Maximum time a connection can remain idle

		// verify the connection
		if err = DB.Ping(); err != nil {
			log.Fatalf("Failed to verify the database connection: %v", err)
		}

		log.Println("Database connection successfully initialized")
	})
}

// closes the database connection gracefully
func Close() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection successfully closed")
		}
	}
}
