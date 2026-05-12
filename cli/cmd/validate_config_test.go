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

func TestValidateConfigSchemasDir(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	schemasDir := path.Join(currentDir, "testdata", "schemas-dir")

	t.Run("good spec validates offline without plugin spawn", func(t *testing.T) {
		cmd := NewCmdRoot()
		testConfig := path.Join(currentDir, "testdata", "validate-config-schemas-dir.yml")
		baseArgs := testCommandArgs(t)

		args := append([]string{"validate-config", testConfig, "--schemas-dir", schemasDir}, baseArgs...)
		cmd.SetArgs(args)
		err := cmd.Execute()
		require.NoError(t, err)

		b, logFileError := os.ReadFile(baseArgs[3])
		require.NoError(t, logFileError, "failed to read cloudquery.log")
		logContent := string(b)
		require.Contains(t, logContent, "Validating source against local schema")
		require.Contains(t, logContent, "Validating destination against local schema")
		// No plugin spawn happened, so no "Initializing source/destination" lines.
		require.NotContains(t, logContent, "Initializing source")
		require.NotContains(t, logContent, "Initializing destination")
	})

	t.Run("spec violating schema fails offline", func(t *testing.T) {
		cmd := NewCmdRoot()
		testConfig := path.Join(currentDir, "testdata", "validate-config-schemas-dir-bad.yml")
		baseArgs := testCommandArgs(t)

		args := append([]string{"validate-config", testConfig, "--schemas-dir", schemasDir}, baseArgs...)
		cmd.SetArgs(args)
		err := cmd.Execute()
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to validate source config src")
	})
}

func TestLookupSchemaFile(t *testing.T) {
	dir := t.TempDir()
	target := path.Join(dir, "aws.json")
	require.NoError(t, os.WriteFile(target, []byte("{}"), 0o644))

	require.Equal(t, target, lookupSchemaFile(dir, "aws"))
	require.Equal(t, "", lookupSchemaFile(dir, "gcp"))
	require.Equal(t, "", lookupSchemaFile("", "aws"))
}
