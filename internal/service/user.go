package service

import (
	"context"
	"forum/internal/models"
	"forum/internal/repository"
	"forum/pkg/utils"
)

type UserService struct {
	userRepo  repository.User
	secretKey string
}

func newUserService(userRepo repository.User, secretKey string) *UserService {
	return &UserService{userRepo: userRepo, secretKey: secretKey}
}

func (s *UserService) Create(ctx context.Context, user models.User) error {
	if err := utils.IsValidRegister(&user); err != nil {
		return err
	}
	return s.userRepo.Create(ctx, user)
}

func (s *UserService) SignIn(ctx context.Context, nickname, password string) {
	
}
