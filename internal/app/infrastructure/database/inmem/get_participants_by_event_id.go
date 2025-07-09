package inmem

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (m *InMem) GetParticipantsByEventID(_ context.Context, eventID int64) ([]entity.Participant, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	totalParticipants := 0
	for _, participant := range m.participants {
		if participant.EventID != eventID {
			continue
		}
		totalParticipants++
	}

	participants := make([]entity.Participant, 0, totalParticipants)
	for _, participant := range m.participants {
		if participant.EventID != eventID {
			continue
		}
		participants = append(participants, participant)
	}

	return participants, nil
}
