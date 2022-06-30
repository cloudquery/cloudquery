//go:build darwin || linux

package cmd

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_setUlimit(t *testing.T) {
	var initialLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &initialLimit)
	require.NoError(t, err)

	err = setUlimit(ulimitUnix)
	require.NoError(t, err)

	var actualLimit syscall.Rlimit
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &actualLimit)
	require.NoError(t, err)

	if initialLimit.Max < ulimitUnix {
		require.Equal(t, ulimitUnix, actualLimit.Max)
	} else {
		require.Equal(t, initialLimit.Max, actualLimit.Max)
	}

	if initialLimit.Cur < ulimitUnix {
		require.Equal(t, ulimitUnix, actualLimit.Cur)
	} else {
		require.Equal(t, initialLimit.Cur, actualLimit.Cur)
	}
}
