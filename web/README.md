# web/ — готовый фронтенд

**Не переписывай**, пока бэкенд не заработает. Потом можно улучшать.

## Что уже есть

| Файл | Что делает |
|------|------------|
| `templates/login.html` | Форма входа |
| `templates/register.html` | Форма регистрации |
| `templates/chat.html` | Интерфейс чата |
| `static/css/style.css` | Стили |
| `static/js/auth.js` | fetch → /api/login, /api/register |
| `static/js/chat.js` | fetch + WebSocket |

## Какие запросы шлёт фронтенд

Читай эти файлы **перед** написанием handlers:

### auth.js
- `POST /api/register` → `{username, email, password}`
- `POST /api/login` → `{email, password}`
- Ожидает ответ `{token, user}`

### chat.js
- `GET /api/me` + `Authorization: Bearer token`
- `GET /api/messages`
- `POST /api/messages` → `{content}`
- `WebSocket ws://host/api/ws`
- Ожидает WS: `{type: "message", data: {...}}`

Полный контракт — в [ЗАДАНИЕ.md](../ЗАДАНИЕ.md#контракт-api-шпаргалка).
