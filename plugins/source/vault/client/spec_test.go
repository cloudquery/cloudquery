package client

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSpec_SetDefaults(t *testing.T) {
	tests := map[string]struct {
		concurrency int
		expected    int
	}{
		"Should use a default if no value is supplied": {
			concurrency: 0,
			expected:    10000,
		},
		"Should use the configured value if supplied": {
			concurrency: 500,
			expected:    500,
		},
	}

	for desc, tC := range tests {
		t.Run(desc, func(t *testing.T) {
			s := Spec{Concurrency: tC.concurrency}

			s.SetDefaults()

			assert.Equal(t, tC.expected, s.Concurrency)
		})
	}
}

func TestSpec_Validate(t *testing.T) {
	tests := map[string]struct {
		address       string
		expectedError error
	}{
		"A valid URL should not return an error": {
			address: "http://localhost:8200",
		},
		"An invalid URL should return an error": {
			address:       "bad_address",
			expectedError: errors.New("invalid vault address provided \"bad_address\": parse \"bad_address\": invalid URI for request"),
		},
		"An empty URL should return an error": {
			address:       "",
			expectedError: errors.New("no vault address provided"),
		},
	}

	for desc, tC := range tests {
		t.Run(desc, func(t *testing.T) {
			s := Spec{
				VaultAddress: tC.address,
			}

			err := s.Validate()

			if tC.expectedError != nil {
				require.EqualError(t, err, tC.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
