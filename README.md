# Godima Chat —

Онлайн-чат с авторизацией. **Фронтенд готов**, бэкенд пишешь ты.

## С чего начать

1. Прочитай этот файл целиком (хотя бы бегло)
2. Открой **[ЗАДАНИЕ.md](./ЗАДАНИЕ.md)** — там пошаговая инструкция с примерами кода
3. Пиши **по одному методу за раз**, проверяй, потом переходи к следующему
4. В каждой папке `internal/*` есть краткий `README.md` — напоминание, что делать в этой папке

## Что уже готово

| Файл / папка | Назначение |
|--------------|------------|
| `web/templates/` | HTML: login, register, chat |
| `web/static/` | CSS и JavaScript |
| `migrations/001_initial.sql` | Таблицы `users` и `messages` |
| `docker-compose.yml` | PostgreSQL в Docker |
| `.env.example` | Настройки подключения |

## Быстрый старт

```bash
cp .env.example .env
make db-up
```

Когда напишешь `cmd/server/main.go`:

```bash
go mod tidy
make run
```

Открой http://localhost:8080

## Структура (что ты пишешь)

```
cmd/server/main.go
internal/config/config.go
internal/database/database.go
internal/models/models.go
internal/auth/auth.go
internal/repository/user.go
internal/repository/message.go
internal/middleware/auth.go
internal/handlers/api.go
internal/handlers/pages.go
internal/hub/hub.go
internal/hub/client.go
```

## Документация для ученика

| Файл | Для чего |
|------|----------|
| **[ЗАДАНИЕ.md](./ЗАДАНИЕ.md)** | Полная пошаговая инструкция с примерами кода (30 шагов) |
| **[ШПАРГАЛКА.md](./ШПАРГАЛКА.md)** | Краткий чеклист: что писать и в каком порядке |
| `internal/*/README.md` | Напоминание по каждой папке |
