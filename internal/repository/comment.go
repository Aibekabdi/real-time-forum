package repository

import (
	"context"
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type CommentRepository struct {
	db *sqlx.DB
}

func newCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(ctx context.Context, comment models.Comments) error
func (r *CommentRepository) Delete(ctx context.Context, commentID uint) error
func (r *CommentRepository) GetByPostID(ctx context.Context, postID uint) ([]models.Comments, error)
