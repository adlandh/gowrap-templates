package sentry

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/getsentry/sentry-go"
	"github.com/stretchr/testify/require"
)

func TestDecorateTag(t *testing.T) {
	span := sentry.StartSpan(context.Background(), "test")

	t.Run("test span decorator with nil span", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("expected nil, got %v", r)
			}
		}()

		SpanDecorator(nil, nil, nil)
		require.Equal(t, span.Status, sentry.SpanStatusUndefined)
	})

	t.Run("test decorate tag with nil", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("expected nil, got %v", r)
			}
		}()

		decorateTag(span, "test", "test", nil)
		SpanDecorator(span, nil, nil)
		require.Equal(t, span.Status, sentry.SpanStatusOK)
	})

	t.Run("test decorate tag with nil response", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("expected nil, got %v", r)
			}
		}()

		var v *http.Response

		decorateTag(span, "test", "test", v)
		SpanDecorator(span, nil, nil)
		require.Equal(t, span.Status, sentry.SpanStatusOK)
	})

	t.Run("test decorate tag with nil request", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("expected nil, got %v", r)
			}
		}()

		var v *http.Request

		decorateTag(span, "test", "test", v)
		SpanDecorator(span, nil, nil)
		require.Equal(t, span.Status, sentry.SpanStatusOK)
	})

	t.Run("test decorate tag with nil error", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("expected nil, got %v", r)
			}
		}()

		var v error

		decorateTag(span, "test", "test", v)
		SpanDecorator(span, nil, nil)
		require.Equal(t, span.Status, sentry.SpanStatusOK)
	})

	t.Run("test decorate tag with error", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("expected nil, got %v", r)
			}
		}()

		v := errors.New("test")

		SpanDecorator(span, nil, map[string]interface{}{
			"error": v,
		})

		require.Equal(t, span.Status, sentry.SpanStatusInternalError)
	})
}
