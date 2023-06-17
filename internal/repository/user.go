package repository

import (
	"context"
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func newUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(ctx context.Context, user models.User) error {
	query := `INSERT INTO users (nickname, gender, age, firstname, lastname, email, password) VALUES($1, $2, $3, $4, $5, $6, $7)`
	prep, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer prep.Close()
	if _, err = prep.ExecContext(ctx, user.Nickname, user.Gender, user.Age, user.FirstName, user.LastName, user.Email, user.Password); err != nil {
		return err
	}
	return nil
}
