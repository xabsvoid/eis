package postgres

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

const createEventSQL = `insert into events (title, date)
values ($1, $2)
returning id, updated_at, created_at`

func (db *Postgres) CreateEvent(ctx context.Context, event entity.Event) (entity.Event, error) {
	row := db.QueryRow(ctx, createEventSQL, event.Title, event.Date)

	err := row.Scan(&event.ID, &event.UpdatedAt, &event.CreatedAt)
	if err != nil {
		return event, err
	}

	return event, nil
}
