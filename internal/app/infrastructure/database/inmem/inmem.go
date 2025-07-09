package inmem

import (
	"errors"
	"sync"

	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

var ErrNotFound = errors.New("not found")

type InMem struct {
	mutex         sync.RWMutex
	eventID       int64
	participantID int64
	personID      int64
	events        map[int64]entity.Event
	participants  map[int64]entity.Participant
	persons       map[int64]entity.Person
}

func NewInMem() *InMem {
	return &InMem{
		mutex:         sync.RWMutex{},
		eventID:       0,
		participantID: 0,
		personID:      0,
		events:        make(map[int64]entity.Event),
		participants:  make(map[int64]entity.Participant),
		persons:       make(map[int64]entity.Person),
	}
}

func (m *InMem) getNextEventID() int64 {
	m.eventID++
	return m.eventID
}

func (m *InMem) getNextParticipantID() int64 {
	m.participantID++
	return m.participantID
}

func (m *InMem) getNextPersonID() int64 {
	m.personID++
	return m.personID
}
