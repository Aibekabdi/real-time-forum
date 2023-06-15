package service

import (
	"context"
	"forum/internal/models"
	"forum/internal/repository"
	"forum/pkg/utils"
)

type UserService struct {
	userRepo repository.User
}

func newUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) Create(ctx context.Context, user models.User) error {
	if err := utils.IsValidRegister(&user); err != nil {
		return err
	}
	return u.userRepo.Create(ctx, user)
}
