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

func (r *CommentRepository) Create(ctx context.Context, comment models.Comments) error
func (r *CommentRepository) Delete(ctx context.Context, commentID uint) error

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

		if err := rows.Scan(&comment.ID, &comment.Title, &comment.Text, &comment.Author.ID, &comment.Author.Nickname); err != nil {
			return nil, err
		}

		comment.PostId = postID
		comments = append(comments, comment)
	}

	return comments, nil
}
