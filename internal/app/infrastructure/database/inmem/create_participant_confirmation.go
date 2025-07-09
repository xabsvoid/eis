package inmem

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/valueobject"
)

func (m *InMem) CreateParticipantConfirmation(_ context.Context, participantID int64) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	participant, ok := m.participants[participantID]
	if !ok {
		return ErrNotFound
	}
	participant.Confirmation = valueobject.ParticipantConfirmation{
		Exist: true,
	}
	participant.Metadata = refresh(participant.Metadata)
	m.participants[participantID] = participant
	return nil
}
