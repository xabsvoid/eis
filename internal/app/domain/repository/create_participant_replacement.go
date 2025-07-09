package repository

import "context"

type CreateParticipantReplacement interface {
	CreateParticipantReplacement(ctx context.Context, participantID, personID int64) error
}
