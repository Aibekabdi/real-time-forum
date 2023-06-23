package repository

import (
	"context"
	"database/sql"
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
			LEFT JOIN tags t ON pt.tag_id = t.id
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
				Author: models.User{Id: authorID, Nickname: author},
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
