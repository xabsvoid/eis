package repository

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

type GetParticipantsByEventID interface {
	GetParticipantsByEventID(ctx context.Context, eventID int64) ([]entity.Participant, error)
}
