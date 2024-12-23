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
	FindUserById(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int64) error
	ExistUserByEmail(ctx context.Context, email string) (bool, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository: repository}
}

// create a new user
func (s *userService) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	return s.repository.Create(ctx, user)
}

// find user by id
func (s *userService) FindUserById(ctx context.Context, id int64) (*models.User, error) {
	return s.repository.FindById(ctx, id)
}

// update a user
func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	return s.repository.Update(ctx, user)
}

// delete a user
func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

// exist user by email
func (s *userService) ExistUserByEmail(ctx context.Context, email string) (bool, error) {
	return s.repository.ExistsByEmail(ctx, email)
}

// find user by email
func (s *userService) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repository.FindByEmail(ctx, email)
}
