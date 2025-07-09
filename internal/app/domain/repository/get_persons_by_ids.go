package repository

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

type GetPersonsByIDs interface {
	GetPersonsByIDs(ctx context.Context, ids []int64) ([]entity.Person, error)
}
