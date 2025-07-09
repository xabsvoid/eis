package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handlers) PostParticipantsParticipantIdCheckIn(ctx echo.Context, participantID int64) error { //nolint:revive,stylecheck
	var reqBody PostParticipantsParticipantIdCheckInJSONRequestBody
	err := ctx.Bind(&reqBody)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, convertError(err))
	}

	exist := convert(reqBody.Exist)

	err = h.service.UpdateParticipantCheckInByParticipantID(ctx.Request().Context(), participantID, exist)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, convertError(err))
	}

	return ctx.NoContent(http.StatusOK)
}
