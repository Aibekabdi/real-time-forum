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

func (r *TagRepository) Create(ctx context.Context, tags []models.Tags, postID uint) error {
	for _, tag := range tags {
		var (
			tagID uint
			err   error
		)

		query := `
		WITH inserted_tags AS (
			INSERT INTO tags (name) VALUES ($1)
				ON CONFLICT (name) DO NOTHING
				RETURNING id
			)
			SELECT id FROM inserted_tags
			UNION ALL
			SELECT id FROM tags WHERE name = $1;`

		prep, err := r.db.PrepareContext(ctx, query)
		if err != nil {
			return err
		}

		if err := prep.QueryRowContext(ctx, tag.Name).Scan(&tagID); err != nil {
			return err
		}
		prep.Close()

		if err := r.CreateTagPostConnection(ctx, postID, tagID); err != nil {
			return err
		}
	}
	return nil
}

func (r *TagRepository) CreateTagPostConnection(ctx context.Context, postID, tagID uint) error {
	query := `
	INSERT INTO post_tags (post_id, tag_id)
	SELECT $1, $2
	WHERE NOT EXISTS (
		SELECT 1 FROM post_tags WHERE post_id = $1 AND tag_id = $2
	);
`

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

func (r *TagRepository) Delete(ctx context.Context, tagID uint) error {
	query := "DELETE FROM tags WHERE id = $1;"
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer prep.Close()
	if _, err := prep.ExecContext(ctx, tagID); err != nil {
		return err
	}
	return nil
}
