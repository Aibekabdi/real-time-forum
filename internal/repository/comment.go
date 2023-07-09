package repository

import (
	"context"
	"forum/internal/models"
	"log"

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
	SELECT c.id, c.text, u.id, u.nickname
	FROM comments c
	INNER JOIN users u ON c.user_id = u.id
	WHERE c.post_id = $1;
	`

	prep, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

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
	prep.Close()
	for _, comment := range comments {
		votes, err := r.getLikesAndDislikes(comment.ID)
		if err != nil {
			return nil, err
		}
		comment.Vote = votes
	}
	return comments, nil
}

func (r *CommentRepository) InsertorDelete(ctx context.Context, input models.CommentVote) error {
	var (
		query string
	)
	exists, err := r.checkCommentLikeExists(ctx, input.CommentID, input.UserID)
	if err != nil {
		return err
	}
	if exists {
		currentType, err := r.getCommentLikeType(ctx, input.CommentID, input.UserID)
		if err != nil {
			return nil
		}
		if currentType != input.LikeType {
			query = `UPDATE comments_likes SET type = $1 WHERE comment_id = $2 AND user_id = $3`
			prep, err := r.db.PrepareContext(ctx, query)
			if err != nil {
				return err
			}
			defer prep.Close()
			if _, err := prep.ExecContext(ctx, input.LikeType, input.CommentID, input.UserID); err != nil {
				return err
			}
			log.Printf("user : %v's vote updated in post:%v\n", input.UserID, input.CommentID)
			return nil
		} else {
			query = `DELETE FROM comments_likes WHERE comment_id = $1 AND user_id = $2`
			prep, err := r.db.PrepareContext(ctx, query)
			if err != nil {
				return err
			}
			defer prep.Close()
			if _, err := prep.ExecContext(ctx, input.CommentID, input.UserID); err != nil {
				return err
			}
			log.Printf("user : %v's vote deleted in post:%v\n", input.UserID, input.CommentID)
			return nil
		}
	} else {
		query = `INSERT INTO comments_likes (comment_id, user_id, type) VALUES ($1, $2, $3)`
		prep, err := r.db.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer prep.Close()
		if _, err := prep.ExecContext(ctx, input.CommentID, input.UserID, input.LikeType); err != nil {
			return err
		}
	}
	return nil
}

func (r *CommentRepository) checkCommentLikeExists(ctx context.Context, commentID, userID uint) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM comments_likes WHERE comment_id = $1 AND user_id = $2)",
		commentID, userID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *CommentRepository) getCommentLikeType(ctx context.Context, commentID, userID uint) (int, error) {
	var likeType int
	err := r.db.QueryRowContext(ctx, "SELECT type FROM comments_likes WHERE comment_id = $1 AND user_id = $2",
		commentID, userID).Scan(&likeType)
	if err != nil {
		return 0, err
	}
	return likeType, nil
}

func (r *CommentRepository) getLikesAndDislikes(commentID uint) (models.Vote, error) {
	query := `
	SELECT 
		COALESCE(COUNT(CASE WHEN type = -1 THEN 1 END), 0) AS dislikes,
		COALESCE(COUNT(CASE WHEN type = 1 THEN 1 END), 0) AS likes
	FROM posts_likes
	WHERE post_id = $1;
	`

	var dislikes, likes uint
	err := r.db.QueryRow(query, commentID).Scan(&dislikes, &likes)
	if err != nil {
		return models.Vote{}, err
	}
	vote := models.Vote{
		Likes:    likes,
		Dislikes: dislikes,
	}
	return vote, nil
}
