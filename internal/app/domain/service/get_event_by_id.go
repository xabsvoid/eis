package service

import (
	"context"
	"fmt"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (s Service) GetEventByID(ctx context.Context, id int64) (entity.Event, error) {
	event, err := s.repository.GetEventByID(ctx, id)
	if err != nil {
		return entity.Event{}, fmt.Errorf("repository.GetEventByID: %w", err)
	}

	return event, nil
}
