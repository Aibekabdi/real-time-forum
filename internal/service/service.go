package service

import (
	"context"
	"forum/internal/models"
	"forum/internal/repository"
)

type Auth interface {
	Create(ctx context.Context, user models.User) error
	SignIn(ctx context.Context, user models.SigningInput) (string, error)
	ParseToken(accessToken string) (models.UserToken, error)
}

type Post interface {
	Create(ctx context.Context, post models.Post) (uint, error)
	Delete(ctx context.Context, postID, userID uint) error
	GetALL(ctx context.Context) ([]models.Post, error)
	GetByID(ctx context.Context, postID uint) (models.Post, error)
	InsertorDelete(ctx context.Context, input models.PostVote) error
}

type Comment interface {
	Create(ctx context.Context, comment models.Comments) (uint, error)
	Delete(ctx context.Context, commentID, userID uint) error
	InsertorDelete(ctx context.Context, input models.CommentVote) error
}

type Service struct {
	Auth
	Post
	Comment
}

func NewService(repo *repository.Repository, secretKey string) *Service {
	return &Service{
		Auth:    newAuthService(repo.User, secretKey),
		Post:    newPostService(repo.Post, repo.Tag, repo.Comment),
		Comment: newCommentService(repo.Comment),
	}
}
