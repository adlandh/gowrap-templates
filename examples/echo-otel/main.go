package main

import (
	echo_otel_middleware "github.com/adlandh/echo-otel-middleware"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func main() {
	// Create otel exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		panic(err)
	}

	// Create otel resource
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("echo-example"),
		),
	)
	if err != nil {
		panic(err)
	}

	// Create trace provider
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exp),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	handlers := NewServerInterfaceWithTracing(NewHandlers(), "handlers")

	e := echo.New()
	e.Use(echo_otel_middleware.MiddlewareWithConfig(
		echo_otel_middleware.OtelConfig{
			AreHeadersDump: true,
			IsBodyDump:     true,
			LimitHTTPBody:  true,
			LimitSize:      60_000,
		}))

	RegisterHandlers(e, handlers)

	e.Logger.Fatal(e.Start(":1234"))
}
