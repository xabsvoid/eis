package inmem

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (m *InMem) GetEvents(_ context.Context) ([]entity.Event, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	events := make([]entity.Event, 0, len(m.events))
	for _, event := range m.events {
		events = append(events, event)
	}

	return events, nil
}
