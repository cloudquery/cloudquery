package client

import (
	"errors"
	"fmt"
	"strings"
)

const (
	backslash               = `\`
	unicodeEscapedBackslash = backslash + "u005C"
)

var errNULInIdentifier = errors.New("identifier contains a NUL byte")

var identifierReplacer = strings.NewReplacer(
	backslash, unicodeEscapedBackslash,
	"`", "``",
)

func quoteIdentifier(name string) (string, error) {
	if strings.IndexByte(name, 0) >= 0 {
		return "", fmt.Errorf("%w: %q", errNULInIdentifier, name)
	}
	return "`" + identifierReplacer.Replace(name) + "`", nil
}
