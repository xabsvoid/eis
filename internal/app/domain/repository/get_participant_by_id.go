package repository

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

type GetParticipantByID interface {
	GetParticipantByID(ctx context.Context, id int64) (entity.Participant, error)
}
