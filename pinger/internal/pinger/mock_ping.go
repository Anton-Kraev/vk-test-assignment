package pinger

import (
	"math/rand"
	"time"

	"github.com/Anton-Kraev/vk-test-assignment/pinger/internal/models"
)

func pingContainers(containers []models.Container) []models.Ping {
	pings := make([]models.Ping, len(containers))

	for i, cont := range containers {
		pings[i] = models.Ping{
			ContainerID:    cont.ID,
			Success:        rand.Float32() > 0.2,
			AttemptTime:    time.Now(),
			ResponseTimeMS: rand.Intn(100),
		}
	}

	return pings
}
