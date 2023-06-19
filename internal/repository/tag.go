package repository

import (
	"context"
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type TagRepository struct {
	db *sqlx.DB
}

func newTagRepository(db *sqlx.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) Create(ctx context.Context, tag models.Tags, postID uint) error {
	var (
		tagID uint
		err   error
	)
	query := "INSERT INTO tags (text) VALUES ($1) RETURNING id;"

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	if err := prep.QueryRowContext(ctx, tag.Text).Scan(&tagID); err != nil {
		return err
	}
	prep.Close()

	if err := r.CreateTagPostConnection(ctx, tagID, postID); err != nil {
		return err
	}

	return nil
}

func (r *TagRepository) CreateTagPostConnection(ctx context.Context, tagID uint, postID uint) error {
	query := "INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2);"

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer prep.Close()

	if _, err := prep.ExecContext(ctx, postID, tagID); err != nil {
		return err
	}

	return nil
}
