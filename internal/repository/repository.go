package repository

import "database/sql"

// todo repository
type Repository struct{}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
