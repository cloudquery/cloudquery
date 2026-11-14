package client

import "strings"

const (
	backslash               = `\`
	unicodeEscapedBackslash = backslash + "u005C"
)

var identifierReplacer = strings.NewReplacer(
	"\x00", "",
	backslash, unicodeEscapedBackslash,
	"`", "``",
)

func quoteIdentifier(name string) string {
	return "`" + identifierReplacer.Replace(name) + "`"
}
