package entity

import (
	"time"

	"github.com/xabsvoid/eis/internal/app/domain/model/valueobject"
)

type Participant struct {
	ID           int64
	EventID      int64
	PersonID     int64
	CheckIn      valueobject.ParticipantCheckIn
	Confirmation valueobject.ParticipantConfirmation
	Replacement  valueobject.ParticipantReplacement
	valueobject.Metadata
}

func NewParticipant(eventID, personID int64) Participant {
	return Participant{
		ID:       0,
		EventID:  eventID,
		PersonID: personID,
		CheckIn: valueobject.ParticipantCheckIn{
			Exist: false,
		},
		Confirmation: valueobject.ParticipantConfirmation{
			Exist: false,
		},
		Replacement: valueobject.ParticipantReplacement{
			Exist:    false,
			PersonID: 0,
		},
		Metadata: valueobject.Metadata{
			UpdatedAt: time.Time{},
			CreatedAt: time.Time{},
		},
	}
}
