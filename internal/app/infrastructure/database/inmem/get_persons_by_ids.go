package inmem

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (m *InMem) GetPersonsByIDs(_ context.Context, ids []int64) ([]entity.Person, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	mIDs := make(map[int64]struct{})
	for _, id := range ids {
		mIDs[id] = struct{}{}
	}

	persons := make([]entity.Person, 0, len(mIDs))
	for id := range mIDs {
		if person, ok := m.persons[id]; ok {
			persons = append(persons, person)
		}
	}

	return persons, nil
}
