package client

import "strings"

var identifierReplacer = strings.NewReplacer(
	"\x00", "",
	`\u0060`, "``",
	"`", "``",
)

func quoteIdentifier(name string) string {
	return "`" + identifierReplacer.Replace(name) + "`"
}
