package inmem

import (
	"context"
)

func (m *InMem) DeleteParticipantCheckInByParticipantID(_ context.Context, participantID int64) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	participant, ok := m.participants[participantID]
	if !ok {
		return ErrNotFound
	}
	participant.CheckIn.Exist = false
	participant.Metadata = refresh(participant.Metadata)
	m.participants[participantID] = participant
	return nil
}
