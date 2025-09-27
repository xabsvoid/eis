package postgres

import "context"

const createParticipantCheckInSQL = `update participants
set check_in=true,
    updated_at=now()
where id=$1`

func (db *Postgres) CreateParticipantCheckIn(ctx context.Context, participantID int64) error {
	_, err := db.Exec(ctx, createParticipantCheckInSQL, participantID)
	if err != nil {
		return err
	}

	return nil
}
