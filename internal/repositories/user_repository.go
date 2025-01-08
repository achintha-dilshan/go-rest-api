package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
)

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (int64, error)
	FindById(ctx context.Context, id int64) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int64) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// inserts a new user into the database
func (r *userRepository) Create(ctx context.Context, user *models.User) (int64, error) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// retrieves a user by ID
func (r *userRepository) FindById(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email, password FROM users WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.Id, &user.Name, &user.Email, &user.Password,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &user, err
}

// updates a user's details in the database
func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET name = ?, email = ?, password = ?, updated_at = NOW() WHERE id = ?"
	re, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Id)

	fmt.Println(re.LastInsertId())

	return err
}

// removes a user from the database
func (r *userRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

// checks if a user exists by email
func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists int
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)

	return exists == 1, err
}

// retrieves a user by email
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id, &user.Name, &user.Email, &user.Password,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &user, err
}
