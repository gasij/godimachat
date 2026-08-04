package middleware

import (
	"context"
	"godima/internal/auth"
	"net/http"
	"strings"
)

type contextKey string

const UserContextKey contextKey = "user"

func extractToken(r *http.Request) string {
	// Из заголовка
	h := r.Header.Get("Authorization")
	if strings.HasPrefix(h, "Bearer ") {
		return strings.TrimPrefix(h, "Bearer ")
	}
	// Из cookie (браузер шлёт автоматически)
	if c, err := r.Cookie("token"); err == nil {
		return c.Value
	}
	return ""
}

func Auth(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := extractToken(r)
			if token == "" {
				http.Error(w, `{"error":"требуется авторизация"}`, http.StatusUnauthorized)
				return
			}

			claims, err := auth.ParseToken(token, jwtSecret)
			if err != nil {
				http.Error(w, `{"error":"невалидный токен"}`, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUser(r *http.Request) *auth.Claims {
	claims, _ := r.Context().Value(UserContextKey).(*auth.Claims)
	return claims
}
