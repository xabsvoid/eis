package repository

import "context"

type DeleteParticipantReplacementByParticipantID interface {
	DeleteParticipantReplacementByParticipantID(ctx context.Context, participantID int64) error
}
