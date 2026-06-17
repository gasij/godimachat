# internal/database

**ЗАДАНИЕ.md:** Шаг 5

## Файл: `database.go`

Одна функция:

```go
func Connect(ctx context.Context, dsn string) (*pgxpool.Pool, error)
```

1. `pgxpool.New(ctx, dsn)`
2. `pool.Ping(ctx)`
3. Вернуть pool или ошибку

## Документация

- https://github.com/jackc/pgx/wiki/Getting-started-with-pgx

## Проверка

Временный main: подключись и выведи «БД подключена!»
