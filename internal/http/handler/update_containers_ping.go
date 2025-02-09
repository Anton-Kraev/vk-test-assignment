package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/Anton-Kraev/vk-test-assignment/internal/models"
	repo "github.com/Anton-Kraev/vk-test-assignment/internal/repository"
)

type ping struct {
	ContainerID    int       `json:"container_id" validate:"required"`
	Success        bool      `json:"success" validate:"required"`
	AttemptTime    time.Time `json:"attempt_time" validate:"required"`
	ResponseTimeMS int       `json:"response_time_ms"`
}

func (p ping) toDomain() models.Container {
	cont := models.Container{ID: p.ContainerID, LastPingAttempt: p.AttemptTime}

	if p.Success {
		cont.LastSuccefulPing = p.AttemptTime
		cont.ResponseTimeMS = p.ResponseTimeMS
	}

	return cont
}

type updateContainersPingRequest struct {
	Pings []ping `json:"pings" validate:"required"`
}

func (h Handler) UpdateContainersPing(c echo.Context) error {
	var req updateContainersPingRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var conts []models.Container

	for _, p := range req.Pings {
		conts = append(conts, p.toDomain())
	}

	err := h.repo.UpdateContainers(c.Request().Context(), conts)

	switch {
	case errors.As(err, &repo.ErrContainerNotFound{}):
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	case err != nil:
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	default:
		return c.NoContent(http.StatusOK)
	}
}
