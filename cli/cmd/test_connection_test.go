package cmd

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTestConnection(t *testing.T) {
	configs := []struct {
		name   string
		config string
		errors []string
	}{
		{
			name:   "multiple test sources should pass validation",
			config: "multiple-sources.yml",
		},
		{
			name:   "bad AWS and Postgres auth should fail validation",
			config: "test-connection-bad-connection.yml",
			errors: []string{"failed to init source cloudflare", "failed to init destination postgresql"},
		},
	}
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			cqDir := t.TempDir()
			logFileName := path.Join(cqDir, "cloudquery.log")
			t.Cleanup(func() {
				CloseLogFile()
			})
			testConfig := path.Join(currentDir, "testdata", tc.config)
			cmd := NewCmdRoot()
			cmd.SetArgs([]string{"test-connection", testConfig, "--cq-dir", cqDir, "--log-file-name", logFileName})
			err := cmd.Execute()
			if tc.errors != nil {
				for _, e := range tc.errors {
					assert.Contains(t, err.Error(), e)
				}
			} else {
				assert.NoError(t, err)
			}

			// check that log was written and contains some lines from the plugin
			b, logFileError := os.ReadFile(path.Join(cqDir, "cloudquery.log"))
			logContent := string(b)
			require.NoError(t, logFileError, "failed to read cloudquery.log")
			require.NotEmpty(t, logContent, "cloudquery.log empty; expected some logs")
		})
	}
}
