package main

// Code generated by gowrap. DO NOT EDIT.
// template: https://raw.githubusercontent.com/adlandh/gowrap-templates/main/echo-sentry.gotmpl
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p github.com/adlandh/gowrap-templates/examples/echo-sentry -i ServerInterface -t https://raw.githubusercontent.com/adlandh/gowrap-templates/main/echo-sentry.gotmpl -o openapi_sentry_gen.go -l ""

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// ServerInterfaceWithSentry implements ServerInterface interface instrumented with opentracing spans
type ServerInterfaceWithSentry struct {
	ServerInterface
	_instance      string
	_spanDecorator func(span *sentry.Span, params, results map[string]interface{})
}

// NewServerInterfaceWithSentry returns ServerInterfaceWithSentry
func NewServerInterfaceWithSentry(base ServerInterface, instance string, spanDecorator ...func(span *sentry.Span, params, results map[string]interface{})) ServerInterfaceWithSentry {
	d := ServerInterfaceWithSentry{
		ServerInterface: base,
		_instance:       instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	} else {
		d._spanDecorator = d._defaultSpanDecorator
	}

	return d
}

func (_d ServerInterfaceWithSentry) _defaultSpanDecorator(span *sentry.Span, params, results map[string]interface{}) {
	for p := range params {
		switch params[p].(type) {
		case context.Context:
		case *http.Request:
			span.SetTag("param."+p+".method", params[p].(*http.Request).Method)
			val, _ := json.Marshal(params[p].(*http.Request).Header)
			span.SetTag("param."+p+".headers", string(val))
		case *http.Response:
			val, _ := json.Marshal(params[p].(*http.Response).Header)
			span.SetTag("param."+p+".headers", string(val))
		case echo.Context:
		default:
			val, _ := json.Marshal(params[p])
			span.SetTag("param."+p, string(val))
		}
	}

	for p := range results {
		switch results[p].(type) {
		case context.Context:
		case *http.Response:
			val, _ := json.Marshal(results[p].(*http.Response).Header)
			span.SetTag("result."+p+".headers", string(val))
		default:
			val, _ := json.Marshal(results[p])
			span.SetTag("result."+p, string(val))
		}
	}
}

// GetGreeting implements ServerInterface
func (_d ServerInterfaceWithSentry) GetGreeting(ctx echo.Context) (err error) {
	request := ctx.Request()
	savedCtx := request.Context()
	span := sentry.StartSpan(savedCtx, _d._instance+".ServerInterface.GetGreeting", sentry.TransactionName("ServerInterface.GetGreeting"))
	ctxNew := span.Context()

	defer func() {
		_d._spanDecorator(span, map[string]interface{}{
			"ctx": ctx}, map[string]interface{}{
			"err": err})
		if err != nil {
			span.SetTag("event", "error")
			span.SetTag("message", err.Error())
		}

		span.Finish()
	}()
	ctx.SetRequest(request.WithContext(ctxNew))
	return _d.ServerInterface.GetGreeting(ctx)
}

// SayHello implements ServerInterface
func (_d ServerInterfaceWithSentry) SayHello(ctx echo.Context, name string) (err error) {
	request := ctx.Request()
	savedCtx := request.Context()
	span := sentry.StartSpan(savedCtx, _d._instance+".ServerInterface.SayHello", sentry.TransactionName("ServerInterface.SayHello"))
	ctxNew := span.Context()

	defer func() {
		_d._spanDecorator(span, map[string]interface{}{
			"ctx":  ctx,
			"name": name}, map[string]interface{}{
			"err": err})
		if err != nil {
			span.SetTag("event", "error")
			span.SetTag("message", err.Error())
		}

		span.Finish()
	}()
	ctx.SetRequest(request.WithContext(ctxNew))
	return _d.ServerInterface.SayHello(ctx, name)
}

// SetGreeting implements ServerInterface
func (_d ServerInterfaceWithSentry) SetGreeting(ctx echo.Context, greeting string) (err error) {
	request := ctx.Request()
	savedCtx := request.Context()
	span := sentry.StartSpan(savedCtx, _d._instance+".ServerInterface.SetGreeting", sentry.TransactionName("ServerInterface.SetGreeting"))
	ctxNew := span.Context()

	defer func() {
		_d._spanDecorator(span, map[string]interface{}{
			"ctx":      ctx,
			"greeting": greeting}, map[string]interface{}{
			"err": err})
		if err != nil {
			span.SetTag("event", "error")
			span.SetTag("message", err.Error())
		}

		span.Finish()
	}()
	ctx.SetRequest(request.WithContext(ctxNew))
	return _d.ServerInterface.SetGreeting(ctx, greeting)
}
