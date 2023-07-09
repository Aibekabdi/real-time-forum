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
	GetByID(ctx context.Context, postID uint) (models.Post, error)
	InsertorDelete(ctx context.Context, input models.PostVote) error
	GetALLByTag(ctx context.Context, tagName string) ([]models.Post, error)
	GetALLByUserID(ctx context.Context, userID uint) ([]models.Post, error)
}

type Tag interface {
	Create(ctx context.Context, tag []models.Tags, postID uint) error
	CreateTagPostConnection(ctx context.Context, tagID uint, postID uint) error
	Delete(ctx context.Context, tagID uint) error
	GetByPostID(ctx context.Context, postID uint) ([]models.Tags, error)
}

type Comment interface {
	Create(ctx context.Context, comment models.Comments) (uint, error)
	Delete(ctx context.Context, commentID, userID uint) error
	GetByPostID(ctx context.Context, postID uint) ([]models.Comments, error)
	InsertorDelete(ctx context.Context, input models.CommentVote) error
}

type Repository struct {
	User
	Post
	Tag
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    newUserRepository(db),
		Post:    newPostRepository(db),
		Tag:     newTagRepository(db),
		Comment: newCommentRepository(db),
	}
}
