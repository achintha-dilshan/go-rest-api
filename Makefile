# Variables
BUILD_OUTPUT = bin/go-rest-api
DB_DRIVER = mysql
DB_STRING = root@tcp(127.0.0.1:3306)/go_rest_api?parseTime=true
MIGRATIONS_DIR = database/migrations
GOOSE_CMD = goose

# Build the Go application
build:
	@go build -o $(BUILD_OUTPUT) cmd/main.go

# Run the built application
run: build
	@./$(BUILD_OUTPUT)

# Run tests
test:
	@go test -v ./...

# Goose migration commands
create-migration:
	@goose -dir $(MIGRATIONS_DIR) create $(name) sql

migrate-up:
	@goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) up

migrate-down:
	@goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) down

migrate-rollback:
	@goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) down 1

migrate-status:
	@goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) status

# Help for available commands
help:
	@echo "Available commands:"
	@echo "  make build                  - Build the application"
	@echo "  make run                    - Run the application"
	@echo "  make test                   - Run tests"
	@echo "  make create-migration name=<migration_name>  - Create a new migration with the given name"
	@echo "  make migrate-up             - Apply all available migrations"
	@echo "  make migrate-down           - Rollback all migrations"
	@echo "  make migrate-rollback       - Rollback the last migration"
	@echo "  make migrate-status         - Show migration status"
