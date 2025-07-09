package repository

import "context"

type DeleteParticipantConfirmationByParticipantID interface {
	DeleteParticipantConfirmationByParticipantID(ctx context.Context, participantID int64) error
}
