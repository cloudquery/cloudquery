package cmd

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/google/uuid"
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
			errors: []string{"cloudflare (cloudquery/cloudflare@v9.6.0)", "postgresql (cloudquery/postgresql@v8.6.2)"},
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

func TestTestConnection_IsolatedPluginEnvironmentsInCloud(t *testing.T) {
	configs := []struct {
		name   string
		config string
		err    string
	}{
		{
			name:   "source-with-env",
			config: "source-with-env.yml",
		},
	}
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)

	t.Setenv("CLOUDQUERY_API_KEY", "cqstc_123")
	t.Setenv("CQ_CLOUD", "1")
	t.Setenv("_CQ_TEAM_NAME", "test_team")
	t.Setenv("_CQ_SYNC_NAME", "test_sync")
	t.Setenv("_CQ_SYNC_RUN_ID", uuid.Must(uuid.NewUUID()).String())
	t.Setenv("__SOURCE_TEST__TEST_KEY", "test_value")
	t.Setenv("NOT_TEST_ENV", "should_not_be_visible_to_plugin")

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			testConfig := path.Join(currentDir, "testdata", tc.config)
			cmd := NewCmdRoot()
			cmd.SetArgs(append([]string{"test-connection", testConfig}, testCommandArgs(t)...))
			err := cmd.Execute()
			if tc.err != "" {
				assert.Contains(t, err.Error(), tc.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
