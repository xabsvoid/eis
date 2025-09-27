package postgres

import "context"

const deleteParticipantCheckInByParticipantIDSQL = `update participants
set check_in=false,
    updated_at=now()
where id=$1`

func (db *Postgres) DeleteParticipantCheckInByParticipantID(ctx context.Context, participantID int64) error {
	_, err := db.Exec(ctx, deleteParticipantCheckInByParticipantIDSQL, participantID)
	if err != nil {
		return err
	}

	return nil
}
