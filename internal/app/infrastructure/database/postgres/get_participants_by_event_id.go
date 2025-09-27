package postgres

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

const getParticipantsByEventIDSQL = `select id, event_id, person_id, replacement_person_id, check_in, confirmed, updated_at, created_at
from participants
where event_id=$1`

func (db *Postgres) GetParticipantsByEventID(ctx context.Context, eventID int64) ([]entity.Participant, error) {
	var participants []entity.Participant

	rows, err := db.Query(ctx, getParticipantsByEventIDSQL, eventID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var participant entity.Participant
		var replacementPersonID *int64
		err = rows.Scan(
			&participant.ID,
			&participant.EventID,
			&participant.PersonID,
			&replacementPersonID,
			&participant.CheckIn.Exist,
			&participant.Confirmation.Exist,
			&participant.UpdatedAt,
			&participant.CreatedAt)
		if err != nil {
			return nil, err
		}

		if replacementPersonID != nil {
			participant.Replacement.Exist = true
			participant.Replacement.PersonID = *replacementPersonID
		}

		participants = append(participants, participant)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return participants, nil
}
