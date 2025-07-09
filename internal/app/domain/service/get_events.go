package service

import (
	"context"
	"fmt"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (s Service) GetEvents(ctx context.Context) ([]entity.Event, error) {
	events, err := s.repository.GetEvents(ctx)
	if err != nil {
		return nil, fmt.Errorf("repository.GetEvents: %w", err)
	}

	return events, nil
}
