package postgres

import "context"

const createParticipantReplacementSQL = `update participants
set replacement_person_id=$2,
    updated_at=now()
where id=$1`

func (db *Postgres) CreateParticipantReplacement(ctx context.Context, participantID, personID int64) error {
	_, err := db.Exec(ctx, createParticipantReplacementSQL, participantID, personID)
	if err != nil {
		return err
	}

	return nil
}
