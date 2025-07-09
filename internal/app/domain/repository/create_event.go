package repository

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

type CreateEvent interface {
	CreateEvent(ctx context.Context, event entity.Event) (entity.Event, error)
}
