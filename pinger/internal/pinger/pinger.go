package pinger

import (
	"context"
	"log"
	"time"

	"github.com/Anton-Kraev/vk-test-assignment/pinger/internal/models"
)

type backendClient interface {
	GetContainers() ([]models.Container, error)
	UpdateContainers(pings []models.Ping) error
}

type Pinger struct {
	client backendClient
}

func New(client backendClient) Pinger {
	return Pinger{client: client}
}

func (p Pinger) Start(ctx context.Context, pingInterval time.Duration) {
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("[INFO] Stopping processing events")

			return
		case <-ticker.C:
		}

		containers, err := p.client.GetContainers()
		if err != nil {
			log.Println("[ERROR] Error getting containers:", err)

			continue
		}

		pings := pingContainers(containers)

		if err = p.client.UpdateContainers(pings); err != nil {
			log.Println("[ERROR] Error updating containers", err)

			continue
		}

		log.Println("[INFO] Pinging containers done")
	}
}
