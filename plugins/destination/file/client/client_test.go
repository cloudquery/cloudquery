package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/filetypes/v4/csv"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/assert"
)

func testFormats() []filetypes.FileSpec {
	return []filetypes.FileSpec{
		{
			Format: filetypes.FormatTypeCSV,
			FormatSpec: csv.Spec{
				SkipHeader: true,
				Delimiter:  ",",
			},
		},
		{
			Format: filetypes.FormatTypeJSON,
		},
		//{
		//	Format: filetypes.FormatTypeParquet,
		//},
	}
}

type testSpec struct {
	Spec
	testName string
	baseDir  string
}

func testSpecsWithoutFormat(t *testing.T) []testSpec {
	var (
		ret []testSpec
		bd  string // temp variable to hold tempdir/basedir dir for each test case
	)

	bd = t.TempDir()
	ret = append(ret, testSpec{
		testName: "Directory",
		baseDir:  bd,
		Spec: Spec{
			Directory: bd,
		},
	})

	bd = t.TempDir()
	ret = append(ret, testSpec{
		testName: "DirectoryWithTable",
		baseDir:  bd,
		Spec: Spec{
			Directory: filepath.Join(bd, "{{TABLE}}", "data.{{FORMAT}}"),
		},
	})

	bd = t.TempDir()
	ret = append(ret, testSpec{
		testName: "Path",
		baseDir:  bd,
		Spec: Spec{
			Path: filepath.Join(bd, "{{TABLE}}.{{FORMAT}}"),
		},
	})

	bd = t.TempDir()
	ret = append(ret, testSpec{
		testName: "PathWithTable",
		baseDir:  bd,
		Spec: Spec{
			Path: filepath.Join(bd, "{{TABLE}}", "data.{{FORMAT}}"),
		},
	})

	return ret
}

func testSpecs(t *testing.T) []testSpec {
	var ret []testSpec
	formats := testFormats()
	for _, s := range testSpecsWithoutFormat(t) {
		s := s
		s.NoRotate = true
		for i := range formats {
			s2 := s
			s2.testName += ":" + string(formats[i].Format)
			s2.FileSpec = &formats[i]
			ret = append(ret, s2)
		}
	}

	return ret
}

func TestPlugin(t *testing.T) {
	for _, ts := range testSpecs(t) {
		ts := ts
		t.Run(ts.testName, func(t *testing.T) {
			testPlugin(t, &ts.Spec)

			fi, err := os.Stat(ts.baseDir)
			assert.NoError(t, err)
			assert.Truef(t, fi.IsDir(), "basedir %s is not a directory", ts.baseDir)

			fileCount := 0
			assert.NoError(t, filepath.WalkDir(ts.baseDir, func(path string, d os.DirEntry, err error) error {
				assert.NoError(t, err)
				if err != nil {
					return err
				}
				t.Log("walking path", path)
				if !d.IsDir() {
					fileCount++
				}
				if !assert.NotContainsf(t, path, "{", "path %s still contains template", path) {
					return fmt.Errorf("test failed")
				}
				return nil
			}))

			assert.NotZero(t, fileCount, "no files written to %s", ts.baseDir)
		})
	}
}

func testPlugin(t *testing.T, spec *Spec) {
	ctx := context.Background()
	p := plugin.NewPlugin("file", "development", New)
	b, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, b); err != nil {
		t.Fatal(err)
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.PluginTestSuiteTests{
			SkipUpsert:      true,
			SkipMigrate:     true,
			SkipDeleteStale: true,
		},
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			TimePrecision: time.Millisecond,
		}),
	)
}
