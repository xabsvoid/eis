package postgres

import "context"

const deleteParticipantByIDSQL = `delete from participants
where id=$1`

func (db *Postgres) DeleteParticipantByID(ctx context.Context, id int64) error {
	_, err := db.Exec(ctx, deleteParticipantByIDSQL, id)
	if err != nil {
		return err
	}

	return nil
}
