// Package sentry provides helpers for sentry wrappers
package sentry

import (
	"context"
	"io"
	"net/http"

	"github.com/adlandh/gowrap-templates/helpers"
	"github.com/getsentry/sentry-go"
	"github.com/goccy/go-json"
	"github.com/labstack/echo/v4"
)

func SetTag(span *sentry.Span, tag, value string) {
	if tag == "" || value == "" {
		return
	}

	span.SetTag(helpers.PrepareTagName(tag), helpers.PrepareTagValue(value))
}

func SetErrorTags(span *sentry.Span, err error) {
	if err == nil {
		return
	}

	SetTag(span, "event", "error")
	SetTag(span, "message", err.Error())
}

func SpanDecorator(span *sentry.Span, params, results map[string]interface{}) {
	if span == nil {
		return
	}

	span.Status = sentry.SpanStatusOK

	for p, v := range params {
		decorateTag(span, helpers.ParamPrefix, p, v)
	}

	for p, v := range results {
		decorateTag(span, helpers.ResultPrefix, p, v)
	}
}

//nolint:cyclop
func decorateTag(span *sentry.Span, prefix string, p string, v any) {
	switch v := v.(type) {
	case context.Context:
	case io.Reader:
	case io.Writer:
	case echo.Context:
	case *http.Request:
		if v == nil {
			return
		}

		SetTag(span, prefix+"."+p+".method", v.Method)
		val, _ := json.Marshal(v.Header)
		SetTag(span, prefix+"."+p+".headers", string(val))
	case *http.Response:
		if v == nil {
			return
		}

		val, _ := json.Marshal(v.Header)
		SetTag(span, prefix+"."+p+".headers", string(val))
	case []byte:
		if v == nil {
			return
		}

		SetTag(span, prefix+"."+p, string(v))
	case error:
		if v == nil {
			return
		}

		SetTag(span, prefix+"."+p, v.Error())

		if prefix == helpers.ResultPrefix {
			span.Status = sentry.SpanStatusInternalError
			SetErrorTags(span, v)
		}
	default:
		if v == nil {
			return
		}

		val, _ := json.Marshal(v)
		SetTag(span, prefix+"."+p, string(val))
	}
}
