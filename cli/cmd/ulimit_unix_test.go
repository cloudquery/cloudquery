//go:build darwin || linux

package cmd

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This test call 'setUlimit', and expects the 'softUlimit' to increase up to 'hardUlimit'.
// Note that because ulimit is a process-wide variable, this test has some risk of causing flakiness in
// unrelated tests..
func Test_setUlimit_when_low(t *testing.T) {
	// We choose relatively "high" numbers because ulimit is a process-wide resource, and we don't
	// want to choose values so small that they make other tests fail.
	const smallSoftLimit = 500
	const smallHardLimit = 512

	// set to low values
	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &syscall.Rlimit{Cur: smallSoftLimit, Max: smallHardLimit})
	require.NoError(t, err)

	err = setUlimit()
	require.NoError(t, err)

	var limits syscall.Rlimit
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limits)
	require.NoError(t, err)

	assert.Equal(t, limits, syscall.Rlimit{Cur: smallHardLimit, Max: smallHardLimit})
}
