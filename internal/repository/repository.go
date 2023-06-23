package repository

import (
	"context"
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, loggindField string) (uint, string, error)
}

type Post interface {
	Create(ctx context.Context, post models.Post) (uint, error)
	Delete(ctx context.Context, postID, userID uint) error
	GetALL(ctx context.Context) ([]models.Post, error)
}

type Tag interface {
	Create(ctx context.Context, tag []models.Tags, postID uint) error
	CreateTagPostConnection(ctx context.Context, tagID uint, postID uint) error
	Delete(ctx context.Context, tagID uint) error
}

type Repository struct {
	User
	Post
	Tag
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: newUserRepository(db),
		Post: newPostRepository(db),
		Tag:  newTagRepository(db),
	}
}
