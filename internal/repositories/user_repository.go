package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
)

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (int64, error)
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int64) error
	ExistUserByEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser inserts a new user into the database
func (r *userRepository) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error retrieving last insert ID: %v", err)
		return 0, err
	}

	return id, nil
}

// GetUserById retrieves a user by ID
func (r *userRepository) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email FROM users WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.Id, &user.Name, &user.Email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		log.Printf("Error retrieving user by ID: %v", err)
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates a user's details in the database
func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Id)

	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}

// DeleteUser removes a user from the database
func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	return nil
}

// ExistUserByEmail checks if a user exists by email
func (r *userRepository) ExistUserByEmail(ctx context.Context, email string) (bool, error) {
	var exists int
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)

	if err != nil {
		log.Printf("Error checking if user exists: %v", err)
		return false, err
	}

	return exists == 1, nil
}

// GetUserByEmail retrieves a user by email
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id, &user.Name, &user.Email, &user.Password,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		log.Printf("Error retrieving user by email: %v", err)
		return nil, err
	}

	return &user, nil
}
