# internal/auth

**ЗАДАНИЕ.md:** Шаги 8 – 9

## Порядок написания

| # | Функция | Шаг |
|---|---------|-----|
| 1 | `HashPassword` | 8.1 |
| 2 | `CheckPassword` | 8.2 |
| 3 | `Claims` (структура) | 9.1 |
| 4 | `GenerateToken` | 9.2 |
| 5 | `ParseToken` | 9.3 |

## Тесты

Создай `auth_test.go` после шагов 8 и 9:

```bash
go test ./internal/auth/...
```

## Документация

- bcrypt: https://pkg.go.dev/golang.org/x/crypto/bcrypt
- JWT: https://github.com/golang-jwt/jwt
