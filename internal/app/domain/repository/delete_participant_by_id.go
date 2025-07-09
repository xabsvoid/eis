package repository

import "context"

type DeleteParticipantByID interface {
	DeleteParticipantByID(ctx context.Context, id int64) error
}
