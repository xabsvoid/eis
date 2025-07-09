package service

import (
	"context"
	"fmt"
)

func (s Service) UpdateParticipantCheckInByParticipantID(ctx context.Context, participantID int64, exist bool) error {
	if !exist {
		err := s.repository.DeleteParticipantCheckInByParticipantID(ctx, participantID)
		if err != nil {
			return fmt.Errorf("repository.DeleteParticipantCheckInByParticipantID: %w", err)
		}

		return nil // check-in deleted successfully
	}

	err := s.repository.CreateParticipantCheckIn(ctx, participantID)
	if err != nil {
		return fmt.Errorf("repository.CreateParticipantCheckIn: %w", err)
	}

	return nil // check-in created successfully
}
