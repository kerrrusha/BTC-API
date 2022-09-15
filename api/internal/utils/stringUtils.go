package utils

import "strings"

func RemoveRedundantGaps(str string) string {
	return strings.ReplaceAll(str, "\"", "")
}

func SubstringBetween(main string, start string, end string) string {
	startIndex := strings.Index(main, start) + len(start)
	endIndex := startIndex + strings.Index(main[startIndex:], end)
	return main[startIndex:endIndex]
}
