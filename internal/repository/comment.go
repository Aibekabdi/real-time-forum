package repository

import (
	"context"
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type CommentRepository struct {
	db *sqlx.DB
}

func newCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(ctx context.Context, comment models.Comments) (uint, error) {
	var (
		id  uint
		err error
	)
	query := "INSERT INTO comments (user_id, post_id, text) VALUES ($1, $2, $3) RETURNING id;"

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer prep.Close()

	if err := prep.QueryRowContext(ctx, comment.Author.ID, comment.PostId, comment.Text).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CommentRepository) Delete(ctx context.Context, commentID, userID uint) error {
	query := "DELETE FROM comments WHERE id = $1 and user_id ;"
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer prep.Close()

	if _, err := prep.ExecContext(ctx, commentID, userID); err != nil {
		return err
	}
	return nil
}

func (r *CommentRepository) GetByPostID(ctx context.Context, postID uint) ([]models.Comments, error) {
	query := `
	SELECT c.id, c.title, c.text, u.id, u.nickname
	FROM comments c
	INNER JOIN users u ON c.user_id = u.id
	WHERE c.post_id = $1;
	`

	prep, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer prep.Close()

	rows, err := prep.QueryContext(ctx, postID)
	if err != nil {
		return nil, err
	}

	var comments []models.Comments
	for rows.Next() {
		comment := models.Comments{}

		if err := rows.Scan(&comment.ID, &comment.Text, &comment.Author.ID, &comment.Author.Nickname); err != nil {
			return nil, err
		}

		comment.PostId = postID
		comments = append(comments, comment)
	}

	return comments, nil
}
