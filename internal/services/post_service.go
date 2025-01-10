package services

import (
	"context"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
	"github.com/achintha-dilshan/go-rest-api/internal/repositories"
)

type postService struct {
	repository repositories.PostRepository
}

type PostService interface {
	CreatePost(ctx context.Context, post *models.Post) (int64, error)
	FindAll(ctx context.Context) ([]*models.Post, error)
	FindPostById(ctx context.Context, id int64) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id int64) error
}

func NewPostService(repository repositories.PostRepository) PostService {
	return &postService{
		repository: repository,
	}
}

// create a new post
func (s *postService) CreatePost(ctx context.Context, post *models.Post) (int64, error) {
	return s.repository.Create(ctx, post)
}

// find all posts
func (s *postService) FindAll(ctx context.Context) ([]*models.Post, error) {
	return s.repository.FindAll(ctx)
}

// find post by id
func (s *postService) FindPostById(ctx context.Context, id int64) (*models.Post, error) {
	return s.repository.FindById(ctx, id)
}

// update a post
func (s *postService) UpdatePost(ctx context.Context, post *models.Post) error {
	return s.repository.Update(ctx, post)
}

// delete a post
func (s *postService) DeletePost(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}
