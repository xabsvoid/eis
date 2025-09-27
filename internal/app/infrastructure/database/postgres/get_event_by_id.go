package postgres

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

const getEventByIDSQL = `select id, title, date, updated_at, created_at
from events
where id=$1`

func (db *Postgres) GetEventByID(ctx context.Context, id int64) (entity.Event, error) {
	row := db.QueryRow(ctx, getEventByIDSQL, id)

	var event entity.Event
	err := row.Scan(
		&event.ID,
		&event.Title,
		&event.Date,
		&event.UpdatedAt,
		&event.CreatedAt)
	if err != nil {
		return event, err
	}

	return event, nil
}
