package inmem

import "context"

func (m *InMem) DeleteParticipantByID(_ context.Context, id int64) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.participants, id)
	return nil
}
