package inmem

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (m *InMem) CreateEvent(_ context.Context, event entity.Event) (entity.Event, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	event.ID = m.getNextEventID()
	event.Metadata = refresh(event.Metadata)
	m.events[event.ID] = event
	return event, nil
}
