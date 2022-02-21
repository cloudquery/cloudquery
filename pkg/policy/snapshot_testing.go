package policy

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/google/go-cmp/cmp"
)

func (p *Policy) Test(ctx context.Context, e *Executor, source, snapshotDirectory, tempDirectory string) error {

	tests, err := FindAllTestCases(snapshotDirectory)
	if err != nil {
		return err
	}
	for _, test := range tests {

		selector := strings.TrimPrefix(test, snapshotDirectory+"/")
		selector = strings.Split(selector, "/tests")[0]
		pol := p.Filter(selector)
		tables, _ := FindAllTables(test)
		sort.Sort(sort.Reverse(sort.StringSlice(tables)))

		for _, table := range tables {
			e.log.Info("restoring table ", "table", table)
			err = RestoreSnapshot(ctx, e.conn, e.log, table)
			if err != nil {
				e.log.Error("failed to restore snapshot", "err", err)
				return err
			}
		}

		err = StoreOutput(ctx, e, &pol, tempDirectory)
		if err != nil {
			e.log.Error("failed to StoreOutput", "err", err)
			return err
		}
		f2, _ := OpenAndParse(path.Join(test, "snapshot_data.csv"))
		f1, _ := OpenAndParse(path.Join(tempDirectory, "snapshot_data.csv"))
		if err := compareArbitraryArrays(f1, f2); err != nil {

			e.log.Error("Failed test case", "case", test, "got", f1, "expected", f2)
			return err
		}
	}

	return nil
}

func compareArbitraryArrays(f1, f2 [][]string) error {

	if len(f1) != len(f2) {
		return errors.New("query results should have same number of items")
	}
	for _, item1 := range f1 {
		diffItemPresent := false
		for _, item2 := range f2 {
			diff := cmp.Diff(item1, item2)
			if diff == "" {
				diffItemPresent = true
			}
		}
		if !diffItemPresent {
			return fmt.Errorf("item %+v, not found", item1)
		}
	}
	for item1 := range f2 {
		diffItemPresent := false
		for item2 := range f1 {
			diff := cmp.Diff(item1, item2)
			if diff == "" {
				diffItemPresent = true
			}
		}
		if !diffItemPresent {
			return fmt.Errorf("item %+v, not found", item1)
		}
	}
	return nil

}
func OpenAndParse(filePath string) ([][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, err
}

func FindAllTestCases(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == "snapshot_data.csv" {
			files = append(files, strings.TrimSuffix(path, "/snapshot_data.csv"))
		}
		return nil
	})
	return files, err
}

func FindAllTables(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasPrefix(info.Name(), "table_") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
