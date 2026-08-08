package handlers

import (
	"encoding/json"
	"godima/internal/auth"
	"godima/internal/hub"
	"godima/internal/middleware"
	"godima/internal/models"
	"godima/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

type Handler struct {
	users     *repository.UserRepository
	messages  *repository.MessageRepository
	hub       *hub.Hub
	jwtSecret string
}

func New(users *repository.UserRepository, messages *repository.MessageRepository,
	h *hub.Hub, jwtSecret string) *Handler {
	return &Handler{users: users, messages: messages, hub: h, jwtSecret: jwtSecret}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func setTokenCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400,
	})
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "неверный JSON"})
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "заполни все поля"})
		return
	}

	hash, err := auth.HashPassword(req.Password)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "ошибка сервера"})
		return
	}

	user, err := h.users.Create(r.Context(), req.Username, req.Email, hash)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "ошибка сервера"})
		return
	}

	token, _ := auth.GenerateToken(user.ID, user.Username, h.jwtSecret)
	setTokenCookie(w, token)
	writeJSON(w, http.StatusCreated, map[string]any{"token": token, "user": user})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "неверный JSON"})
		return
	}

	user, err := h.users.FindByEmail(r.Context(), req.Email)
	if err != nil || user == nil || !auth.CheckPassword(user.Password, req.Password) {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "неверный email или пароль"})
		return
	}

	token, _ := auth.GenerateToken(user.ID, user.Username, h.jwtSecret)
	setTokenCookie(w, token)
	writeJSON(w, http.StatusOK, map[string]any{
		"token": token,
		"user":  map[string]any{"id": user.ID, "username": user.Username, "email": user.Email},
	})
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	claims := middleware.GetUser(r)
	writeJSON(w, http.StatusOK, map[string]any{
		"id":       claims.UserID,
		"username": claims.Username,
	})
}

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := h.messages.GetRecent(r.Context(), 50)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "ошибка сервера"})
		return
	}
	if messages == nil {
		messages = []models.Message{}
	}
	writeJSON(w, http.StatusOK, messages)
}

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	claims := middleware.GetUser(r)

	var req models.SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "неверный JSON"})
		return
	}
	if req.Content == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "пустое сообщение"})
		return
	}

	msg, err := h.messages.Create(r.Context(), claims.UserID, req.Content)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "ошибка сервера"})
		return
	}

	msg.Username = claims.Username
	h.hub.BroadcastMessage(*msg)

	writeJSON(w, http.StatusCreated, msg)
}

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/api/register", h.Register)
	r.Post("/api/login", h.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(h.jwtSecret))
		r.Get("/api/me", h.Me)
		r.Get("/api/messages", h.GetMessages)
		r.Post("/api/messages", h.SendMessage)
		r.Get("/api/ws", h.WebSocket)
	})

	return r
}

func (h *Handler) WebSocket(w http.ResponseWriter, r *http.Request) {
	claims := middleware.GetUser(r)

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &hub.Client{
		Hub: h.hub, Conn: conn,
		Send:     make(chan []byte, 256),
		Username: claims.Username,
	}

	h.hub.Register(client)
	go client.WritePump()
	go client.ReadPump()
}
