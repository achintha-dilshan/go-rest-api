package services

import (
	"context"
	"errors"

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
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository: repository}
}

// create a new user
func (s *userService) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	exists, err := s.repository.ExistUserByEmail(ctx, user.Email)

	if err != nil {
		return 0, err
	}

	if exists {
		return 0, errors.New("email already exists")
	}

	return s.repository.CreateUser(ctx, user)
}

// get user by id
func (s *userService) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.repository.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// update a user
func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	existingUser, err := s.repository.GetUserById(ctx, user.Id)

	if err != nil {
		return errors.New("user not found")
	}

	if existingUser.Email != user.Email {
		emailExists, err := s.repository.ExistUserByEmail(ctx, user.Email)

		if err != nil {
			return err
		}

		if emailExists {
			return errors.New("email already exists")
		}
	}

	return s.repository.UpdateUser(ctx, user)
}

// delete a user
func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	_, err := s.repository.GetUserById(ctx, id)

	if err != nil {
		return errors.New("user not found")
	}

	return s.repository.DeleteUser(ctx, id)
}

// get user by email
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.repository.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
