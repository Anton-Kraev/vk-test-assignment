package models

import (
	"net"
	"time"
)

type Container struct {
	ID               int        `json:"id"`
	IP               net.IP     `json:"ip"`
	LastPingAttempt  *time.Time `json:"last_ping_attempt,omitempty"`
	LastSuccefulPing *time.Time `json:"last_successful_ping,omitempty"`
	ResponseTimeMS   int        `json:"response_time_ms,omitempty"`
}

type Ping struct {
	ContainerID    int       `json:"container_id"`
	Success        bool      `json:"success"`
	AttemptTime    time.Time `json:"attempt_time"`
	ResponseTimeMS int       `json:"response_time_ms"`
}
