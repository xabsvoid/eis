package repository

import "context"

type CreateParticipantConfirmation interface {
	CreateParticipantConfirmation(ctx context.Context, participantID int64) error
}
