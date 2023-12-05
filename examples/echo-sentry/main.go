package main

import (
	"log"

	echo_sentry_middleware "github.com/adlandh/echo-sentry-middleware"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

//go:generate oapi-codegen -old-config-style -generate types,server -o "openapi_gen.go" -package "main" "api.yaml"
//go:generate gowrap gen -i ServerInterface -t https://raw.githubusercontent.com/adlandh/gowrap-templates/main/echo-sentry.gotmpl -o openapi_sentry_gen.go -g
//go:generate gowrap gen -i AppInterface -t https://raw.githubusercontent.com/adlandh/gowrap-templates/main/sentry.gotmpl -o aplication_sentry_gen.go -g

func main() {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		EnableTracing:      true,
		TracesSampleRate:   1.0,
		ProfilesSampleRate: 1.0,
	}); err != nil {
		log.Fatalln("Sentry initialization failed:", err)
	}

	e := echo.New()
	e.Use(sentryecho.New(sentryecho.Options{}))
	e.Use(echo_sentry_middleware.MiddlewareWithConfig(
		echo_sentry_middleware.SentryConfig{
			// if you would like to save your request or response headers as tags, set AreHeadersDump to true
			AreHeadersDump: true,
			// if you would like to save your request or response body as tags, set IsBodyDump to true
			IsBodyDump: true,
		}))

	app := NewAppInterfaceWithSentry[string](&App{"Hello"}, "app")

	handlers := NewServerInterfaceWithSentry(NewHandlers(app), "handlers")

	RegisterHandlers(e, handlers)

	e.Logger.Fatal(e.Start(":1234"))
}
