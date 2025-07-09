package service

import (
	"context"
	"fmt"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (s Service) UpdateParticipantReplacementByParticipantID(ctx context.Context, participantID int64, exist bool, person entity.Person) error {
	if !exist {
		err := s.repository.DeleteParticipantReplacementByParticipantID(ctx, participantID)
		if err != nil {
			return fmt.Errorf("repository.DeleteParticipantReplacementByParticipantID: %w", err)
		}

		return nil // replacement deleted successfully
	}

	var err error
	person, err = s.repository.CreatePerson(ctx, person)
	if err != nil {
		return fmt.Errorf("repository.CreatePerson: %w", err)
	}

	err = s.repository.CreateParticipantReplacement(ctx, participantID, person.ID)
	if err != nil {
		return fmt.Errorf("repository.CreateParticipantReplacement: %w", err)
	}

	return nil // replacement created successfully
}
