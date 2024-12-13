package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (int64, error)
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int64) error
	ExistUserByEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

// create a new user
func (r *userRepository) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	query := "INSERT INTO users (name, email, password, created_at) VALUES (?, ?, ?, NOW())"
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)

	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID: %w", err)
	}

	return id, nil
}

// get user by id
func (r *userRepository) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email FROM users WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.Id, &user.Name, &user.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, fmt.Errorf("error getting user by ID: %w", err)
	}

	return &user, nil
}

// update user
func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Id)

	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

// delete user
func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}

// check if the user already exists by email
func (r *userRepository) ExistUserByEmail(ctx context.Context, email string) (bool, error) {
	var exists int
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("error checking if user exists by email: %w", err)
	}

	return exists == 1, nil
}

// get user by email
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email FROM users WHERE email = ?"
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id, &user.Name, &user.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, fmt.Errorf("error getting user by email: %w", err)
	}

	return &user, nil
}
