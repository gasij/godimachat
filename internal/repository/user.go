package repository

import (
	"context"
	"errors"
	"fmt"
	"godima/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, username, email, passwordHash string) (*models.User, error) {
	query := `
        INSERT INTO users (username, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, username, email, created_at
    `
	var user models.User
	err := r.db.QueryRow(ctx, query, username, email, passwordHash).Scan(
		&user.ID, &user.Username, &user.Email, &user.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("создание пользователя: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, username, email, password, created_at FROM users WHERE email = $1`

	var user models.User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // пользователь не найден — это не ошибка
		}
		return nil, fmt.Errorf("поиск пользователя: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT id, username, email, created_at FROM users WHERE id = $1`

	var user models.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("поиск пользователя: %w", err)
	}
	return &user, nil
}
