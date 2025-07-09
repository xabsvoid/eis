package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (h Handlers) PostEventsEventIdParticipants(ctx echo.Context, eventId int64) error { //nolint:revive,stylecheck
	var reqBody PostEventsEventIdParticipantsJSONRequestBody
	err := ctx.Bind(&reqBody)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, convertError(err))
	}

	person := entity.NewPerson(
		convert(reqBody.FirstName),
		convert(reqBody.LastName),
		convert(reqBody.MiddleName),
		convert(reqBody.Phone),
	)

	var participant entity.Participant
	participant, person, err = h.service.CreateParticipant(ctx.Request().Context(), eventId, person)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, convertError(err))
	}

	respBody := convertParticipant(participant, []entity.Person{person})

	return ctx.JSON(http.StatusCreated, respBody)
}
