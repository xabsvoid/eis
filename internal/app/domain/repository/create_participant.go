package repository

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

type CreateParticipant interface {
	CreateParticipant(ctx context.Context, participant entity.Participant) (entity.Participant, error)
}
