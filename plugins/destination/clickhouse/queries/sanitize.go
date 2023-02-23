package queries

import (
	"strings"
)

func sanitizeID(id string) string {
	return "`" + strings.ReplaceAll(id, string([]byte{0}), ``) + "`"
}

func sanitized(elems ...string) []string {
	result := make([]string, len(elems))
	for i, column := range elems {
		result[i] = sanitizeID(column)
	}
	return result
}
