package handler

import (
	"context"

	"github.com/Anton-Kraev/vk-test-assignment/internal/models"
)

type repository interface {
	GetContainers(ctx context.Context) ([]models.Container, error)
	UpdateContainers(ctx context.Context, containers []models.Container) error
}

type Handler struct {
	repo repository
}

func New(repo repository) Handler {
	return Handler{repo: repo}
}
