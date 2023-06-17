package service

import (
	"context"
	"forum/internal/models"
	"forum/internal/repository"
)

type User interface {
	Create(ctx context.Context, user models.User) error
}
type Service struct {
	User
}

func NewService(repo *repository.Repository, secretKey string) *Service {
	return &Service{
		User: newUserService(repo.User, secretKey),
	}
}
