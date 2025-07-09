package repository

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

type GetEvents interface {
	GetEvents(ctx context.Context) ([]entity.Event, error)
}
