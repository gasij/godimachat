# internal/handlers

**ЗАДАНИЕ.md:** Шаги 20 – 26, 28

## api.go — порядок

| # | Что |
|---|-----|
| 1 | `Handler` + `New` |
| 2 | `writeJSON`, `setTokenCookie` |
| 3 | `Register` |
| 4 | `Login` |
| 5 | `Me` |
| 6 | `GetMessages` |
| 7 | `SendMessage` |
| 8 | `WebSocket` |
| 9 | `Routes` |

## pages.go

| Метод | Шаблон | URL |
|-------|--------|-----|
| `Login` | login.html | /login |
| `Register` | register.html | /register |
| `Chat` | chat.html | /chat |

## Проверка через curl

Примеры в [ЗАДАНИЕ.md](../../ЗАДАНИЕ.md) — шаги 21–23.

## Контракт с фронтендом

Обязательно прочитай раздел «Контракт API» в [ЗАДАНИЕ.md](../../ЗАДАНИЕ.md).
