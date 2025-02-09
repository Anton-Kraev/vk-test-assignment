package handler

import (
	"net"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/Anton-Kraev/vk-test-assignment/internal/models"
)

type container struct {
	ID               int        `json:"id"`
	IP               net.IP     `json:"ip"`
	LastPingAttempt  *time.Time `json:"last_ping_attempt,omitempty"`
	LastSuccefulPing *time.Time `json:"last_successful_ping,omitempty"`
	ResponseTimeMS   int        `json:"response_time_ms,omitempty"`
}

func containerFromDomain(c models.Container) container {
	cont := container{
		ID:               c.ID,
		IP:               c.IP,
		LastPingAttempt:  &c.LastPingAttempt,
		LastSuccefulPing: &c.LastSuccefulPing,
		ResponseTimeMS:   c.ResponseTimeMS,
	}

	if c.LastPingAttempt.IsZero() {
		cont.LastPingAttempt = nil
	}

	if c.LastSuccefulPing.IsZero() {
		cont.LastSuccefulPing = nil
	}

	return cont
}

type getContainersResponse struct {
	Containers []container `json:"containers"`
}

func (h Handler) GetContainers(c echo.Context) error {
	containers, err := h.repo.GetContainers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := getContainersResponse{Containers: make([]container, 0)}

	for _, cont := range containers {
		resp.Containers = append(resp.Containers, containerFromDomain(cont))
	}

	return c.JSON(http.StatusOK, resp)
}
