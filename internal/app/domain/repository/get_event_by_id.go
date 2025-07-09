package repository

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

type GetEventByID interface {
	GetEventByID(ctx context.Context, id int64) (entity.Event, error)
}
