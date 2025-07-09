package service

import (
	"context"
	"fmt"
)

func (s Service) DeleteParticipantByID(ctx context.Context, id int64) error {
	err := s.repository.DeleteParticipantByID(ctx, id)
	if err != nil {
		return fmt.Errorf("repository.DeleteParticipantByID: %w", err)
	}

	return nil
}
