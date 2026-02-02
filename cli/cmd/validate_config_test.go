package cmd

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateConfig(t *testing.T) {
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
			config: "validate-config-error.yml",
			errors: []string{"failed to validate source config cloudflare", "failed to validate destination config postgresql"},
		},
	}
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			cmd := NewCmdRoot()
			testConfig := path.Join(currentDir, "testdata", tc.config)
			baseArgs := testCommandArgs(t)

			args := append([]string{"validate-config", testConfig}, baseArgs...)
			cmd.SetArgs(args)
			err := cmd.Execute()
			// check that log was written and contains some lines from the plugin
			b, logFileError := os.ReadFile(baseArgs[3])
			logContent := string(b)
			require.NoError(t, logFileError, "failed to read cloudquery.log")
			require.NotEmpty(t, logContent, "cloudquery.log empty; expected some logs")
			require.NotContains(t, logContent, "skipping validation")

			if tc.errors != nil {
				for _, e := range tc.errors {
					require.Contains(t, err.Error(), e)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}
