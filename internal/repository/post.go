package repository

import (
	"context"
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	db *sqlx.DB
}

func newPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, post models.Post) (uint, error) {
	var (
		id  uint
		err error
	)
	query := "INSERT INTO posts (user_id, title, text) VALUES ($1, $2, $3) RETURNING id;"

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	defer prep.Close()

	if err := prep.QueryRowContext(ctx, post.Author.Id, post.Title, post.Text).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
