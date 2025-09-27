package postgres

import "context"

const createParticipantConfirmationSQL = `update participants
set confirmed=true,
    updated_at=now()
where id=$1`

func (db *Postgres) CreateParticipantConfirmation(ctx context.Context, participantID int64) error {
	_, err := db.Exec(ctx, createParticipantConfirmationSQL, participantID)
	if err != nil {
		return err
	}

	return nil
}
