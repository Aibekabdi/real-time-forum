package repository

import (
	"context"
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(ctx context.Context, user models.User) error
}
type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: newUserRepository(db),
	}
}
