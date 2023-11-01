//go:build windows

package client

import (
	"strings"
)

// escapePath properly escapes the `\` character in Window's file paths.
func escapePath(p string) string {
	return strings.ReplaceAll(p, "\\", "\\\\")
}
