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
/*
фронт в chat.js ждет у сообщений поле username а твоя модель Message не содержит такого поля. 
Поэтому нужно добавить поле Username в модель Message, чтобы фронт мог получить имя пользователя вместе с сообщением.
модно добавить отдельную response модель  например MessageResponse что бы api возращал id , user_id username ,content,
created_at. дальше идешь чисто по шпаргалке ч
*/
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
