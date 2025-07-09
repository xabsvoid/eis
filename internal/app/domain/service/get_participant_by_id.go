package service

import (
	"context"
	"fmt"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (s Service) GetParticipantByID(ctx context.Context, id int64) (entity.Participant, []entity.Person, error) {
	participant, err := s.repository.GetParticipantByID(ctx, id)
	if err != nil {
		return entity.Participant{}, nil, fmt.Errorf("repository.GetParticipantByID: %w", err)
	}

	const maxPersonIDsLen = 2
	personIDs := make([]int64, 0, maxPersonIDsLen)
	personIDs = append(personIDs, participant.PersonID)
	if participant.Replacement.Exist {
		personIDs = append(personIDs, participant.Replacement.PersonID)
	}

	var persons []entity.Person
	persons, err = s.repository.GetPersonsByIDs(ctx, personIDs)
	if err != nil {
		return entity.Participant{}, nil, fmt.Errorf("repository.GetPersonsByIDs: %w", err)
	}

	return participant, persons, nil
}
