package service

import (
	"errors"
	"forum/internal/models"
	"forum/internal/repository"
)

type AuthorizationService struct {
	repo repository.Authorization
}

func newAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (a *AuthorizationService) Signup(input *models.User) error {
	if err := isEmailValid(input.Email); err != nil {
		return err
	} else if err := isPasswordValid(input.Password); err != nil {
		return err
	} else if isEmpty(input.FirstName) {
		return errors.New("empty first name field")
	} else if isEmpty(input.LastName) {
		return errors.New("empty last name field")
	} else if isEmpty(input.Gender) {
		return errors.New("empty gender field")
	} else if isEmpty(input.Age) {
		return errors.New("empty age field")
	}
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		return err
	}
	input.Password = hashedPassword
	if err := a.repo.Signup(input); err != nil {
		return err
	}
	return nil
}
