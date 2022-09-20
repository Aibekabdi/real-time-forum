package service

import "forum/internal/repository"

type Authorization interface {
}

type Service struct {
}

func NewService(r *repository.Repository) *Service {
	return &Service{}
}
