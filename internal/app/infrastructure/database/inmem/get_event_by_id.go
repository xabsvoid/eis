package inmem

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (m *InMem) GetEventByID(_ context.Context, id int64) (entity.Event, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	event, ok := m.events[id]
	if !ok {
		return entity.Event{}, ErrNotFound
	}
	return event, nil
}
