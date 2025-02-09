package app

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/Anton-Kraev/vk-test-assignment/internal/http/handler"
	"github.com/Anton-Kraev/vk-test-assignment/internal/http/server"
	"github.com/Anton-Kraev/vk-test-assignment/internal/repository/postgres"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening postgres connection pool: %v", err)
	}

	repo := postgres.NewRepository(db)
	hndl := handler.New(repo)

	if err = server.Start(hndl); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
