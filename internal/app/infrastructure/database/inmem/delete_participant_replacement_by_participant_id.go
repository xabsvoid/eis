package inmem

import (
	"context"
)

func (m *InMem) DeleteParticipantReplacementByParticipantID(_ context.Context, participantID int64) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	participant, ok := m.participants[participantID]
	if !ok {
		return ErrNotFound
	}
	participant.Replacement.Exist = false
	participant.Replacement.PersonID = 0
	participant.Metadata = refresh(participant.Metadata)
	m.participants[participantID] = participant
	return nil
}
