package helpers

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
