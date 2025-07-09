package repository

import "context"

type DeleteEventByID interface {
	DeleteEventByID(ctx context.Context, id int64) error
}
