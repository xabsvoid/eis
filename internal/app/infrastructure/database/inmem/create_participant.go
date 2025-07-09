package inmem

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (m *InMem) CreateParticipant(_ context.Context, participant entity.Participant) (entity.Participant, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	participant.ID = m.getNextParticipantID()
	participant.Metadata = refresh(participant.Metadata)
	m.participants[participant.ID] = participant
	return participant, nil
}
