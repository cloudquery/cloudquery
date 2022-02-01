package testing

import (
	"context"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/cloudquery/cq-provider-sdk/testlog"
	"github.com/hashicorp/go-hclog"
)

// pol policy.Policy
func TestPolicy(t *testing.T, source, selector, snapshotDirectory string) {
	// var pols policy.Policies
	// pols = append(pols, &policy.Policy{Name: "aws", Source: source})
	// Setup dependencies
	t.Helper()
	uniqueTempDir, err := os.MkdirTemp(os.TempDir(), "*-myOptionalSuffix")
	if err != nil {
		t.Fatal(err)
	}
	l := testlog.New(t)
	l.SetLevel(hclog.Info)

	ctx := context.Background()
	conn, err := sdkdb.New(ctx, hclog.NewNullLogger(), "postgres://postgres:pass@localhost:5432/postgres")
	if err != nil {
		t.Fatal(err)
	}

	policyManager := policy.NewManager(uniqueTempDir, conn, l)
	p, err := policyManager.Load(ctx, &policy.Policy{Name: "test-policy", Source: source})
	if err != nil {
		t.Fatal(err)
	}
	pol := p.Filter(selector)

	clie := policy.NewExecutor(conn, l, nil)

	testPath := path.Join(snapshotDirectory, "query-"+pol.Checks[0].Name, "tests", "")
	tests, _ := FindAllTestCases(testPath)
	for _, test := range tests {
		tables, _ := FindAllTables(path.Join(testPath, test))
		sort.Sort(sort.Reverse(sort.StringSlice(tables)))

		for _, table := range tables {
			l.Info("restoring table ", "table", table)
			err = clie.RestoreSnapshot(ctx, table)
			if err != nil {
				t.Fatal(err)
			}
		}

		if err != nil {
			t.Fatalf("Error creating temp dir: %+v", err)
		}

		// 		c. Run query
		err = clie.StoreOutput(ctx, &pol, uniqueTempDir)
		if err != nil {
			t.Fatalf("Error storing output: %+v", err)
		}
		f2, _ := OpenAndParse(path.Join(testPath, test, "data.csv"))
		f1, _ := OpenAndParse(path.Join(uniqueTempDir, "data.csv"))
		compareArbitraryArrays(t, f1, f2)
	}

}

func compareArbitraryArrays(t *testing.T, f1, f2 [][]string) {
	assert.Equal(t, len(f1), len(f2), "Query results should have same number of items.")
	for _, item1 := range f1 {
		diffItemPresent := false
		for _, item2 := range f2 {
			diff := cmp.Diff(item1, item2)
			if diff == "" {
				diffItemPresent = true
			}
		}
		if !diffItemPresent {
			t.Fatalf("Item %+v, not found", item1)
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
			t.Fatalf("Item %+v, not found", item1)
		}
	}

}
func OpenAndParse(filePath string) ([][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(csvFile)
	csvFile.Close()
	r := csv.NewReader(strings.NewReader(string(byteValue)))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records, err

}

func FindAllTestCases(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != root {
			files = append(files, info.Name())
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
