package entity

import (
	"time"

	"github.com/xabsvoid/eis/internal/app/domain/model/valueobject"
)

type Event struct {
	ID    int64
	Title string
	Date  time.Time
	valueobject.Metadata
}

func NewEvent(title string, date time.Time) Event {
	return Event{
		ID:    0,
		Title: title,
		Date:  date,
		Metadata: valueobject.Metadata{
			UpdatedAt: time.Time{},
			CreatedAt: time.Time{},
		},
	}
}
