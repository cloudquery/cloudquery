package getmodules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitHubDetector_Detect(t *testing.T) {
	cases := []struct {
		Name           string
		Source         string
		ExpectedSource string
		ExpectedFound  bool
		ExpectedError  error
	}{
		{
			Name:           "base",
			Source:         "github.com/cloudquery-policies/test_policy",
			ExpectedSource: "git::https://github.com/cloudquery-policies/test_policy.git?ref=v0.0.1",
			ExpectedFound:  true,
			ExpectedError:  nil,
		},
		{
			Name:           "base",
			Source:         "github.com/cloudquery-policies/aws?ref=v0.0.1",
			ExpectedSource: "git::https://github.com/cloudquery-policies/aws.git?ref=v0.0.1",
			ExpectedFound:  true,
			ExpectedError:  nil,
		},
		{
			Name:           "base",
			Source:         "github.com/cloudquery-policies/aws?ref=v0.0.1&other=something",
			ExpectedSource: "git::https://github.com/cloudquery-policies/aws.git?ref=v0.0.1&other=something",
			ExpectedFound:  true,
			ExpectedError:  nil,
		},
	}

	detector := new(GitHubDetector)
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			source, found, err := detector.Detect(tc.Source, "")
			assert.Equal(t, tc.ExpectedSource, source)
			assert.Equal(t, tc.ExpectedFound, found)
			assert.Equal(t, tc.ExpectedError, err)
		})
	}

}
