package client

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cloudquery/filetypes/v2"
	"github.com/cloudquery/filetypes/v2/csv"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/specs"
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
	TestName string
}

func testSpecsWithoutFormat(t *testing.T) []testSpec {
	return []testSpec{
		{
			TestName: "Directory",
			Spec: Spec{
				Directory: t.TempDir(),
			},
		},
		{
			TestName: "DirectoryWithTable",
			Spec: Spec{
				Directory: path.Join(t.TempDir(), "{{TABLE}}/data.{{FORMAT}}"),
			},
		},
		{
			TestName: "Path",
			Spec: Spec{
				Path: path.Join(t.TempDir(), "{{TABLE}}.{{FORMAT}}"),
			},
		},
		{
			TestName: "PathWithTable",
			Spec: Spec{
				Path: path.Join(t.TempDir(), "{{TABLE}}/data.{{FORMAT}}"),
			},
		},
	}
}

func testSpecs(t *testing.T) []testSpec {
	var ret []testSpec
	formats := testFormats()
	for _, s := range testSpecsWithoutFormat(t) {
		s := s
		s.NoRotate = true
		for i := range formats {
			s2 := s
			s2.TestName += ":" + string(formats[i].Format)
			s2.FileSpec = &formats[i]
			ret = append(ret, s2)
		}
	}

	return ret
}

func TestPlugin(t *testing.T) {
	for _, ts := range testSpecs(t) {
		ts := ts
		t.Run(ts.TestName, func(t *testing.T) {
			testPlugin(t, &ts.Spec)

			dirOrPath := ts.Spec.Directory
			if dirOrPath == "" {
				dirOrPath = ts.Spec.Path
			}
			baseDir := strings.Split(dirOrPath, "/{")[0]

			fi, err := os.Stat(baseDir)
			assert.NoError(t, err)
			assert.Truef(t, fi.IsDir(), "basedir %s is not a directory", baseDir)
			assert.NoError(t, filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
				assert.NoError(t, err)
				if err != nil {
					return err
				}
				t.Log("walking path", path)
				assert.NotContainsf(t, path, "{", "path %s still contains template", path)
				if t.Failed() {
					return fmt.Errorf("test failed")
				}
				return nil
			}))
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
