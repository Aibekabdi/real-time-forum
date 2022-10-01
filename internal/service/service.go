package service

import (
	"forum/internal/models"
	"forum/internal/repository"
)

type Authorization interface {
	Signup(input *models.User) error
}

type Service struct {
	Authorization
}

func NewService(r *repository.Repository) *Service {
	return &Service{Authorization: newAuthorizationService(r.Authorization)}
}
