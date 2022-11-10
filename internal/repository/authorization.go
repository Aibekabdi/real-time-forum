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
	stmt, err := a.db.Prepare("INSERT INTO users(email, username, firstname, lastname, password_hash, gender, age) VALUES ($1,$2,$3,$4,$5,$6,$7);")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(input.Email, input.Nickname, input.FirstName, input.LastName, input.Password, input.Gender, input.Age); err != nil {
		return err
	}
	stmt.Close()
	return nil
}
