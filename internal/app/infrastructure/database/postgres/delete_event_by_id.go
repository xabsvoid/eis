package postgres

import "context"

const (
	deleteEventByIDSQL = `delete from events
where id=$1`
	deleteEventParticipantsByEventIDSQL = `delete from participants
where event_id=$1`
)

func (db *Postgres) DeleteEventByID(ctx context.Context, eventID int64) error {
	_, err := db.Exec(ctx, deleteEventByIDSQL, eventID)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, deleteEventParticipantsByEventIDSQL, eventID)
	if err != nil {
		return err
	}

	return nil
}
