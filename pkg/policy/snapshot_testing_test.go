package policy

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func testHelperCreatePaths(dir, filePath string) error {
	err := os.MkdirAll(path.Join(dir, filePath[:strings.LastIndex(filePath, "/")]), 0755)
	if err != nil {
		return err
	}
	_, err = os.Create(path.Join(dir, filePath))
	if err != nil {
		return err
	}
	return nil
}

func TestFindAllTables(t *testing.T) {
	tests := []struct {
		paths  []string
		tables []string
		err    error
	}{
		{
			paths:  []string{"path/to/actual/table/table_testTable.sql"},
			tables: []string{"path/to/actual/table/table_testTable.sql"},
		}, {
			paths:  []string{"path/to/fake/table/table1_testTable.sql"},
			tables: []string{},
		},
	}
	for _, test := range tests {
		uniqueTempDir, err := os.MkdirTemp(os.TempDir(), "*-myOptionalSuffix")
		if err != nil {
			t.Fatal(err)
		}

		for _, testPath := range test.paths {
			err = testHelperCreatePaths(uniqueTempDir, testPath)
			if err != nil {
				t.Fatal(err)
			}

		}

		resp, err := FindAllTables(uniqueTempDir)
		errDiff := cmp.Diff(err, test.err, cmpopts.EquateErrors())
		if errDiff != "" {
			t.Fatal(errDiff)
		}

		tables := make([]string, 0)
		for _, table := range test.tables {
			tables = append(tables, path.Join(uniqueTempDir, table))
		}
		respDiff := cmp.Diff(resp, tables, cmpopts.EquateEmpty())
		if respDiff != "" {
			t.Fatal(respDiff)
		}
	}
}
func TestFindAllTestCases(t *testing.T) {

	tests := []struct {
		paths      []string
		foundPaths []string
		err        error
	}{
		{
			err:        nil,
			paths:      []string{"/path/to/actual/test/snapshot_data.csv"},
			foundPaths: []string{"/path/to/actual/test"},
		},
		{
			err:        nil,
			paths:      []string{"/path/to/fake/test/ssnapshot_data.csv"},
			foundPaths: []string{},
		},
		{
			err:        nil,
			paths:      []string{"/path/to/invalid/test/snapshot_data.sql"},
			foundPaths: []string{},
		},
	}
	for _, test := range tests {
		uniqueTempDir, err := os.MkdirTemp(os.TempDir(), "*-myOptionalSuffix")
		if err != nil {
			t.Fatal(err)
		}

		for _, testPath := range test.paths {
			err = testHelperCreatePaths(uniqueTempDir, testPath)
			if err != nil {
				t.Fatal(err)
			}

		}

		resp, err := FindAllTestCases(uniqueTempDir)
		errDiff := cmp.Diff(err, test.err, cmpopts.EquateErrors())
		if errDiff != "" {
			t.Fatal(errDiff)
		}

		fixedPaths := make([]string, 0)
		for _, foundPath := range test.foundPaths {
			fixedPaths = append(fixedPaths, path.Join(uniqueTempDir, foundPath))
		}
		respDiff := cmp.Diff(resp, fixedPaths, cmpopts.EquateEmpty())
		if respDiff != "" {
			t.Fatal(respDiff)
		}
	}
}
