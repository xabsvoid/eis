package inmem

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/valueobject"
)

func (m *InMem) CreateParticipantReplacement(_ context.Context, participantID, personID int64) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	participant, ok := m.participants[participantID]
	if !ok {
		return ErrNotFound
	}
	participant.Replacement = valueobject.ParticipantReplacement{
		Exist:    true,
		PersonID: personID,
	}
	participant.Metadata = refresh(participant.Metadata)
	m.participants[participantID] = participant
	return nil
}
