# Simple REST API

This is a simple REST API built using GoLang, featuring the [Chi Router](https://github.com/go-chi/chi) for routing, [Goose](https://github.com/pressly/goose) for database migrations, and MySQL as the database. The API implements JWT-based authentication for secure access to protected endpoints.

## Features

- User authentication and management
- JWT authentication for secure API access
- CRUD operations for posts
- Database migrations with Goose

## Prerequisites

- GoLang installed (1.19 or higher recommended)
- MySQL server running
- Goose CLI installed for database migrations

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/achintha-dilshan/go-rest-api.git
   cd go-rest-api
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up the database:

   - Create a MySQL database.
   - Configure the database connection in the `.env` file.
   - Run database migrations

4. Start the server:

   ```bash
   go run main.go
   ```

## Endpoints

### Authentication

- **POST /auth/register**
  - Registers a new user.

- **POST /auth/login**
  - Logs in a user and returns a JWT token.

### User Management

- **POST /user/password-reset**
  - Resets the password for the logged-in user.

- **PATCH /user/update**
  - Updates the profile of the logged-in user.

- **DELETE /user/delete**
  - Deletes the logged-in user's account.

### Posts

- **GET /posts**
  - Retrieves all posts.

- **GET /posts/{id}**
  - Retrieves a single post by its ID.

- **POST /posts**
  - Creates a new post.

- **PATCH /posts/{id}**
  - Updates an existing post by its ID.

- **DELETE /posts/{id}**
  - Deletes an existing post by its ID.

## Usage

1. Use an API client like Postman or cURL to test the endpoints.
2. Include the JWT token in the `Authorization` header for protected routes:

   ```http
   Authorization: Bearer <your-jwt-token>
   ```

## Feedback

I'm a beginner and would greatly appreciate any thoughts and advice you may have. Feel free to create issues or share suggestions on how to improve this project.
