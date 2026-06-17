# internal/hub

**ЗАДАНИЕ.md:** Шаги 27 – 28

## hub.go — порядок

| # | Что |
|---|-----|
| 1 | `Client`, `Hub` структуры |
| 2 | `New()` |
| 3 | `Run()` — в горутине! |
| 4 | `Register(client)` |
| 5 | `BroadcastMessage(msg)` |

## client.go

| # | Метод |
|---|-------|
| 1 | `ReadPump()` |
| 2 | `WritePump()` |

## Формат WebSocket

Фронтенд (`chat.js`) ждёт:

```json
{"type": "message", "data": {"id": 1, "user_id": 1, "username": "ivan", "content": "...", "created_at": "..."}}
```

## Документация

- https://github.com/gorilla/websocket
- Пример чата: https://github.com/gorilla/websocket/tree/main/examples/chat
