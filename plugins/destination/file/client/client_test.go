package client

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudquery/filetypes/v3"
	"github.com/cloudquery/filetypes/v3/csv"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
	"github.com/stretchr/testify/assert"
)

var migrateStrategy = destination.MigrateStrategy{
	AddColumn:           specs.MigrateModeForced,
	AddColumnNotNull:    specs.MigrateModeForced,
	RemoveColumn:        specs.MigrateModeForced,
	RemoveColumnNotNull: specs.MigrateModeForced,
	ChangeColumn:        specs.MigrateModeForced,
}

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
		{
			Format: filetypes.FormatTypeParquet,
		},
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
				assert.NotContainsf(t, path, "{", "path %s still contains template", path)
				if t.Failed() {
					return fmt.Errorf("test failed")
				}
				return nil
			}))

			assert.NotZero(t, fileCount, "no files written to %s", ts.baseDir)
		})
	}
}

func testPlugin(t *testing.T, spec *Spec) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("file", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: spec,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipSecondAppend:          true,
			SkipDeleteStale:           true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,

			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		},
	)
}
