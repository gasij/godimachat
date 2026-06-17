# internal/config

**ЗАДАНИЕ.md:** Шаги 1.1 – 1.5

## Порядок написания

| # | Что | Функция |
|---|-----|---------|
| 1 | Структура | `Config` |
| 2 | Хелпер | `getEnv(key, fallback)` |
| 3 | Метод | `DSN() string` |
| 4 | Загрузка | `Load() (*Config, error)` |

## Откуда брать поля

Смотри `.env.example`:
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`
- `SERVER_PORT`, `JWT_SECRET`

## Документация

- https://pkg.go.dev/os#Getenv
