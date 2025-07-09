package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handlers) GetEventsEventIdParticipants(ctx echo.Context, eventID int64) error { //nolint:revive,stylecheck
	participants, persons, err := h.service.GetParticipantsByEventID(ctx.Request().Context(), eventID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, convertError(err))
	}

	respParticipants := make([]Participant, 0, len(participants))
	for _, participant := range participants {
		respParticipants = append(respParticipants, convertParticipant(participant, persons))
	}

	respStatistic := calculateStatistic(participants)

	respBody := Participants{
		Participants: &respParticipants,
		Statistics:   &respStatistic,
	}

	return ctx.JSON(http.StatusOK, respBody)
}
