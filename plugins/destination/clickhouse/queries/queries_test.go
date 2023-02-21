package queries

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func ensureContents(t *testing.T, data string, path string) {
	contents, err := os.ReadFile("testdata/" + path)
	require.NoError(t, err)
	require.Equal(t, string(contents), data)
}
