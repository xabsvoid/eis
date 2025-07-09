package inmem

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (m *InMem) GetParticipantByID(_ context.Context, id int64) (entity.Participant, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	participant, ok := m.participants[id]
	if !ok {
		return entity.Participant{}, ErrNotFound
	}

	return participant, nil
}
