package service

import (
	"context"
	"fmt"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (s Service) CreateEvent(ctx context.Context, event entity.Event) (entity.Event, error) {
	var err error
	event, err = s.repository.CreateEvent(ctx, event)
	if err != nil {
		return entity.Event{}, fmt.Errorf("repository.CreateEvent: %w", err)
	}

	return event, nil
}
