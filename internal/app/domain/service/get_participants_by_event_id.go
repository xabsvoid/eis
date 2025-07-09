package service

import (
	"context"
	"fmt"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (s Service) GetParticipantsByEventID(ctx context.Context, eventID int64) ([]entity.Participant, []entity.Person, error) {
	participants, err := s.repository.GetParticipantsByEventID(ctx, eventID)
	if err != nil {
		return nil, nil, fmt.Errorf("repository.GetParticipantsByEventID: %w", err)
	}

	if len(participants) == 0 {
		return nil, nil, nil
	}

	totaPersonIDs := 0
	for _, participant := range participants {
		totaPersonIDs++
		if participant.Replacement.Exist {
			totaPersonIDs++
		}
	}

	personIDs := make([]int64, 0, totaPersonIDs)
	for _, participant := range participants {
		personIDs = append(personIDs, participant.PersonID)
		if participant.Replacement.Exist {
			personIDs = append(personIDs, participant.Replacement.PersonID)
		}
	}

	var persons []entity.Person
	persons, err = s.repository.GetPersonsByIDs(ctx, personIDs)
	if err != nil {
		return nil, nil, fmt.Errorf("repository.GetPersonsByIDs: %w", err)
	}

	return participants, persons, nil
}
