package repository

import (
	"database/sql"
	"forum/internal/models"
)

type AuthorizationRepository struct {
	db *sql.DB
}

func newAuthorizationRepository(db *sql.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

func (a AuthorizationRepository) Signup(input *models.User) error {
	return nil
}
