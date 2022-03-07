package cmd

import (
	"encoding/json"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestFilterConfigProviders(t *testing.T) {
	testCases := []struct {
		name        string
		input       config.Config
		filter      []string
		expected    config.Config
		expectError string
	}{
		{
			name:     "empty config",
			input:    config.Config{},
			filter:   nil,
			expected: config.Config{},
		},
		{
			name: "config with unmatching filter",
			input: config.Config{
				Providers: []*config.Provider{
					{
						Name:      "aws",
						Resources: []string{"res1", "res2"},
					},
				},
				CloudQuery: config.CloudQuery{
					Providers: []*config.RequiredProvider{
						{
							Name: "aws",
						},
					},
				},
			},
			filter:      []string{"gcp"},
			expectError: "nothing to fetch",
		},
		{
			name: "config with resource filter",
			input: config.Config{
				Providers: []*config.Provider{
					{
						Name:      "aws",
						Resources: []string{"res1", "res2"},
					},
				},
				CloudQuery: config.CloudQuery{
					Providers: []*config.RequiredProvider{
						{
							Name: "aws",
						},
					},
				},
			},
			filter: []string{"aws:res2"},
			expected: config.Config{
				Providers: []*config.Provider{
					{
						Name:      "aws",
						Resources: []string{"res2"},
					},
				},
				CloudQuery: config.CloudQuery{
					Providers: []*config.RequiredProvider{
						{
							Name: "aws",
						},
					},
				},
			},
		},
		{
			name: "config with aliases, one alias and one other provider",
			input: config.Config{
				Providers: []*config.Provider{
					{
						Name:      "aws",
						Resources: []string{"res1", "res2"},
					},
					{
						Name:      "aws",
						Alias:     "aws2",
						Resources: []string{"res2", "res3"},
					},
					{
						Name:      "gcp",
						Resources: []string{"gcpres1", "gcpres2"},
					},
				},
				CloudQuery: config.CloudQuery{
					Providers: []*config.RequiredProvider{
						{
							Name: "aws",
						},
						{
							Name: "gcp",
						},
					},
				},
			},
			filter: []string{"aws2:res2", "gcp"},
			expected: config.Config{
				Providers: []*config.Provider{
					{
						Name:      "aws",
						Alias:     "aws2",
						Resources: []string{"res2"},
					},
					{
						Name:      "gcp",
						Resources: []string{"gcpres1", "gcpres2"},
					},
				},
				CloudQuery: config.CloudQuery{
					Providers: []*config.RequiredProvider{
						{
							Name: "aws",
						},
						{
							Name: "gcp",
						},
					},
				},
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			out := tc.input
			err := filterConfigProviders(tc.filter)(&out)
			if tc.expectError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectError)
				if t.Failed() {
					actual, _ := json.MarshalIndent(out, "", "  ")
					assert.Empty(t, string(actual))
				}
				return
			}
			actual, _ := json.MarshalIndent(out, "", "  ")
			expected, _ := json.MarshalIndent(tc.expected, "", "  ")
			assert.Equal(t, string(expected), string(actual))
		})
	}
}
