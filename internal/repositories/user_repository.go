package repositories

import (
	"context"
	"database/sql"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (int64, error)
	GetById(ctx context.Context, id int64) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int64) error
	ExistByEmail(ctx context.Context, email string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// create a new user
func (r *userRepository) Create(ctx context.Context, user *models.User) (int64, error) {
	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, NOW(), NULL)"
	result, err := r.db.Exec(query, user.Name, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

// get user by id
func (r *userRepository) GetById(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email FROM users WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(
		&user.Id, &user.Name, &user.Email,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// update user
func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Id)

	return err
}

// delete user
func (r *userRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)

	return err
}

// check if the user already exists by email
func (r *userRepository) ExistByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	err := r.db.QueryRow(query, email).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

// get user by email
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email FROM users WHERE email = ?"
	err := r.db.QueryRow(query, email).Scan(
		&user.Id, &user.Name, &user.Email,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
