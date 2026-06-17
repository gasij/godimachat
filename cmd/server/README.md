# cmd/server

## Файл: `main.go`

**Когда писать:** в самом конце, после всех пакетов (Шаг 30 в [ЗАДАНИЕ.md](../../ЗАДАНИЕ.md)).

## Что делает main.go

1. `godotenv.Load()` — читает `.env`
2. `config.Load()` — настройки
3. `database.Connect()` — PostgreSQL
4. Создаёт репозитории, hub, handlers
5. Настраивает роутер (статика + HTML + API)
6. Запускает сервер + graceful shutdown

## Маршруты

| URL | Откуда |
|-----|--------|
| `/static/*` | `web/static/` |
| `/login`, `/register`, `/chat` | `pagesHandler` |
| `/api/*` | `apiHandler.Routes()` |

Полный пример — в [ЗАДАНИЕ.md, Шаг 30](../../ЗАДАНИЕ.md#шаг-30-maingo).
