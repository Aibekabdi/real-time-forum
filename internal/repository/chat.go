package repository

import (
	"context"
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type ChatRepository struct {
	db *sqlx.DB
}

func newChatRepository(db *sqlx.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) GetMessages(ctx context.Context, senderID, receiverID, lastMessageID, limit uint) ([]models.Message, error) {
	query := `
	SELECT id, sender_id, receiver_id, content, sent_at, read
	FROM messages
	WHERE (sender_id = $1 and receiver_id = $2)
	OR (sender_id = $2 and receiver_id = $1)
	AND 
	CASE WHEN $3 = 0 THEN true ELSE id < $3 END
	ORDER BY id DESC LIMIT $4;
`
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := prep.QueryContext(ctx, senderID, receiverID, lastMessageID, limit)
	if err != nil {
		return nil, err
	}

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, message.SenderID, message.ReceiverID, message.Content, message.SentAt, message.Read); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	prep.Close()
	return messages, nil
}

func (r *ChatRepository) Create(ctx context.Context, message *models.Message) (uint, error) {
	var (
		id  uint
		err error
	)
	query := `
		INSERT INTO messages (sender_id, receiver_id, content, sent_at, read) VALUES ($1, $2, $3, $5, $6) RETURNING id;
	`

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, nil
	}
	if err = prep.QueryRowContext(ctx, message.SenderID, message.ReceiverID, message.Content, message.SentAt, message.Read).Scan(&id); err != nil {
		return 0, nil
	}

	return id, nil
}
