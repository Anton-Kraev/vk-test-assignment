package client

import (
	"fmt"

	"github.com/Anton-Kraev/vk-test-assignment/pinger/internal/models"
)

type updateContainerRequest struct {
	Pings []models.Ping `json:"pings"`
}

func (c Client) UpdateContainers(pings []models.Ping) error {
	const op = "client.Client.UpdateContainers"

	resp, err := c.http.R().
		SetBody(&updateContainerRequest{Pings: pings}).
		Patch("")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if resp.IsError() {
		return fmt.Errorf("%s: failed to update containers: %s", op, resp.Status())
	}

	return nil
}
