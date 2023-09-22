package helpers

import "strings"

func LimitString(str string, size int) string {
	if size <= 0 || len(str) <= size {
		return str
	}

	return str[:size]
}

func LimitStringWithDots(str string, size int) string {
	if size <= 0 || len(str) <= size {
		return str
	}

	return str[:size-3] + "..."
}

func PrepareTagValue(str string) string {
	str = strings.Replace(str, "\n", " ", -1)
	return LimitStringWithDots(str, 200)
}

func PrepareTagName(str string) string {
	return LimitString(str, 32)
}
