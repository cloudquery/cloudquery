package policy

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type sourceTest struct {
	Name          string
	Source        string
	Expected      bool
	ExpectedMeta  *Meta
	ErrorExpected bool
}

func init() {
	_, filename, _, _ := runtime.Caller(0)
	err := os.Chdir(path.Dir(filename))
	if err != nil {
		panic(err)
	}
}

func TestLoadSource(t *testing.T) {
	sourceTests := []sourceTest{
		{
			Name:         "local_directory",
			Source:       "tests/local",
			Expected:     true,
			ExpectedMeta: &Meta{Type: "file", Version: "", SubPolicy: "", Directory: "tests/output/local"},
		},

		{
			Name:          "non-existing_directory",
			Source:        "tests/not_exist",
			ErrorExpected: true,
		},
		{
			Name:         "hub",
			Source:       "aws",
			Expected:     true,
			ExpectedMeta: &Meta{Type: "hub", Version: "", SubPolicy: "", Directory: "tests/output/github.com/cloudquery-policies/aws"},
		},
		{
			Name:         "hub with ref",
			Source:       "aws?ref=v0.1.0",
			Expected:     true,
			ExpectedMeta: &Meta{Type: "hub", Version: "", SubPolicy: "", Directory: "tests/output/github.com/cloudquery-policies/aws"},
		},
		{
			Name:         "hub with @",
			Source:       "aws@v0.1.0",
			Expected:     true,
			ExpectedMeta: &Meta{Type: "hub", Version: "", SubPolicy: "", Directory: "tests/output/github.com/cloudquery-policies/aws"},
		},
		{
			Name:         "github",
			Source:       "github.com/cloudquery-policies/aws",
			Expected:     true,
			ExpectedMeta: &Meta{Type: "github", Version: "", SubPolicy: "", Directory: "tests/output/github.com/cloudquery-policies/aws"},
		},
		{
			Name:          "non-existing-github",
			Source:        "user:pass@github.com/cloudquery-policies/blabla",
			ErrorExpected: true,
		},
		{
			Name:         "force github link",
			Source:       "git::https://github.com/cloudquery-policies/aws.git?ref=v0.0.1",
			Expected:     true,
			ExpectedMeta: &Meta{Type: "git", Version: "v0.0.1", SubPolicy: "", Directory: "tests/output/github.com/cloudquery-policies/aws"},
		},
		{
			Name:          "https github link no force",
			Source:        "https://github.com/cloudquery-policies/aws?ref=v0.0.1",
			ErrorExpected: true,
		},
		{
			Name:         "github reference specific commit",
			Source:       "github.com/cloudquery-policies/aws?ref=68bede6e",
			Expected:     true,
			ExpectedMeta: &Meta{Type: "github", Version: "68bede6e", SubPolicy: "", Directory: "tests/output/github.com/cloudquery-policies/aws"},
		},

		{
			Name:         "subpolicy path",
			Source:       "github.com/cloudquery-policies/aws//subpolicy?ref=68bede6e",
			Expected:     true,
			ExpectedMeta: &Meta{Type: "github", Version: "68bede6e", SubPolicy: "subpolicy", Directory: "tests/output/github.com/cloudquery-policies/aws"},
		},
		{
			Name:         "subpolicy path with specific rev using @",
			Source:       "github.com/cloudquery-policies/aws//cis_v1.2.0@v0.1.0",
			Expected:     true,
			ExpectedMeta: &Meta{Type: "github", Version: "v0.1.0", SubPolicy: "cis_v1.2.0", Directory: "tests/output/github.com/cloudquery-policies/aws"},
		},
	}

	for _, s := range sourceTests {
		t.Run(s.Name, func(t *testing.T) {
			data, meta, err := LoadSource(context.Background(), "tests/output", s.Source)
			if s.ErrorExpected {
				require.Error(t, err)
				return
			}
			require.Nil(t, err)
			assert.Equal(t, s.ExpectedMeta.Type, meta.Type)
			assert.Equal(t, filepath.ToSlash(s.ExpectedMeta.Directory), filepath.ToSlash(meta.Directory), "unexpected saved policy directory")
			assert.NotNil(t, data)
		})
	}
}

func TestDetectPolicy(t *testing.T) {
	p, found, err := DetectPolicy("aws", "")
	require.Nil(t, err)
	assert.Equal(t, "aws", p.Source)
	assert.True(t, found)
	_, found, err = DetectPolicy("not-exist", "")
	require.Nil(t, err)
	assert.False(t, found)
}
