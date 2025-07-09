package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handlers) GetEvents(ctx echo.Context) error {
	events, err := h.service.GetEvents(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, convertError(err))
	}

	respBody := make(Events, 0, len(events))
	for _, event := range events {
		respBody = append(respBody, convertEvent(event))
	}

	return ctx.JSON(http.StatusOK, respBody)
}
