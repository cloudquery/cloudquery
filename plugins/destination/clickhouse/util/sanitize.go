package util

import (
	"strings"
)

func SanitizeID(id string) string {
	return "`" + strings.ReplaceAll(id, string([]byte{0}), ``) + "`"
}

func Sanitized(elems ...string) []string {
	result := make([]string, len(elems))
	for i, column := range elems {
		result[i] = SanitizeID(column)
	}
	return result
}
