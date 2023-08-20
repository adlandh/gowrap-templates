package main

import (
	"context"
	"fmt"
)

type AppInterface interface {
	SetGreeting(ctx context.Context, greeting string) (err error)
	GetGreeting(ctx context.Context) (greeting string)
	GenFullGreeting(ctx context.Context, name string) (greeting string, err error)
}

var _ AppInterface = (*App)(nil)

type App struct {
	greeting string
}

func (a *App) SetGreeting(_ context.Context, greeting string) error {
	if greeting == "" {
		return fmt.Errorf("empty greeting")
	}

	a.greeting = greeting

	return nil
}

func (a *App) GetGreeting(_ context.Context) (greeting string) {
	return a.greeting
}

func (a *App) GenFullGreeting(_ context.Context, name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("empty name")
	}

	return a.greeting + ", " + name + "!", nil
}
