package repository

import (
	"database/sql"
	"forum/internal/models"
)

type Authorization interface {
	Signup(input *models.User) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Authorization: newAuthorizationRepository(db)}
}
