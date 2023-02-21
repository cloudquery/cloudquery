package cmd

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func getSyncCommand(t *testing.T, config string) (*cobra.Command, string) {
	t.Helper()

	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	testConfig := path.Join(currentDir, "testdata", config)
	tmpDir := t.TempDir()
	logFileName := path.Join(tmpDir, "cloudquery.log")
	cmd := NewCmdRoot()
	cmd.SetArgs([]string{"sync", testConfig, "--cq-dir", tmpDir, "--log-file-name", logFileName})
	return cmd, tmpDir
}

func TestSync(t *testing.T) {
	configs := []struct {
		name        string
		config      string
		err         string
		logMessages []string
	}{
		{
			name:   "should finish successfully for valid config",
			config: "sync-success.yml",
		},
		{
			name:   "should fail with missing path error when path is missing",
			config: "sync-missing-path-error.yml",
			err:    "Error: failed to validate destination test: path is required",
		},
		{
			name:   "should sync multiple sources and destinations",
			config: "multiple-sources-destinations.yml",
			logMessages: []string{
				`Start sync destinations=\["test\-1.*?","test\-2.*?"\] module=cli source="test\-1.*?"`,
				`Start sync destinations=\["test\-2.*?","test\-1.*?"\] module=cli source="test\-2.*?"`},
		},
	}

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			defer CloseLogFile()
			cmd, cqDir := getSyncCommand(t, tc.config)
			commandError := cmd.Execute()

			// check that log was written and contains some lines from the plugin
			b, logFileError := os.ReadFile(path.Join(cqDir, "cloudquery.log"))
			logContent := string(b)
			require.NoError(t, logFileError, "failed to read cloudquery.log")
			require.NotEmpty(t, logContent, "cloudquery.log empty; expected some logs")

			for _, logMessage := range tc.logMessages {
				require.Regexp(t, logMessage, logContent, "log message not found in cloudquery.log")
			}

			if tc.err == "" {
				require.NoError(t, commandError)
			} else {
				require.Contains(t, commandError.Error(), tc.err)
			}
		})
	}
}
