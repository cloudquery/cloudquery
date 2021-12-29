package policy

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/cloudquery/cloudquery/internal/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	err := os.Chdir(path.Dir(filename))
	if err != nil {
		panic(err)
	}
}

type sourceTest struct {
	Name          string
	Source        string
	Expected      string
	ExpectedMeta  *Meta
	ErrorExpected bool
}

func TestLoadSource(t *testing.T) {
	sourceTests := []sourceTest{
		{
			Name:         "local_directory",
			Source:       "tests/local",
			Expected:     "d05b984bc7837467dceda36a8598c600f3fb624ca24a9337f5a890dab0927662",
			ExpectedMeta: &Meta{Type: "local", Version: "", SubPath: "", Directory: "tests\\output\\local"},
		},

		{
			Name:          "non-existing_directory",
			Source:        "tests/not_exist",
			ErrorExpected: true,
		},
		{
			Name:         "hub",
			Source:       "aws",
			Expected:     "94bd44e6f17851a2dd41d1df683724bb932ca3bf6aa6400986294194df4022b6",
			ExpectedMeta: &Meta{Type: "hub", Version: "", SubPath: "", Directory: "tests/output/aws"},
		},
		{
			Name:         "github",
			Source:       "github.com/cloudquery-policies/aws",
			Expected:     "94bd44e6f17851a2dd41d1df683724bb932ca3bf6aa6400986294194df4022b6",
			ExpectedMeta: &Meta{Type: "github", Version: "", SubPath: "", Directory: "tests\\output\\github.com\\cloudquery-policies\\aws"},
		},
		{
			Name:          "non-existing-github",
			Source:        "github.com/cloudquery-policies/blabla",
			ErrorExpected: true,
		},
	}

	for _, s := range sourceTests {
		t.Run(s.Name, func(t *testing.T) {
			data, meta, err := LoadSource(context.Background(), "tests/output", s.Source)
			if s.ErrorExpected {
				require.Error(t, err)
				return
			} else {
				require.Nil(t, err)
			}
			assert.Equal(t, s.ExpectedMeta.Type, meta.Type)
			assert.Equal(t, filepath.ToSlash(s.ExpectedMeta.Directory), filepath.ToSlash(meta.Directory), "unexpected saved policy directory")
			assert.Equal(t, s.Expected, hash.SHA256(data))
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
