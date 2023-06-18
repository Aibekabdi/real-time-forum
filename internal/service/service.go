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
type Service struct {
	Auth
}

func NewService(repo *repository.Repository, secretKey string) *Service {
	return &Service{
		Auth: newAuthService(repo.User, secretKey),
	}
}
