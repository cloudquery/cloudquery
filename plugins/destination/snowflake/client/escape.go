//go:build !windows

package client

// escapePath is not for non-Windows environments
func escapePath(p string) string {
	return p
}
