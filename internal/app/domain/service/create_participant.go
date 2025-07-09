package service

import (
	"context"
	"fmt"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (s Service) CreateParticipant(ctx context.Context, eventID int64, person entity.Person) (entity.Participant, entity.Person, error) {
	var err error
	person, err = s.repository.CreatePerson(ctx, person)
	if err != nil {
		return entity.Participant{}, entity.Person{}, fmt.Errorf("repository.CreatePerson: %w", err)
	}

	participant := entity.NewParticipant(eventID, person.ID)
	participant, err = s.repository.CreateParticipant(ctx, participant)
	if err != nil {
		return entity.Participant{}, entity.Person{}, fmt.Errorf("repository.CreateParticipant: %w", err)
	}

	return participant, person, nil
}
