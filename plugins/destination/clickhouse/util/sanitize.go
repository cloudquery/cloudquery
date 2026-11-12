package util

import (
	"strings"
)

func UnquoteID(sanitizedID string) string {
	if len(sanitizedID) >= 2 && sanitizedID[0] == '`' && sanitizedID[len(sanitizedID)-1] == '`' {
		return unescapeBacktickIdentifier(sanitizedID[1 : len(sanitizedID)-1])
	}
	return strings.Trim(sanitizedID, "`")
}

func unescapeBacktickIdentifier(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '`':
			if i+1 < len(s) && s[i+1] == '`' {
				b.WriteByte('`')
				i++
				continue
			}
			b.WriteByte('`')
		case '\\':
			if i+1 < len(s) && s[i+1] == '\\' {
				b.WriteByte('\\')
				i++
				continue
			}
			b.WriteByte('\\')
		default:
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

func SanitizeID(id string) string {
	id = strings.ReplaceAll(id, string([]byte{0}), "")
	id = strings.ReplaceAll(id, `\`, `\\`)
	id = strings.ReplaceAll(id, "`", "``")
	return "`" + id + "`"
}

func Sanitized(elems ...string) []string {
	result := make([]string, len(elems))
	for i, column := range elems {
		result[i] = SanitizeID(column)
	}
	return result
}
