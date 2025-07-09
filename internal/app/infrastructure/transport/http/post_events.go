package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
)

func (h Handlers) PostEvents(ctx echo.Context) error {
	var reqBody PostEventsJSONRequestBody
	err := ctx.Bind(&reqBody)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, convertError(err))
	}

	event := entity.NewEvent(
		convert(reqBody.Title),
		convert(reqBody.Date),
	)

	event, err = h.service.CreateEvent(ctx.Request().Context(), event)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, convertError(err))
	}

	respBody := convertEvent(event)

	return ctx.JSON(http.StatusCreated, respBody)
}
