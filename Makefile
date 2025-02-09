.PHONY:
.SILENT:

postgres:
	docker-compose up --build -d postgres

migrate:
	./backend/scripts/migration.sh

backend:
	docker-compose up --build -d backend

frontend:
	docker-compose up --build -d frontend

pinger:
	docker-compose up --build -d pinger

app:
	docker-compose up --build -d backend frontend pinger
