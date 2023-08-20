package main

// Code generated by gowrap. DO NOT EDIT.
// template: ../../sentry.gotmpl
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p github.com/adlandh/gowrap-templates/examples/echo-sentry -i AppInterface -t https://raw.githubusercontent.com/adlandh/gowrap-templates/main/sentry.gotmpl -o aplication_sentry_gen.go -l ""

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/getsentry/sentry-go"
)

// AppInterfaceWithSentry implements AppInterface interface instrumented with opentracing spans
type AppInterfaceWithSentry struct {
	AppInterface
	_instance string
}

// NewAppInterfaceWithSentry returns AppInterfaceWithSentry
func NewAppInterfaceWithSentry(base AppInterface, instance string) AppInterfaceWithSentry {
	d := AppInterfaceWithSentry{
		AppInterface: base,
		_instance:    instance,
	}

	return d
}

func (_d AppInterfaceWithSentry) _spanDecorator(span *sentry.Span, params, results map[string]interface{}) {
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

// GenFullGreeting implements AppInterface
func (_d AppInterfaceWithSentry) GenFullGreeting(ctx context.Context, name string) (greeting string, err error) {
	span := sentry.StartSpan(ctx, _d._instance+".AppInterface.GenFullGreeting", sentry.TransactionName("AppInterface.GenFullGreeting"))
	ctx = span.Context()

	defer func() {
		_d._spanDecorator(span, map[string]interface{}{
			"ctx":  ctx,
			"name": name}, map[string]interface{}{
			"greeting": greeting,
			"err":      err})
		if err != nil {
			span.SetTag("event", "error")
			span.SetTag("message", err.Error())
		}

		span.Finish()
	}()
	return _d.AppInterface.GenFullGreeting(ctx, name)
}

// GetGreeting implements AppInterface
func (_d AppInterfaceWithSentry) GetGreeting(ctx context.Context) (greeting string) {
	span := sentry.StartSpan(ctx, _d._instance+".AppInterface.GetGreeting", sentry.TransactionName("AppInterface.GetGreeting"))
	ctx = span.Context()

	defer func() {
		_d._spanDecorator(span, map[string]interface{}{
			"ctx": ctx}, map[string]interface{}{
			"greeting": greeting})
		span.Finish()
	}()
	return _d.AppInterface.GetGreeting(ctx)
}

// SetGreeting implements AppInterface
func (_d AppInterfaceWithSentry) SetGreeting(ctx context.Context, greeting string) (err error) {
	span := sentry.StartSpan(ctx, _d._instance+".AppInterface.SetGreeting", sentry.TransactionName("AppInterface.SetGreeting"))
	ctx = span.Context()

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
	return _d.AppInterface.SetGreeting(ctx, greeting)
}
