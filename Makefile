.PHONY:
.SILENT:

all: postgres migrate app

postgres:
	docker-compose up --build -d postgres

app:
	go run ./cmd/app/main.go

migrate:
	./scripts/migration.sh
