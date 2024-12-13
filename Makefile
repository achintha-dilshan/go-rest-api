build:
	@go build -o bin/go-rest-api cmd/main.go

run: build
	@./bin/go-rest-api

test:
	@go test -v ./..