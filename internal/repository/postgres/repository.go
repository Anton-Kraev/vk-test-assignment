package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Anton-Kraev/vk-test-assignment/internal/models"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return Repository{db: db}
}

func (r Repository) GetContainers(ctx context.Context) ([]models.Container, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) UpdateContainers(ctx context.Context, containers []models.Container) error {
	//TODO implement me
	panic("implement me")
}
