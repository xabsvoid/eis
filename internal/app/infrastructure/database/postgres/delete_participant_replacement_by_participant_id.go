package postgres

import "context"

const deleteParticipantReplacementByParticipantIDSQL = `update participants
set replacement_person_id=null,
    updated_at=now()
where id=$1`

func (db *Postgres) DeleteParticipantReplacementByParticipantID(ctx context.Context, participantID int64) error {
	_, err := db.Exec(ctx, deleteParticipantReplacementByParticipantIDSQL, participantID)
	if err != nil {
		return err
	}

	return nil
}
