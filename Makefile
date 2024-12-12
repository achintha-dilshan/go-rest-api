build:
	@go build -o bin/go-rest-api cmd/main.go

run: build
	@./bin/go-rest-api

test:
	@go test -v ./..

migration:
	@migrate create -ext sql -dir database/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/main.go up

migrate-down:
	@go run cmd/main.go down