package cmd

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/cloudquery/cloudquery/cli/v6/internal/enum"
	"github.com/stretchr/testify/require"
)

func TestInitLogging_OverwriteTruncatesExistingFile(t *testing.T) {
	dir := t.TempDir()
	logFileName := filepath.Join(dir, "cloudquery.log")
	require.NoError(t, os.WriteFile(logFileName, []byte("old content\n"), 0644))

	logLevel := enum.NewEnum([]string{"trace", "debug", "info", "warn", "error"}, "info")
	logFormat := enum.NewEnum([]string{"text", "json"}, "json")

	f, err := initLogging(false, logLevel, logFormat, false, logFileName, true)
	require.NoError(t, err)
	t.Cleanup(func() {
		if f != nil {
			f.Close()
		}
	})

	got, err := os.ReadFile(logFileName)
	require.NoError(t, err)
	require.NotContains(t, string(got), "old content")
}

func TestInitLogging_AppendPreservesExistingFile(t *testing.T) {
	dir := t.TempDir()
	logFileName := filepath.Join(dir, "cloudquery.log")
	require.NoError(t, os.WriteFile(logFileName, []byte("old content\n"), 0644))

	logLevel := enum.NewEnum([]string{"trace", "debug", "info", "warn", "error"}, "info")
	logFormat := enum.NewEnum([]string{"text", "json"}, "json")

	f, err := initLogging(false, logLevel, logFormat, false, logFileName, false)
	require.NoError(t, err)
	t.Cleanup(func() {
		if f != nil {
			f.Close()
		}
	})

	got, err := os.ReadFile(logFileName)
	require.NoError(t, err)
	require.Contains(t, string(got), "old content")
}

func TestInitLogging_CreatesOwnerOnlyFile(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("file permission bits are not supported on Windows")
	}
	logFileName := filepath.Join(t.TempDir(), "cloudquery.log")

	logLevel := enum.NewEnum([]string{"trace", "debug", "info", "warn", "error"}, "info")
	logFormat := enum.NewEnum([]string{"text", "json"}, "json")

	f, err := initLogging(false, logLevel, logFormat, false, logFileName, false)
	require.NoError(t, err)
	t.Cleanup(func() {
		if f != nil {
			f.Close()
		}
	})

	info, err := os.Stat(logFileName)
	require.NoError(t, err)
	require.Equal(t, os.FileMode(0600), info.Mode().Perm())
}
