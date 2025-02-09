package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/Anton-Kraev/vk-test-assignment/backend/internal/models"
	"github.com/Anton-Kraev/vk-test-assignment/backend/internal/repository"
)

func (r Repository) UpdateContainers(ctx context.Context, containers []models.Container) error {
	const op = "postgres.Repository.UpdateContainers"

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	for _, cont := range containers {
		if err = updateContainer(ctx, tx, cont); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func updateContainer(ctx context.Context, tx pgx.Tx, cont models.Container) error {
	const op = "postgres.updateContainer"

	var (
		query = "UPDATE container SET last_ping_attempt = $2, %s WHERE id = $1"
		args  = []any{cont.ID, cont.LastPingAttempt}
	)

	if cont.ResponseTimeMS != 0 {
		query = fmt.Sprintf(query, "last_successful_ping = $3, response_time_ms = $4")
		args = append(args, cont.LastSuccefulPing, cont.ResponseTimeMS)
	} else {
		query = fmt.Sprintf(query, "response_time_ms = NULL")
	}

	stats, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if stats.RowsAffected() == 0 {
		return repository.NewErrContainerNotFound(op, cont.ID)
	}

	return nil
}
