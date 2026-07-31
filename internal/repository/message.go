package repository

import (
	"context"
	"fmt"
	"godima/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MessageRepository struct {
	db *pgxpool.Pool
}

func NewMessageRepository(db *pgxpool.Pool) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(ctx context.Context, userID int, content string) (*models.Message, error) {
	query := `
        INSERT INTO messages (user_id, content)
        VALUES ($1, $2)
        RETURNING id, user_id, content, created_at
    `
	var msg models.Message
	err := r.db.QueryRow(ctx, query, userID, content).Scan(
		&msg.ID, &msg.UserID, &msg.Content, &msg.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("создание сообщения: %w", err)
	}
	return &msg, nil
}

func (r *MessageRepository) GetRecent(ctx context.Context, limit int) ([]models.Message, error) {
	query := `
        SELECT m.id, m.user_id, u.username, m.content, m.created_at
        FROM messages m
        JOIN users u ON u.id = m.user_id
        ORDER BY m.created_at DESC
        LIMIT $1
    `
	rows, err := r.db.Query(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("получение сообщений: %w", err)
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.UserID, &msg.Username, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return messages, nil
}
