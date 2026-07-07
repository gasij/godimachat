package auth

import "testing"

func TestPassword(t *testing.T) {
	hash, err := HashPassword("secret123")
	if err != nil {
		t.Fatal(err)
	}
	if !CheckPassword(hash, "secret123") {
		t.Error("пароль должен совпадать")
	}
	if CheckPassword(hash, "wrong") {
		t.Error("неверный пароль не должен проходить")
	}
}

func TestJWT(t *testing.T) {
	token, err := GenerateToken(1, "ivan", "test-secret")
	if err != nil {
		t.Fatal(err)
	}
	claims, err := ParseToken(token, "test-secret")
	if err != nil {
		t.Fatal(err)
	}
	if claims.UserID != 1 || claims.Username != "ivan" {
		t.Error("неверные claims")
	}
}
