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
			errors: []string{"cloudflare (cloudquery/cloudflare@v6.1.2)", "postgresql (cloudquery/postgresql@v7.3.5)"},
		},
	}
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			testConfig := path.Join(currentDir, "testdata", tc.config)
			cmd := NewCmdRoot()
			baseArgs := testCommandArgs(t)
			cmd.SetArgs(append([]string{"test-connection", testConfig}, baseArgs...))
			err := cmd.Execute()
			if len(tc.errors) > 0 {
				var errs *testConnectionFailures
				require.ErrorAs(t, err, &errs)
				require.Len(t, errs.failed, len(tc.errors))
				for i, want := range tc.errors {
					assert.Equal(t, want, errs.failed[i].PluginRef)
				}
			} else {
				assert.NoError(t, err)
			}

			// check that log was written and contains some lines from the plugin
			b, logFileError := os.ReadFile(baseArgs[3])
			logContent := string(b)
			require.NoError(t, logFileError, "failed to read cloudquery.log")
			require.NotEmpty(t, logContent, "cloudquery.log empty; expected some logs")
		})
	}
}
