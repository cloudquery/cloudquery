package util

import (
	"strings"
)

var (
	idEscaper   = strings.NewReplacer(string([]byte{0}), ``, `\`, `\\`, "`", "\\`")
	idUnescaper = strings.NewReplacer(`\\`, `\`, "\\`", "`")
)

func UnquoteID(sanitizedID string) string {
	if len(sanitizedID) < 2 || !strings.HasPrefix(sanitizedID, "`") || !strings.HasSuffix(sanitizedID, "`") {
		return strings.Trim(sanitizedID, "`")
	}
	return idUnescaper.Replace(sanitizedID[1 : len(sanitizedID)-1])
}

func SanitizeID(id string) string {
	return "`" + idEscaper.Replace(id) + "`"
}

func Sanitized(elems ...string) []string {
	result := make([]string, len(elems))
	for i, column := range elems {
		result[i] = SanitizeID(column)
	}
	return result
}
