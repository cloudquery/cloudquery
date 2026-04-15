package cmd

import (
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// fencedCodeBlock matches fenced code blocks (``` ... ```) so we can strip
// them before inspecting markdown heading structure.
var fencedCodeBlock = regexp.MustCompile("(?ms)^```[^\n]*\n.*?^```")

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
	slices.Sort(fnames)
	slices.Sort(docFiles)
	require.Equal(t, docFiles, fnames)
}

// TestDocHeadingStructure verifies that every generated markdown page has
// exactly one H1 heading (the command name), no all-caps "SEE ALSO" headings,
// and at most one "## See Also" section.
func TestDocHeadingStructure(t *testing.T) {
	tmpDir := t.TempDir()
	cmd := NewCmdRoot()
	cmd.SetArgs(append([]string{"doc", tmpDir}, testCommandArgs(t)...))
	require.NoError(t, cmd.Execute())

	entries, err := os.ReadDir(tmpDir)
	require.NoError(t, err)

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		data, err := os.ReadFile(tmpDir + "/" + e.Name())
		require.NoError(t, err)
		content := string(data)

		t.Run(e.Name(), func(t *testing.T) {
			// Strip fenced code blocks before inspecting heading structure so
			// that bash comments (# ...) inside examples are not counted.
			stripped := fencedCodeBlock.ReplaceAllString(content, "")

			// Every page must have exactly one H1 heading (the command name).
			h1Count := strings.Count(stripped, "\n# ")
			require.Equal(t, 1, h1Count, "expected exactly one H1 heading")

			// No all-caps SEE ALSO heading should remain.
			require.NotContains(t, content, "## SEE ALSO", "cobra's all-caps SEE ALSO should be normalized")

			// There must be no more than one ## See Also section.
			seeAlsoCount := strings.Count(stripped, "\n## See Also\n")
			require.LessOrEqual(t, seeAlsoCount, 1, "duplicate ## See Also sections found")
		})
	}
}
