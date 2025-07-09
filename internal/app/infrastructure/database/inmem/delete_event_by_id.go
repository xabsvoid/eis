package inmem

import "context"

func (m *InMem) DeleteEventByID(_ context.Context, id int64) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.events, id)
	return nil
}
