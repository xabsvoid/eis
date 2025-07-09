package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handlers) GetEventsEventId(ctx echo.Context, eventID int64) error { //nolint:revive,stylecheck
	event, err := h.service.GetEventByID(ctx.Request().Context(), eventID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, convertError(err))
	}

	respBody := convertEvent(event)

	return ctx.JSON(http.StatusOK, respBody)
}
