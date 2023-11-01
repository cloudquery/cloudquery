//go:build windows

package client

// escapePath properly escapes the `\` character in Window's file paths.
func escapePath(p string) string {
	return strings.ReplaceAll(p, "\\", "\\\\")
}
