package cmd

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func getTablesCommand(t *testing.T, config string, format, filter string) (*cobra.Command, string) {
	t.Helper()

	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	testConfig := path.Join(currentDir, "testdata", config)
	tmpDir := t.TempDir()
	outputDirectory := path.Join(tmpDir, "cq-docs")
	cmd := NewCmdRoot()
	args := []string{"tables", testConfig, "--output-dir", outputDirectory}
	if format != "" {
		args = append(args, "--format", format)
	}
	if filter != "" {
		args = append(args, "--filter", filter)
	}
	cmd.SetArgs(append(args, testCommandArgs(t)...))
	return cmd, tmpDir
}

func TestTables(t *testing.T) {
	configs := []struct {
		name   string
		config string
		format string
	}{
		{
			name:   "should generate tables in default format",
			config: "multiple-sources.yml",
		},
		{
			name:   "should generate tables in json format",
			config: "multiple-sources.yml",
			format: "json",
		},
		{
			name:   "should generate tables in markdown format",
			config: "multiple-sources.yml",
			format: "markdown",
		},
	}

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			cmd, docsDir := getTablesCommand(t, tc.config, tc.format, "")
			commandError := cmd.Execute()
			require.NoError(t, commandError)

			if tc.format == "markdown" {
				require.FileExists(t, path.Join(docsDir, "cq-docs/test/README.md"))
			} else {
				require.FileExists(t, path.Join(docsDir, "cq-docs/test/__tables.json"))
			}
		})
	}
}

func TestTablesWithFilter(t *testing.T) {
	configs := []struct {
		name   string
		config string
		filter string
	}{
		{
			name:   "should generate tables in default format",
			config: "cloudflare-tables-with-spec-filter.yml",
		},
		{
			name:   "should generate tables in json format",
			config: "cloudflare-tables-with-spec-filter.yml",
			filter: "spec",
		},
	}

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			cmd, docsDir := getTablesCommand(t, tc.config, "json", tc.filter)
			commandError := cmd.Execute()
			require.NoError(t, commandError)
			expectedFile := path.Join(docsDir, "cq-docs/cloudflare/__tables.json")
			require.FileExists(t, expectedFile)
			file, err := os.Open(expectedFile)
			require.NoError(t, err)
			content, err := io.ReadAll(file)
			require.NoError(t, err)
			_ = file.Close()
			type table struct {
				Name string `json:"name"`
			}
			var tables []table
			err = json.Unmarshal(content, &tables)
			require.NoError(t, err)

			if tc.filter == "spec" {
				require.Len(t, tables, 1)
				require.Equal(t, "cloudflare_access_applications", tables[0].Name)
			} else {
				require.Greater(t, len(tables), 1)
			}
		})
	}
}
