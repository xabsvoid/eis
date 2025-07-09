package repository

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

type CreatePerson interface {
	CreatePerson(ctx context.Context, person entity.Person) (entity.Person, error)
}
