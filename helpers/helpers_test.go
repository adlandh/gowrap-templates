package helpers

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func TestLimitString(t *testing.T) {
	originalStr := "Привет, мир!"
	maxBytes := 9

	result := LimitString(originalStr, maxBytes)

	require.GreaterOrEqual(t, maxBytes, len(result))
	require.True(t, utf8.ValidString(result))
}

func TestLimitStringWithDots(t *testing.T) {
	originalStr := "Привет, мир! Привет, мир!"
	maxBytes := 11

	result := LimitStringWithDots(originalStr, maxBytes)

	require.GreaterOrEqual(t, maxBytes, len(result))
	require.True(t, utf8.ValidString(result))

	maxBytes = 200

	result = LimitStringWithDots(originalStr, maxBytes)

	require.GreaterOrEqual(t, maxBytes, len(result))
	require.True(t, utf8.ValidString(result))
	require.Equal(t, originalStr, result)
}
