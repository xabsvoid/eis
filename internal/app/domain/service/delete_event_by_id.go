package service

import (
	"context"
	"fmt"
)

func (s Service) DeleteEventByID(ctx context.Context, id int64) error {
	err := s.repository.DeleteEventByID(ctx, id)
	if err != nil {
		return fmt.Errorf("repository.DeleteEventByID: %w", err)
	}

	return nil
}
