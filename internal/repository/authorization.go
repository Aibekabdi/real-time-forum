package repository

import (
	"forum/internal/models"

	"errors"

	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository struct {
	db *sqlx.DB
}

func newAuthorizationRepository(db *sqlx.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

func (a AuthorizationRepository) Signup(input *models.User) error {
	if _, err := a.db.Exec("INSERT INTO User(email, nickname, first_name, last_name, password, gender, age) VALUES (?,?,?,?,?,?,?)", input.Email, input.Nickname, input.FirstName, input.LastName, input.Password, input.Gender, input.Age); err != nil {
		return errors.New("email or username is already exist, try another ones")
	}
	return nil
}
