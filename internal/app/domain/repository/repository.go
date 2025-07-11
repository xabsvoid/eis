package repository

type Repository interface {
	CreateEvent
	CreateParticipant
	CreateParticipantCheckIn
	CreateParticipantConfirmation
	CreateParticipantReplacement
	CreatePerson
	DeleteEventByID
	DeleteParticipantByID
	DeleteParticipantCheckInByParticipantID
	DeleteParticipantConfirmationByParticipantID
	DeleteParticipantReplacementByParticipantID
	GetEventByID
	GetEvents
	GetParticipantByID
	GetParticipantsByEventID
	GetPersonsByIDs
}
