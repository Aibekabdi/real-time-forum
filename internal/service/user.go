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

func (s *UserService) UpdatePassword(ctx context.Context, updatePsw models.UpdatePassword, userID uint) error {
	user, err := s.userRepo.GetUserInfo(ctx, userID)
	if err != nil {
		return err
	}

	if err := utils.CompareHashAndPassword(user.Password, updatePsw.OldPassword); err != nil {
		return err
	}

	if err := utils.IsValidPassword(updatePsw.NewPassword); err != nil {
		return err
	}

	newHash, err := utils.GenerateHashPassword(updatePsw.NewPassword)
	if err != nil {
		return err
	}

	if err := s.userRepo.UpdatePassword(ctx, newHash, userID); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserInfo(ctx context.Context, userID uint) (models.User, error) {
	user, err := s.userRepo.GetUserInfo(ctx, userID)
	if err != nil {
		return models.User{}, err
	}
	user.Password = ""
	return user, nil
}
