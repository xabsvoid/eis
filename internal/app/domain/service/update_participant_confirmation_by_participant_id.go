package service

import (
	"context"
	"fmt"
)

func (s Service) UpdateParticipantConfirmationByParticipantID(ctx context.Context, participantID int64, exist bool) error {
	if !exist {
		err := s.repository.DeleteParticipantConfirmationByParticipantID(ctx, participantID)
		if err != nil {
			return fmt.Errorf("repository.DeleteParticipantConfirmationByParticipantID: %w", err)
		}

		return nil // confirmation deleted successfully
	}

	err := s.repository.CreateParticipantConfirmation(ctx, participantID)
	if err != nil {
		return fmt.Errorf("repository.CreateParticipantConfirmation: %w", err)
	}

	return nil // confirmation created successfully
}
