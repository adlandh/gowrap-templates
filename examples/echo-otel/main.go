package main

import (
	echo_otel_middleware "github.com/adlandh/echo-otel-middleware"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

//go:generate oapi-codegen -old-config-style -generate types,server -o "openapi_gen.go" -package "main" "api.yaml"
//go:generate gowrap gen -i ServerInterface -t https://raw.githubusercontent.com/adlandh/gowrap-templates/main/echo-otel.gotmpl -o openapi_otel_gen.go -g
//go:generate gowrap gen -i AppInterface -t https://raw.githubusercontent.com/adlandh/gowrap-templates/main/otel.gotmpl -o aplication_otel_gen.go -g

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

	app := NewAppInterfaceWithTracing[string](&App{"Hello"}, "app")

	handlers := NewServerInterfaceWithTracing(NewHandlers(app), "handlers")

	e := echo.New()
	e.Use(echo_otel_middleware.MiddlewareWithConfig(
		echo_otel_middleware.OtelConfig{
			AreHeadersDump: true,
			IsBodyDump:     true,
			LimitNameSize:  32,
		}))

	RegisterHandlers(e, handlers)

	e.Logger.Fatal(e.Start(":1234"))
}
