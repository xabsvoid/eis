package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handlers) GetParticipantsParticipantId(ctx echo.Context, participantID int64) error { //nolint:revive,stylecheck
	participant, persons, err := h.service.GetParticipantByID(ctx.Request().Context(), participantID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, convertError(err))
	}

	respBody := convertParticipant(participant, persons)

	return ctx.JSON(http.StatusOK, respBody)
}
