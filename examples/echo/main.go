package main

import (
	"context"
	"encoding/json"
	"net/http"

	echo_otel_middleware "github.com/adlandh/echo-otel-middleware"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

func SpanDecorator(span trace.Span, params, results map[string]interface{}) {
	for p := range params {
		switch params[p].(type) {
		case context.Context:
		case *http.Request:
			span.SetAttributes(attribute.String("param."+p+".method", params[p].(*http.Request).Method))
			val, _ := json.Marshal(params[p].(*http.Request).Header)
			span.SetAttributes(attribute.String("param."+p+".headers", string(val)))
		case *http.Response:
			val, _ := json.Marshal(params[p].(*http.Response).Header)
			span.SetAttributes(attribute.String("param."+p+".headers", string(val)))
		case echo.Context:
		default:
			val, _ := json.Marshal(params[p])
			span.SetAttributes(attribute.String("param."+p, string(val)))
		}
	}

	for p := range results {
		switch results[p].(type) {
		case context.Context:
		case *http.Response:
			val, _ := json.Marshal(results[p].(*http.Response).Header)
			span.SetAttributes(attribute.String("result."+p+".headers", string(val)))
		default:
			val, _ := json.Marshal(results[p])
			span.SetAttributes(attribute.String("result."+p, string(val)))
		}
	}
}

func main() {
	// Create otel exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://jaeger:14268/api/traces")))
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
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	handlers := NewServerInterfaceWithTracing(NewHandlers(), "handlers", SpanDecorator)

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
