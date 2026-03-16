package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppendToFile_AppendsData(t *testing.T) {
	f := filepath.Join(t.TempDir(), "summary.json")
	require.NoError(t, os.WriteFile(f, []byte("line1\n"), 0644))

	err := appendToFile(f, []byte("line2\n"))
	require.NoError(t, err)

	got, err := os.ReadFile(f)
	require.NoError(t, err)
	require.Equal(t, "line1\nline2\n", string(got))
}

func TestAppendToFileFallback_ExistingFile(t *testing.T) {
	f := filepath.Join(t.TempDir(), "summary.json")
	require.NoError(t, os.WriteFile(f, []byte("line1\n"), 0644))

	err := appendToFileFallback(f, []byte("line2\n"))
	require.NoError(t, err)

	got, err := os.ReadFile(f)
	require.NoError(t, err)
	require.Equal(t, "line1\nline2\n", string(got))
}

func TestAppendToFileFallback_NewFile(t *testing.T) {
	f := filepath.Join(t.TempDir(), "summary.json")

	err := appendToFileFallback(f, []byte("line1\n"))
	require.NoError(t, err)

	got, err := os.ReadFile(f)
	require.NoError(t, err)
	require.Equal(t, "line1\n", string(got))
}
