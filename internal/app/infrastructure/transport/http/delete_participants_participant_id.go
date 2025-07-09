package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handlers) DeleteParticipantsParticipantId(ctx echo.Context, participantID int64) error { //nolint:revive,stylecheck
	err := h.service.DeleteParticipantByID(ctx.Request().Context(), participantID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, convertError(err))
	}

	return ctx.NoContent(http.StatusOK)
}
