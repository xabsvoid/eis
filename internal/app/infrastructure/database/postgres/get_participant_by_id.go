package postgres

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

const getParticipantByIDSQL = `select id, event_id, person_id, replacement_person_id, check_in, confirmed, updated_at, created_at
from participants
where id=$1`

func (db *Postgres) GetParticipantByID(ctx context.Context, id int64) (entity.Participant, error) {
	row := db.QueryRow(ctx, getParticipantByIDSQL, id)

	var participant entity.Participant
	var replacementPersonID *int64
	err := row.Scan(
		&participant.ID,
		&participant.EventID,
		&participant.PersonID,
		&replacementPersonID,
		&participant.CheckIn.Exist,
		&participant.Confirmation.Exist,
		&participant.UpdatedAt,
		&participant.CreatedAt)
	if err != nil {
		return participant, err
	}

	if replacementPersonID != nil {
		participant.Replacement.Exist = true
		participant.Replacement.PersonID = *replacementPersonID
	}

	return participant, nil
}
