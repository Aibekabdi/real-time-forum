package repository

import (
	"context"
	"database/sql"
	"forum/internal/models"
	"log"

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

	if err := prep.QueryRowContext(ctx, post.Author.ID, post.Title, post.Text).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostRepository) Delete(ctx context.Context, postID, userID uint) error {
	query := "DELETE FROM posts WHERE id = $1 and user_id = $2;"
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer prep.Close()
	if _, err := prep.ExecContext(ctx, postID, userID); err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) GetALL(ctx context.Context) ([]models.Post, error) {
	query := `
			SELECT p.id, p.title, u.id, u.nickname, t.id, t.name
			FROM posts p
			INNER JOIN users u ON p.user_id = u.id
			LEFT JOIN post_tags pt ON pt.post_id = p.id
			LEFT JOIN tags t ON pt.tag_id = t.id;
		`
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	rows, err := prep.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	posts := make(map[uint]*models.Post)
	for rows.Next() {
		var (
			postID   uint
			title    string
			authorID uint
			author   string
			tagID    sql.NullInt64
			tagName  sql.NullString
		)
		err := rows.Scan(&postID, &title, &authorID, &author, &tagID, &tagName)
		if err != nil {
			return nil, err
		}
		if _, ok := posts[postID]; !ok {
			posts[postID] = &models.Post{
				ID:     postID,
				Title:  title,
				Author: models.User{ID: authorID, Nickname: author},
			}
		}
		if tagID.Valid && tagName.Valid {
			posts[postID].Tags = append(posts[postID].Tags, models.Tags{
				ID:   uint(tagID.Int64),
				Name: tagName.String,
			})
		}
	}
	var result []models.Post
	for _, post := range posts {
		result = append(result, *post)
	}
	return result, nil
}

func (r *PostRepository) GetByID(ctx context.Context, postID uint) (models.Post, error) {
	query := `
		SELECT p.id, p.title, p.text, u.id, u.nickname
		FROM posts p
		INNER JOIN users u ON p.user_id = u.id
		WHERE p.id = $1;
	`
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return models.Post{}, err
	}
	post := models.Post{}
	if err := prep.QueryRowContext(ctx, postID).Scan(&post.ID, &post.Title, &post.Text, &post.Author.ID, &post.Author.Nickname); err != nil {
		return models.Post{}, err
	}
	prep.Close()
	votes, err := r.getLikesAndDislikes(postID)
	if err != nil {
		return models.Post{}, err
	}
	post.Vote = votes
	return post, nil
}

func (r *PostRepository) InsertorDelete(ctx context.Context, input models.PostVote) error {
	var (
		query string
	)
	exists, err := r.checkPostLikeExists(ctx, input.PostID, input.UserID)
	if err != nil {
		return err
	}
	if exists {
		currentType, err := r.getPostLikeType(ctx, input.PostID, input.UserID)
		if err != nil {
			return nil
		}
		if currentType != input.LikeType {
			query = `UPDATE posts_likes SET type = $1 WHERE post_id = $2 AND user_id = $3`
			prep, err := r.db.PrepareContext(ctx, query)
			if err != nil {
				return err
			}
			defer prep.Close()
			if _, err := prep.ExecContext(ctx, input.LikeType, input.PostID, input.UserID); err != nil {
				return err
			}
			log.Printf("user : %v's vote updated in post:%v\n", input.UserID, input.PostID)
			return nil
		} else {
			query = `DELETE FROM posts_likes WHERE post_id = $1 AND user_id = $2`
			prep, err := r.db.PrepareContext(ctx, query)
			if err != nil {
				return err
			}
			defer prep.Close()
			if _, err := prep.ExecContext(ctx, input.PostID, input.UserID); err != nil {
				return err
			}
			log.Printf("user : %v's vote deleted in post:%v\n", input.UserID, input.PostID)
			return nil
		}
	} else {
		query = `INSERT INTO posts_likes (post_id, user_id, type) VALUES ($1, $2, $3)`
		prep, err := r.db.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer prep.Close()
		if _, err := prep.ExecContext(ctx, input.PostID, input.UserID, input.LikeType); err != nil {
			return err
		}
	}
	return nil
}

func (r *PostRepository) checkPostLikeExists(ctx context.Context, postID, userID uint) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM posts_likes WHERE post_id = $1 AND user_id = $2)",
		postID, userID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *PostRepository) getPostLikeType(ctx context.Context, postID, userID uint) (int, error) {
	var likeType int
	err := r.db.QueryRowContext(ctx, "SELECT type FROM posts_likes WHERE post_id = $1 AND user_id = $2",
		postID, userID).Scan(&likeType)
	if err != nil {
		return 0, err
	}
	return likeType, nil
}

func (r *PostRepository) getLikesAndDislikes(postID uint) (models.Vote, error) {
	query := `
	SELECT 
		COUNT(CASE WHEN type = -1 THEN 1 END) AS dislikes,
		COUNT(CASE WHEN type = 1 THEN 1 END) AS likes
	FROM posts_likes
	WHERE post_id = $1;
	`

	var dislikes, likes uint
	err := r.db.QueryRow(query, postID).Scan(&dislikes, &likes)
	if err != nil {
		return models.Vote{}, err
	}
	vote := models.Vote{
		Likes:    likes,
		Dislikes: dislikes,
	}
	return vote, nil
}
