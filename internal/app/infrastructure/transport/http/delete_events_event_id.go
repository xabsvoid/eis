package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handlers) DeleteEventsEventId(ctx echo.Context, eventID int64) error { //nolint:revive,stylecheck
	err := h.service.DeleteEventByID(ctx.Request().Context(), eventID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, convertError(err))
	}

	return ctx.NoContent(http.StatusOK)
}
