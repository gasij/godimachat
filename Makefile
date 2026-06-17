.PHONY: help db-up db-down db-logs run build test tidy

help:
	@echo "Доступные команды:"
	@echo "  make db-up    — запустить PostgreSQL в Docker"
	@echo "  make db-down  — остановить PostgreSQL"
	@echo "  make db-logs  — логи базы данных"
	@echo "  make run      — запустить сервер"
	@echo "  make build    — собрать бинарник"
	@echo "  make tidy     — обновить go.mod"

db-up:
	docker compose up -d

db-down:
	docker compose down

db-logs:
	docker compose logs -f postgres

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

tidy:
	go mod tidy

test:
	go test ./...
