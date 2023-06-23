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
}
type Service struct {
	Auth
	Post
}

func NewService(repo *repository.Repository, secretKey string) *Service {
	return &Service{
		Auth: newAuthService(repo.User, secretKey),
		Post: newPostService(repo.Post, repo.Tag),
	}
}
