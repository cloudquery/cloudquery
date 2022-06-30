//go:build darwin || linux

package cmd

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_setUlimit_when_low(t *testing.T) {
	// set to low values
	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &syscall.Rlimit{Cur: 10, Max: 20})
	require.NoError(t, err)

	err = setUlimit(ulimitUnix)
	require.NoError(t, err)
}
