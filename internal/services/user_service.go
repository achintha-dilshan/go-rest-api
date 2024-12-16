package services

import (
	"context"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
	"github.com/achintha-dilshan/go-rest-api/internal/repositories"
)

type userService struct {
	repository repositories.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) (int64, error)
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int64) error
	ExistUserByEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository: repository}
}

// create a new user
func (s *userService) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	return s.repository.CreateUser(ctx, user)
}

// get user by id
func (s *userService) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return s.repository.GetUserById(ctx, id)
}

// update a user
func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	return s.repository.UpdateUser(ctx, user)
}

// delete a user
func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return s.repository.DeleteUser(ctx, id)
}

// exist user by email
func (s *userService) ExistUserByEmail(ctx context.Context, email string) (bool, error) {
	return s.repository.ExistUserByEmail(ctx, email)
}

// get user by email
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repository.GetUserByEmail(ctx, email)
}
