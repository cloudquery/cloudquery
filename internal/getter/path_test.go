package getter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizePath(t *testing.T) {

	tt := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "simple",
			path:     "output/test",
			expected: "output/test",
		},
		{
			name:     "http",
			path:     "https://github.com/cloudquery-policies/aws?ref=v0.0.1",
			expected: "github.com/cloudquery-policies/aws",
		},
		{
			name:     "force git github",
			path:     "git::github.com/cloudquery-policies/aws?ref=v0.0.1",
			expected: "github.com/cloudquery-policies/aws",
		},
		{
			name:     "force git https",
			path:     "git::https://github.com/cloudquery-policies/aws?ref=v0.0.1",
			expected: "github.com/cloudquery-policies/aws",
		},
		{
			name:     "hub",
			path:     "aws",
			expected: "aws",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NormalizePath(tc.path))
		})
	}
}
