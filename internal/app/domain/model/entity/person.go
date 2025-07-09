package entity

import (
	"time"

	"github.com/xabsvoid/eis/internal/app/domain/model/valueobject"
)

type Person struct {
	ID         int64
	FirstName  string
	LastName   string
	MiddleName string
	Phone      string
	valueobject.Metadata
}

func NewPerson(firstName, lastName, middleName, phone string) Person {
	return Person{
		ID:         0,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: middleName,
		Phone:      phone,
		Metadata: valueobject.Metadata{
			UpdatedAt: time.Time{},
			CreatedAt: time.Time{},
		},
	}
}
