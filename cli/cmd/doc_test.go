package cmd

import (
	"os"
	"path/filepath"
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

// generateDocs runs the doc command into a temp directory and returns the path.
// Shared by multiple tests so we only generate once per test run.
func generateDocs(t *testing.T) string {
	t.Helper()
	tmpDir := t.TempDir()
	cmd := NewCmdRoot()
	cmd.SetArgs(append([]string{"doc", tmpDir}, testCommandArgs(t)...))
	require.NoError(t, cmd.Execute())
	return tmpDir
}

func TestDoc(t *testing.T) {
	tmpDir := generateDocs(t)

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

// TestDocHeadingStructure verifies the heading hierarchy of every generated
// markdown page:
//   - exactly one H1 (the command name), and it is the first heading
//   - cobra's H3 sections (Synopsis, Options, SEE ALSO) were promoted to H2
//   - cobra's all-caps ## SEE ALSO heading was normalized to ## See Also
//   - no duplicate ## See Also sections
func TestDocHeadingStructure(t *testing.T) {
	tmpDir := generateDocs(t)

	entries, err := os.ReadDir(tmpDir)
	require.NoError(t, err)

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(tmpDir, e.Name()))
		require.NoError(t, err)
		content := string(data)

		t.Run(e.Name(), func(t *testing.T) {
			// Strip fenced code blocks so bash comments (# ...) inside
			// Examples sections are not mistaken for markdown headings.
			stripped := fencedCodeBlock.ReplaceAllString(content, "")

			// Every page must have exactly one H1 heading.
			h1Count := strings.Count(stripped, "\n# ")
			require.Equal(t, 1, h1Count, "expected exactly one H1 heading")

			// The first heading in the document body must be H1, not H2 or H3.
			// (frontmatter ends with "---\n", so the next heading follows a newline)
			afterFrontmatter := stripped[strings.Index(stripped, "---\n")+4:]
			firstHeadingIdx := strings.Index(afterFrontmatter, "\n#")
			require.NotEqual(t, -1, firstHeadingIdx, "no heading found after frontmatter")
			firstHeading := afterFrontmatter[firstHeadingIdx+1:] // skip the leading \n
			require.True(t, strings.HasPrefix(firstHeading, "# "),
				"first heading must be H1, got: %q", strings.SplitN(firstHeading, "\n", 2)[0])

			// cobra's H3 sections must have been promoted to H2. Verify none of
			// the known cobra section names remain at H3 level.
			for _, section := range []string{"Synopsis", "Options", "Examples", "SEE ALSO", "See Also"} {
				require.NotContains(t, stripped, "\n### "+section,
					"cobra section %q was not promoted from H3 to H2", section)
			}

			// No all-caps SEE ALSO heading should remain anywhere.
			require.NotContains(t, content, "## SEE ALSO",
				"cobra's all-caps SEE ALSO should be normalized to ## See Also")

			// There must be at most one ## See Also section (no duplicate).
			seeAlsoCount := strings.Count(stripped, "\n## See Also\n")
			require.LessOrEqual(t, seeAlsoCount, 1, "duplicate ## See Also sections found")
		})
	}
}
