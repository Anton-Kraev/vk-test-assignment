package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	"github.com/georgysavva/scany/v2/pgxscan"

	"github.com/Anton-Kraev/vk-test-assignment/backend/internal/models"
)

type container struct {
	ID                 int           `db:"id"`
	IP                 net.IP        `db:"ip"`
	LastPingAttempt    sql.NullTime  `db:"last_ping_attempt"`
	LastSuccessfulPing sql.NullTime  `db:"last_successful_ping"`
	ResponseTimeMS     sql.Null[int] `db:"response_time_ms"`
}

func (c container) toDomain() models.Container {
	return models.Container{
		ID:               c.ID,
		IP:               c.IP,
		LastPingAttempt:  c.LastPingAttempt.Time,
		LastSuccefulPing: c.LastSuccessfulPing.Time,
		ResponseTimeMS:   c.ResponseTimeMS.V,
	}
}

func (r Repository) GetContainers(ctx context.Context) ([]models.Container, error) {
	const (
		op    = "postgres.Repository.GetContainers"
		query = `
			SELECT id, ip, last_ping_attempt, last_successful_ping, response_time_ms
			FROM container
		`
	)

	var records []container
	if err := pgxscan.Select(ctx, r.db, &records, query); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	containers := make([]models.Container, len(records))

	for i, rec := range records {
		containers[i] = rec.toDomain()
	}

	return containers, nil
}
