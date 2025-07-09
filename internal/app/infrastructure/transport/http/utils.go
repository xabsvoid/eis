package http

import "github.com/xabsvoid/eis/internal/app/domain/model/entity"

func convert[T any](in *T) T {
	var out T
	if in != nil {
		out = *in
	}
	return out
}

func convertError(err error) map[string]string {
	return map[string]string{"message": err.Error()}
}

func convertEvent(event entity.Event) Event {
	return Event{
		Date:  &event.Date,
		Id:    &event.ID,
		Title: &event.Title,
	}
}

func convertParticipant(participant entity.Participant, persons []entity.Person) Participant {
	person := findPerson(persons, participant.PersonID)
	replacementPerson := findPerson(persons, participant.Replacement.PersonID)

	return Participant{
		CheckIn: &Nullable{
			Exist: &participant.CheckIn.Exist,
		},
		Confirmation: &Nullable{
			Exist: &participant.Confirmation.Exist,
		},
		EventId:    &participant.EventID,
		FirstName:  &person.FirstName,
		Id:         &participant.ID,
		LastName:   &person.LastName,
		MiddleName: &person.MiddleName,
		Phone:      &person.Phone,
		Replacement: &Replacement{
			Exist:      &participant.Replacement.Exist,
			FirstName:  &replacementPerson.FirstName,
			LastName:   &replacementPerson.LastName,
			MiddleName: &replacementPerson.MiddleName,
			Phone:      &replacementPerson.Phone,
		},
	}
}

func findPerson(persons []entity.Person, personID int64) entity.Person {
	for _, person := range persons {
		if person.ID != personID {
			continue
		}
		return person
	}
	return entity.Person{} //nolint:exhaustruct
}

func calculateStatistic(participants []entity.Participant) Statistics {
	checkedIn := int64(0)
	confirmed := int64(0)
	replaced := int64(0)

	for _, participant := range participants {
		if participant.CheckIn.Exist {
			checkedIn++
		}
		if participant.Confirmation.Exist {
			confirmed++
		}
		if participant.Replacement.Exist {
			replaced++
		}
	}

	invited := int64(len(participants))
	notCheckedIn := invited - checkedIn
	notConfirmed := invited - confirmed

	return Statistics{
		CheckedIn:    &checkedIn,
		Confirmed:    &confirmed,
		Invited:      &invited,
		NotCheckedIn: &notCheckedIn,
		NotConfirmed: &notConfirmed,
		Replaced:     &replaced,
	}
}
