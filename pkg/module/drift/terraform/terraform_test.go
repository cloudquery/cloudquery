package terraform

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadState(t *testing.T) {
	rdr := strings.NewReader(`{
  "version": 4,
  "terraform_version": "1.0.10",
  "serial": 9,
  "lineage": "",
  "outputs": {},
  "resources": []
}`)
	data, err := LoadState(rdr)
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestValidateStateVersion(t *testing.T) {
	testCases := []struct {
		input         string
		expectedOK    bool
		expectedError string
	}{
		{
			input:      `{"version": 4}`,
			expectedOK: true,
		},
		{
			input:         `{"version": 3}`,
			expectedOK:    false,
			expectedError: "unsupported tfstate version",
		},
		{
			input:         `{}`,
			expectedOK:    true,
			expectedError: "unspecified tfstate version",
		},
		{
			input:         `{"version": "mama"}`,
			expectedOK:    true,
			expectedError: "unknown tfstate version",
		},
	}
	for i := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			tc := testCases[i]
			data, err := LoadState(strings.NewReader(tc.input))
			assert.NoError(t, err)
			assert.NotNil(t, data)
			ok, err := ValidateStateVersion(data)
			assert.Equal(t, tc.expectedOK, ok)
			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			}
		})
	}
}
