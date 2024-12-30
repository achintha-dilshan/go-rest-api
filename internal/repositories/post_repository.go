package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
)

type postRepository struct {
	db *sql.DB
}

type PostRepository interface {
	Create(ctx context.Context, post *models.Post) (int64, error)
	FindById(ctx context.Context, id int64) (*models.Post, error)
	Update(ctx context.Context, post *models.Post) error
	Delete(ctx context.Context, id int64) error
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

// inserts a new post into the database
func (r *postRepository) Create(ctx context.Context, post *models.Post) (int64, error) {
	query := "INSERT INTO posts (title, body) VALUES (?, ?)"
	result, err := r.db.ExecContext(ctx, query, post.Title, post.Body)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// retrieves a post by ID
func (r *postRepository) FindById(ctx context.Context, id int64) (*models.Post, error) {
	var post models.Post
	query := "SELECT id, title, body FROM posts WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&post.Id, &post.Title, &post.Body,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &post, err
}

// updates a post's details in the database
func (r *postRepository) Update(ctx context.Context, post *models.Post) error {
	query := "UPDATE posts SET title = ?, body = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, post.Title, post.Body, post.Id)

	return err
}

// removes a post from the database
func (r *postRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM posts WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
