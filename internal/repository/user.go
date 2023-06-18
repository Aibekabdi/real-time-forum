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

func (r *UserRepository) Create(ctx context.Context, user models.User) error {
	query := `INSERT INTO users (nickname, gender, age, firstname, lastname, email, password) VALUES($1, $2, $3, $4, $5, $6, $7)`
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer prep.Close()
	if _, err = prep.ExecContext(ctx, user.Nickname, user.Gender, user.Age, user.FirstName, user.LastName, user.Email, user.Password); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, loggindField string) (uint, string, error) {
	query := `SELECT id, password FROM users WHERE nickname = $1 LIMIT 1;`
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, "", err
	}

	defer prep.Close()

	var (
		id       uint
		password string
	)

	if err = prep.QueryRowContext(ctx, loggindField).Scan(&id, &password); err != nil {
		return 0, "", err
	}

	return id, password, nil
}
