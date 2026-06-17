# internal/repository

**ЗАДАНИЕ.md:** Шаги 13 – 15

## user.go — порядок

| # | Метод |
|---|-------|
| 1 | `NewUserRepository` |
| 2 | `Create` |
| 3 | `FindByEmail` |
| 4 | `FindByID` |

## message.go — порядок

| # | Метод |
|---|-------|
| 1 | `NewMessageRepository` |
| 2 | `Create` |
| 3 | `GetRecent` |

## SQL

Схема таблиц — `migrations/001_initial.sql`

Плейсхолдеры в pgx: `$1`, `$2`, `$3` (не `?` как в MySQL!)

## Документация

- https://github.com/jackc/pgx#query

## Проверка

Временный main: создай пользователя, найди по email, создай сообщение.
