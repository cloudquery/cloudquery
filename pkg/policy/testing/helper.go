package testing

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sync"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/cloudquery/cq-provider-sdk/testlog"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

var (
	dbConnOnce sync.Once
	pool       *pgxpool.Pool
	dbErr      error
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func cleanDatabase(ctx context.Context, conn *pgxpool.Conn) error {
	_, err := conn.Exec(ctx, `DROP SCHEMA public CASCADE;
	CREATE SCHEMA public;
	GRANT ALL ON SCHEMA public TO postgres;
	GRANT ALL ON SCHEMA public TO public`)
	return err
}

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
	// _, err = console.FilterPolicies(source, pols)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	policyManager := policy.NewManager(uniqueTempDir, nil, l)
	p, err := policyManager.Load(ctx, &policy.Policy{Name: "test-policy", Source: source})
	if err != nil {
		t.Fatal(err)
	}
	pol := p.Filter(selector)

	pool, err := setupDatabase()
	if err != nil {
		t.Fatal(err)
	}

	e := policy.NewExecutor(nil, l, nil)
	config, err := pgxpool.ParseConfig(pool.Config().ConnString())
	if err != nil {
		t.Fatalf("Error parsing config: %+v", err)
	}
	conn, err := pool.Acquire(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Release()

	testPath := path.Join(snapshotDirectory, "query-"+pol.Checks[0].Name, "tests", "")
	tests, _ := FindAllTestCases(testPath)
	for _, test := range tests {

		err = cleanDatabase(ctx, conn)
		if err != nil {
			t.Fatal(err)
		}

		fileP := path.Join(testPath, test, "pg-dump.sql")
		err = e.RestoreSnapshot(ctx, fileP, config)
		if err != nil {
			t.Fatal(err)
		}
		if err != nil {
			t.Fatalf("Error creating config: %+v", err)
		}

		if err != nil {
			t.Fatalf("Error creating temp dir: %+v", err)
		}
		// 		c. Run query
		err = e.StoreOutput(ctx, &pol, uniqueTempDir, config)
		if err != nil {
			t.Fatalf("Error storing output: %+v", err)
		}
		f2, _ := OpenAndParse(path.Join(testPath, test, "data.json"))
		f1, _ := OpenAndParse(path.Join(uniqueTempDir, "data.json"))
		compareArbitraryArrays(t, f1, f2)
	}

}

func compareArbitraryArrays(t *testing.T, f1, f2 []map[string]interface{}) {
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
func OpenAndParse(filePath string) ([]map[string]interface{}, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, err
	}
	return result, err

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

func setupDatabase() (*pgxpool.Pool, error) {
	dbConnOnce.Do(func() {
		var dbCfg *pgxpool.Config
		dbCfg, dbErr = pgxpool.ParseConfig(getEnv("DATABASE_URL", "host=localhost user=postgres password=pass database=postgres port=5432"))
		if dbErr != nil {
			return
		}
		ctx := context.Background()
		dbCfg.MaxConns = 15
		dbCfg.LazyConnect = true
		pool, dbErr = pgxpool.ConnectConfig(ctx, dbCfg)
	})
	return pool, dbErr

}
