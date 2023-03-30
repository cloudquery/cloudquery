package testing_cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToBeTheOnlyRun(t *testing.T) {
	require.Equal(t, 100, 1000/10)
}
