package cmd

import (
	"os"
	"path"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

var docFiles = []string{
	"cloudquery.md",
	"cloudquery_sync.md",
	"cloudquery_migrate.md",
	"cloudquery_tables.md",
	"cloudquery_install.md",
	"cloudquery_plugin.md",
	"cloudquery_plugin_install.md",
}

func TestDoc(t *testing.T) {
	defer CloseLogFile()
	cmd := NewCmdRoot()
	tmpDir := t.TempDir()
	cqTmpDir := t.TempDir()
	logFileName := path.Join(cqTmpDir, "cloudquery.log")
	cmd.SetArgs([]string{"doc", tmpDir, "--cq-dir", cqTmpDir, "--log-file-name", logFileName})

	err := cmd.Execute()
	require.NoError(t, err)

	files, err := os.ReadDir(tmpDir)
	require.NoError(t, err)

	fnames := make([]string, len(files))
	for i := range files {
		fnames[i] = files[i].Name()
	}
	sort.Strings(fnames)
	sort.Strings(docFiles)
	require.Equal(t, docFiles, fnames)
}
