.PHONY:
.SILENT:

run:
	go run ./cmd/app/main.go

migrate:
	.scripts/migration.sh
