package utils

import "strings"

func StringArraySearch(array []string, query string) int {
	for index, element := range array {
		if strings.EqualFold(element, query) {
			return index
		}
	}

	return -1
}
