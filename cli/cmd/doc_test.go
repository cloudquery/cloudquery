package cmd

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

var docFiles = []string{
	"cloudquery.md",
	"cloudquery_addon.md",
	"cloudquery_addon_download.md",
	"cloudquery_addon_publish.md",
	"cloudquery_init.md",
	"cloudquery_login.md",
	"cloudquery_logout.md",
	"cloudquery_sync.md",
	"cloudquery_migrate.md",
	"cloudquery_tables.md",
	"cloudquery_test-connection.md",
	"cloudquery_validate-config.md",
	"cloudquery_plugin.md",
	"cloudquery_plugin_install.md",
	"cloudquery_plugin_publish.md",
	"cloudquery_switch.md",
}

func TestDoc(t *testing.T) {
	tmpDir := t.TempDir()
	cmd := NewCmdRoot()
	cmd.SetArgs(append([]string{"doc", tmpDir}, testCommandArgs(t)...))

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
