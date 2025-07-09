package inmem

import (
	"context"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (m *InMem) CreatePerson(_ context.Context, person entity.Person) (entity.Person, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	person.ID = m.getNextPersonID()
	person.Metadata = refresh(person.Metadata)
	m.persons[person.ID] = person
	return person, nil
}
