package postgres

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

const getEventsSQL = `select id, title, date, updated_at, created_at
from events`

func (db *Postgres) GetEvents(ctx context.Context) ([]entity.Event, error) {
	var events []entity.Event

	rows, err := db.Query(ctx, getEventsSQL)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event entity.Event
		err = rows.Scan(
			&event.ID,
			&event.Title,
			&event.Date,
			&event.UpdatedAt,
			&event.CreatedAt)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return events, nil
}
