# internal/models

**ЗАДАНИЕ.md:** Шаг 7

## Файл: `models.go`

Пиши структуры **до** repository и handlers.

| Структура | Назначение |
|-----------|------------|
| `User` | Пользователь из БД |
| `Message` | Сообщение из БД |
| `RegisterRequest` | Тело POST /api/register |
| `LoginRequest` | Тело POST /api/login |
| `SendMessageRequest` | Тело POST /api/messages |

## Важно

- `Password` — тег `json:"-"` (не отдавать в API)
- Поля `json:"..."` должны совпадать с тем, что шлёт/ждёт фронтенд

## Смотри

- `web/static/js/auth.js`
- `web/static/js/chat.js`
- `migrations/001_initial.sql`
