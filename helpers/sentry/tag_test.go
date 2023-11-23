package sentry

import (
	"context"
	"net/http"
	"testing"

	"github.com/getsentry/sentry-go"
)

func TestDecorateTag(t *testing.T) {
	span := sentry.StartSpan(context.Background(), "test")

	t.Run("test decorate tag with nil", func(t *testing.T) {
		if r := recover(); r != nil {
			t.Errorf("expected nil, got %v", r)
		}
		decorateTag(span, "test", "test", nil)
		SpanDecorator(span, nil, nil)
	})

	t.Run("test decorate tag with nil response", func(t *testing.T) {
		if r := recover(); r != nil {
			t.Errorf("expected nil, got %v", r)
		}

		var v *http.Response

		decorateTag(span, "test", "test", v)
		SpanDecorator(span, nil, nil)
	})

	t.Run("test decorate tag with nil request", func(t *testing.T) {
		if r := recover(); r != nil {
			t.Errorf("expected nil, got %v", r)
		}

		var v *http.Request

		decorateTag(span, "test", "test", v)
		SpanDecorator(span, nil, nil)
	})

	t.Run("test decorate tag with nil error", func(t *testing.T) {
		if r := recover(); r != nil {
			t.Errorf("expected nil, got %v", r)
		}

		var v error

		decorateTag(span, "test", "test", v)
		SpanDecorator(span, nil, nil)
	})
}
