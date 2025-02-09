package client

import (
	"fmt"

	"github.com/Anton-Kraev/vk-test-assignment/pinger/internal/models"
)

type getContainersResponse struct {
	Containers []models.Container `json:"containers"`
}

func (c Client) GetContainers() ([]models.Container, error) {
	const op = "client.Client.GetContainers"

	var containers getContainersResponse

	resp, err := c.http.R().SetResult(&containers).Get("")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("%s: failed to fetch containers: %s", op, resp.Status())
	}

	return containers.Containers, nil
}
