package app

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/Anton-Kraev/vk-test-assignment/pinger/internal/http/client"
	"github.com/Anton-Kraev/vk-test-assignment/pinger/internal/pinger"
)

const (
	pingInterval   = 5 * time.Second
	backendURL     = "http://backend:8080/containers"
	requestTimeout = 5 * time.Second
)

func Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	backendClient := client.New(backendURL, requestTimeout)

	containerPinger := pinger.New(backendClient)
	containerPinger.Start(ctx, pingInterval)
}
