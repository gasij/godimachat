# internal/middleware

**ЗАДАНИЕ.md:** Шаг 18

## Файл: `auth.go`

| # | Функция |
|---|---------|
| 1 | `extractToken(r)` — из заголовка или cookie |
| 2 | `Auth(jwtSecret)` — middleware |
| 3 | `GetUser(r)` — достать claims из контекста |

## Откуда берётся токен

1. Заголовок: `Authorization: Bearer eyJ...` (так шлёт `chat.js`)
2. Cookie: `token=eyJ...` (ставится при login/register)

## Документация

- https://github.com/go-chi/chi#middleware-handlers
- https://pkg.go.dev/context#WithValue
