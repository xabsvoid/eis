package repository

import "context"

type CreateParticipantCheckIn interface {
	CreateParticipantCheckIn(ctx context.Context, participantID int64) error
}
