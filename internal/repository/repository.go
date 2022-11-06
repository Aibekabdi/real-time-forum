package repository

import (
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Signup(input *models.User) error
}

// todo repository
type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: newAuthorizationRepository(db)}
}
