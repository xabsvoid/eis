package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (h Handlers) PostParticipantsParticipantIdReplacement(ctx echo.Context, participantID int64) error { //nolint:revive,stylecheck
	var reqBody PostParticipantsParticipantIdReplacementJSONRequestBody
	err := ctx.Bind(&reqBody)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, convertError(err))
	}

	exist := convert(reqBody.Exist)
	person := entity.NewPerson(
		convert(reqBody.FirstName),
		convert(reqBody.LastName),
		convert(reqBody.MiddleName),
		convert(reqBody.Phone),
	)

	err = h.service.UpdateParticipantReplacementByParticipantID(ctx.Request().Context(), participantID, exist, person)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, convertError(err))
	}

	return ctx.NoContent(http.StatusOK)
}
