package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration(db *sql.DB, path string) {
	// Initialize the MySQL driver for migrations
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize MySQL driver for migrations: %v", err)
	}

	// Create a new migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+path,
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// Get the migration command from the last command-line argument
	if len(os.Args) < 2 {
		log.Fatalf("Migration command is required. Use 'up' or 'down'.")
	}
	cmd := os.Args[len(os.Args)-1]

	// Execute the migration command
	switch cmd {
	case "up":
		executeMigration(m.Up, "apply")
	case "down":
		executeMigration(m.Down, "roll back")
	default:
		log.Fatalf("Invalid migration direction: %s. Use 'up' or 'down'.", cmd)
	}
}

// executeMigration handles the execution of migration functions and logs the outcome
func executeMigration(migrationFunc func() error, action string) {
	err := migrationFunc()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to %s migrations: %v", action, err)
	}
	if err == migrate.ErrNoChange {
		log.Printf("No migrations to %s.", action)
	} else {
		log.Printf("Migrations %sed successfully.", action)
	}
}
