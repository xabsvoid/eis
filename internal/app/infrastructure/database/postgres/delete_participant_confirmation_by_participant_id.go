package postgres

import "context"

const deleteParticipantConfirmationByParticipantIDSQL = `update participants
set confirmed=false,
    updated_at=now()
where id=$1`

func (db *Postgres) DeleteParticipantConfirmationByParticipantID(ctx context.Context, participantID int64) error {
	_, err := db.Exec(ctx, deleteParticipantConfirmationByParticipantIDSQL, participantID)
	if err != nil {
		return err
	}

	return nil
}
