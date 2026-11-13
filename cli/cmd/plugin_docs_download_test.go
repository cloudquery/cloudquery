package cmd

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadedDocPathStaysInsideDocsDir(t *testing.T) {
	dir := t.TempDir()

	got, err := downloadedDocPath(dir, "overview")
	require.NoError(t, err)
	require.Equal(t, filepath.Join(dir, "overview.md"), got)

	for _, name := range []string{"../x", `..\x`, "..", "a/../../x", `a\..\..\x`, "/tmp/x"} {
		t.Run(name, func(t *testing.T) {
			got, err := downloadedDocPath(dir, name)
			if err != nil {
				return
			}
			require.Equal(t, dir, filepath.Dir(got))
		})
	}
}
