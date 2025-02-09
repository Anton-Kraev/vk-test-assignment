.PHONY:
.SILENT:

all: postgres migrate backend

postgres:
	docker-compose up --build -d postgres

backend:
	docker-compose up --build -d backend

migrate:
	./backend/scripts/migration.sh
