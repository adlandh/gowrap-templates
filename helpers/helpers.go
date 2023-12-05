// Package helpers provides helper functions
package helpers

import "strings"

const (
	ParamPrefix  = "param"
	ResultPrefix = "result"
)

func LimitString(str string, size int) string {
	result := []rune(str)
	if size <= 0 || len(result) <= size {
		return str
	}

	return string(result[:size])
}

func LimitStringWithDots(str string, size int) string {
	result := []rune(str)
	if size <= 0 || len(result) <= size {
		return str
	}

	return string(result[:size-3]) + "..."
}

func PrepareTagValue(str string) string {
	str = strings.ReplaceAll(str, "\n", " ")
	return LimitStringWithDots(str, 200)
}

func PrepareTagName(str string) string {
	return LimitString(str, 32)
}
