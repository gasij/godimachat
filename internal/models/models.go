package models

import (
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Created_at time.Time `json:"created_at"`
}

type Message struct {
	ID         int       `json:"id"`
	User_id    int       `json:"user_id"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SendMessageRequest struct {
	Content string `json:"content"`
}
