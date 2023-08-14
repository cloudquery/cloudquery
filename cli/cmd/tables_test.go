package cmd

import (
	"path"
	"runtime"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func getTablesCommand(t *testing.T, config string, format string) (*cobra.Command, string) {
	t.Helper()

	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	testConfig := path.Join(currentDir, "testdata", config)
	tmpDir := t.TempDir()
	logFileName := path.Join(tmpDir, "cloudquery.log")
	outputDirectory := path.Join(tmpDir, "cq-docs")
	cmd := NewCmdRoot()
	args := []string{"tables", testConfig, "--cq-dir", tmpDir, "--log-file-name", logFileName, "--output-dir", outputDirectory}
	if format != "" {
		args = append(args, "--format", format)
	}
	cmd.SetArgs(args)
	return cmd, tmpDir
}

func TestTables(t *testing.T) {
	configs := []struct {
		name   string
		config string
		format string
		err    string
	}{
		{
			name:   "should generate tables in default format",
			config: "multiple-sources.yml",
			err:    "the CLI tables command is not supported for sources with protocol version 3. Please upvote https://github.com/cloudquery/cloudquery/issues/12270 if you need this functionality",
		},
		{
			name:   "should generate tables in json format",
			config: "multiple-sources.yml",
			format: "json",
			err:    "the CLI tables command is not supported for sources with protocol version 3. Please upvote https://github.com/cloudquery/cloudquery/issues/12270 if you need this functionality",
		},
		{
			name:   "should generate tables in markdown format",
			config: "multiple-sources.yml",
			format: "markdown",
			err:    "the CLI tables command is not supported for sources with protocol version 3. Please upvote https://github.com/cloudquery/cloudquery/issues/12270 if you need this functionality",
		},
	}

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			defer CloseLogFile()
			cmd, _ := getTablesCommand(t, tc.config, tc.format)
			commandError := cmd.Execute()
			require.Error(t, commandError, tc.err)

			// TODO: Enabled this once https://github.com/cloudquery/cloudquery/issues/12270 is resolved
			// if tc.format == "markdown" {
			// 	require.FileExists(t, path.Join(cqDir, "cq-docs/test/README.md"))
			// } else {
			// 	require.FileExists(t, path.Join(cqDir, "cq-docs/test/__tables.json"))
			// }
		})
	}
}
