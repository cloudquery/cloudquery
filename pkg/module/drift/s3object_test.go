package drift

import (
	"sort"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"
)

type mockS3Client struct {
	s3iface.S3API
}

func (m *mockS3Client) ListObjectsV2Pages(_ *s3.ListObjectsV2Input, fn func(*s3.ListObjectsV2Output, bool) bool) error {
	objs := []string{
		"path/to/object.tfstate",
		"a/path/2021-11-15/object.tfstate",
		"a/path/2021-11-16/object.tfstate",
		"a/path/2021-11-17/object.tfstate",
		"a/path/2021-11-21/object.tfstate",
		"a/path/2021-11-17/object.gz",
		"a/path/drift.gz",
	}

	ret := &s3.ListObjectsV2Output{}
	for i := range objs {
		ret.Contents = append(ret.Contents, &s3.Object{Key: aws.String(objs[i])})
	}
	_ = fn(ret, false)
	return nil
}

func TestS3Glob(t *testing.T) {
	table := []struct {
		name              string
		pattern           string
		expected          []string
		expectedSameInput bool
	}{
		{
			name:              "single-file",
			pattern:           "a/path/drift.gz",
			expectedSameInput: true,
		},
		{
			name:    "all-ext",
			pattern: "**/*.tfstate",
			expected: []string{
				"path/to/object.tfstate",
				"a/path/2021-11-15/object.tfstate",
				"a/path/2021-11-16/object.tfstate",
				"a/path/2021-11-17/object.tfstate",
				"a/path/2021-11-21/object.tfstate",
			},
		},
		{
			name:    "prefix-ext",
			pattern: "a/**/*.tfstate",
			expected: []string{
				"a/path/2021-11-15/object.tfstate",
				"a/path/2021-11-16/object.tfstate",
				"a/path/2021-11-17/object.tfstate",
				"a/path/2021-11-21/object.tfstate",
			},
		},
		{
			name:    "prefix-ext-1",
			pattern: "a/**/2021-11-1*/*.tfstate",
			expected: []string{
				"a/path/2021-11-15/object.tfstate",
				"a/path/2021-11-16/object.tfstate",
				"a/path/2021-11-17/object.tfstate",
			},
		},
		{
			name:    "prefix-ext-singlefirst",
			pattern: "a/*/2021**.tfstate",
			expected: []string{
				"a/path/2021-11-15/object.tfstate",
				"a/path/2021-11-16/object.tfstate",
				"a/path/2021-11-17/object.tfstate",
				"a/path/2021-11-21/object.tfstate",
			},
		},
		{
			name:     "prefix-ext-none",
			pattern:  "a/**/2021-11-3*/*.tfstate",
			expected: []string{},
		},
		{
			name:              "none",
			pattern:           "path/to/object.tfstate.gz",
			expectedSameInput: true,
		},
	}

	mockS3 := &mockS3Client{}
	for _, tt := range table {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			matches, err := globS3(mockS3, "", tt.pattern)
			assert.NoError(t, err)
			if tt.expectedSameInput {
				// this is when the pattern doesn't contain any stars
				assert.EqualValues(t, []string{tt.pattern}, matches)
				return
			}

			sort.Strings(matches)
			sort.Strings(tt.expected)
			assert.EqualValues(t, tt.expected, matches)
		})
	}
}
