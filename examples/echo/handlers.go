package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var _ ServerInterface = (*Handlers)(nil)

type Handlers struct {
	Greeting string
}

func NewHandlers() *Handlers {
	return &Handlers{Greeting: "Hello"}
}

func (h *Handlers) GetGreeting(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Result{Result: h.Greeting})
}

func (h *Handlers) SetGreeting(ctx echo.Context, greeting string) error {
	h.Greeting = greeting

	return ctx.JSON(http.StatusOK, Result{Result: h.Greeting})
}

func (h *Handlers) GetHello(ctx echo.Context, name string) error {
	return ctx.JSON(http.StatusOK, Result{Result: h.Greeting + ", " + name + "!"})
}
