// Package helpers provides helper functions
package helpers

import (
	"strings"
	"unicode/utf8"
)

const (
	ParamPrefix  = "param"
	ResultPrefix = "result"
)

func LimitString(str string, size int) string {
	if len(str) <= size {
		return str
	}

	bytes := []byte(str)

	if len(bytes) <= size {
		return str
	}

	validBytes := bytes[:size]
	for !utf8.Valid(validBytes) {
		validBytes = validBytes[:len(validBytes)-1]
	}

	return string(validBytes)
}

func LimitStringWithDots(str string, size int) string {
	if size <= 10 {
		return LimitString(str, size)
	}

	result := LimitString(str, size-3)

	return result + "..."
}

func PrepareTagValue(str string) string {
	str = strings.ReplaceAll(str, "\n", " ")
	return LimitStringWithDots(str, 200)
}

func PrepareTagName(str string) string {
	return LimitString(str, 32)
}
