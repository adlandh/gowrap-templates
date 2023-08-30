package main

import (
	"context"
	"fmt"
)

type AppInterface[T any] interface {
	SetGreeting(ctx context.Context, greeting T) (err error)
	GetGreeting(ctx context.Context) (greeting T)
	GenFullGreeting(ctx context.Context, name T) (greeting string, err error)
}

var _ AppInterface[string] = (*App)(nil)

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
