package postgres

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

const createParticipantSQL = `insert into participants (event_id, person_id)
values ($1, $2)
returning id, updated_at, created_at`

func (db *Postgres) CreateParticipant(ctx context.Context, participant entity.Participant) (entity.Participant, error) {
	row := db.QueryRow(ctx, createParticipantSQL, participant.EventID, participant.PersonID)

	err := row.Scan(&participant.ID, &participant.UpdatedAt, &participant.CreatedAt)
	if err != nil {
		return participant, err
	}

	return participant, nil
}
