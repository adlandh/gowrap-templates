package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var _ ServerInterface = (*Handlers)(nil)

type Handlers struct {
	app AppInterface
}

func NewHandlers(app AppInterface) *Handlers {
	return &Handlers{app: app}
}

func (h *Handlers) GetGreeting(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Result{Result: h.app.GetGreeting(ctx.Request().Context())})
}

func (h *Handlers) SetGreeting(ctx echo.Context, greeting string) error {
	err := h.app.SetGreeting(ctx.Request().Context(), greeting)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, Result{Result: h.app.GetGreeting(ctx.Request().Context())})
}

func (h *Handlers) SayHello(ctx echo.Context, name string) error {
	greeting, err := h.app.GenFullGreeting(ctx.Request().Context(), name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, Result{Result: greeting})
}
