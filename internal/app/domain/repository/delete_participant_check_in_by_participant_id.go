package repository

import "context"

type DeleteParticipantCheckInByParticipantID interface {
	DeleteParticipantCheckInByParticipantID(ctx context.Context, participantID int64) error
}
