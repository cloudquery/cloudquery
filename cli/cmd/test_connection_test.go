package cmd

import (
	"errors"
	"os"
	"path"
	"runtime"
	"strings"
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
			errors: []string{"cloudquery/cloudflare", "cloudquery/postgresql"},
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
				var errs *testConnectionFailureErrors
				if errors.As(err, &errs) {
					AssertErrorsContainAny(t, errs.Unwrap(), tc.errors)
				} else {
					t.Errorf("Expected error to be of type testConnectionFailureErrors, got %v", err)
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

// ErrorsContainAny checks if any error in a slice contains at least one of the substrings.
func ErrorsContainAny(errors []error, substrings []string) bool {
	for _, err := range errors {
		if err != nil {
			errMsg := err.Error()
			for _, substr := range substrings {
				if strings.Contains(errMsg, substr) {
					return true
				}
			}
		}
	}
	return false
}

// AssertErrorsContainAny asserts that at least one error message in the slice contains at least one of the substrings.
func AssertErrorsContainAny(t *testing.T, errors []error, substrings []string) {
	if !ErrorsContainAny(errors, substrings) {
		t.Errorf("Expected at least one error in %v to contain at least one of %v", errors, substrings)
	}
}
