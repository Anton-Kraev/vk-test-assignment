package models

import (
	"net"
	"time"
)

type Container struct {
	ID               int
	IP               net.IP
	LastPingAttempt  time.Time
	LastSuccefulPing time.Time
	ResponseTimeMS   int
}
