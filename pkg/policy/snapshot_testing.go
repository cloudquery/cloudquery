package policy

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/go-hclog"
)

// pol policy.Policy

func TestPolicy(ctx context.Context, l hclog.Logger, source, snapshotDirectory string) error {
	uniqueTempDir, err := os.MkdirTemp(os.TempDir(), "*-myOptionalSuffix")
	if err != nil {
		l.Error("failed to create tempDirectory", "err", err)
		return err
	}

	conn, err := sdkdb.New(ctx, hclog.NewNullLogger(), "postgres://postgres:pass@localhost:5432/postgres")
	if err != nil {
		l.Error("failed to connect to new database", "err", err)
		return err
	}

	policyManager := NewManager(uniqueTempDir, conn, l)
	p, err := policyManager.Load(ctx, &Policy{Name: "test-policy", Source: source})
	if err != nil {
		l.Error("failed to create policy manager", "err", err)
		return err
	}

	e := NewExecutor(conn, l, nil)
	tests, _ := FindAllTestCases(snapshotDirectory)
	for _, test := range tests {

		selector := strings.TrimPrefix(test, snapshotDirectory+"/")
		selector = strings.Split(selector, "/tests")[0]
		pol := p.Filter(selector)
		tables, _ := FindAllTables(test)
		sort.Sort(sort.Reverse(sort.StringSlice(tables)))

		for _, table := range tables {
			l.Info("restoring table ", "table", table)
			err = e.RestoreSnapshot(ctx, table)
			if err != nil {
				l.Error("failed to restore snapshot", "err", err)
				return err
			}
		}

		// 		c. Run query
		err = e.StoreOutput(ctx, &pol, uniqueTempDir)
		if err != nil {
			l.Error("failed to StoreOutput", "err", err)
			return err
		}
		f2, _ := OpenAndParse(path.Join(test, "data.csv"))
		f1, _ := OpenAndParse(path.Join(uniqueTempDir, "data.csv"))
		if err := compareArbitraryArrays(f1, f2); err != nil {

			l.Error("Failed test case", "case", test, "got", f1, "expected", f2)
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
	csvFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(strings.NewReader(string(csvFile)))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, err

}

func FindAllTestCases(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == "data.csv" {
			files = append(files, strings.TrimSuffix(path, "/data.csv"))
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
