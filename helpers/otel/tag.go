package sentry

import (
	"context"
	"io"
	"net/http"

	"github.com/adlandh/gowrap-templates/helpers"
	"github.com/goccy/go-json"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func SetTag(span trace.Span, tag, value string) {
	if tag == "" || value == "" {
		return
	}

	span.SetAttributes(attribute.String(helpers.PrepareTagName(tag), helpers.PrepareTagValue(value)))
}

func SetErrorTags(span trace.Span, err error) {
	if err == nil {
		return
	}
	SetTag(span, "event", "error")
	SetTag(span, "message", err.Error())
}

func SpanDecorator(span trace.Span, params, results map[string]interface{}) {
	for p, v := range params {
		decorateTag(span, "param", p, v)
	}

	for p, v := range results {
		decorateTag(span, "result", p, v)
	}
}

func decorateTag(span trace.Span, prefix string, p string, v any) {
	switch v := v.(type) {
	case context.Context:
	case io.Reader:
	case io.Writer:
	case echo.Context:
	case *http.Request:
		SetTag(span, prefix+"."+p+".method", v.Method)
		val, _ := json.Marshal(v.Header)
		SetTag(span, prefix+"."+p+".headers", string(val))
	case *http.Response:
		val, _ := json.Marshal(v.Header)
		SetTag(span, prefix+"."+p+".headers", string(val))
	case []byte:
		SetTag(span, prefix+"."+p, string(v))
	case error:
		if v != nil {
			span.RecordError(v)
			SetTag(span, prefix+"."+p, v.Error())
			SetErrorTags(span, v)
		}
	default:
		val, _ := json.Marshal(v)
		SetTag(span, "param."+p, string(val))
	}
}
