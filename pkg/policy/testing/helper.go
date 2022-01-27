package testing

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/cloudquery/cq-provider-sdk/testlog"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

// Inputs:
// 	- Directory
// 	- Database DSN

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

func TestResource(t *testing.T, policy *policy.Policy) {
	t.Helper()

	// No need for configuration or db connection, get it out of the way first
	// testTableIdentifiersForProvider(t, resource.Provider)

	pool, err := setupDatabase()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Release()

	l := testlog.New(t)
	l.SetLevel(hclog.Debug)

	// 1. Walk database directory
	// 	For each sub-directory:
	// 		a. Clean DB
	// DROP SCHEMA public CASCADE;
	// CREATE SCHEMA public;
	// GRANT ALL ON SCHEMA public TO postgres;
	// GRANT ALL ON SCHEMA public TO public;
	// 		b. Restore Database in .sql file
	// 		c. Run query
	// 		d. Compare values in output.json

}

func FilterFiles(files []string) map[string][]string {
	t := make(map[string][]string)
	for _, file := range files {
		specific := strings.Split(file, "tests")[1]
		test := strings.Split(specific, "/")
		log.Println(test[2])
		log.Println(test, len(test))
		if _, ok := t[test[0]]; !ok {
			t[test[0]] = []string{}
		}

		// t[test[0]] = append(t[test[0]], test[1])
		log.Println(t[test[0]])
	}
	return t
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func setupDatabase() (*pgxpool.Pool, error) {
	dbConnOnce.Do(func() {
		var dbCfg *pgxpool.Config
		dbCfg, dbErr = pgxpool.ParseConfig(getEnv("DATABASE_URL", "host=localhost user=postgres password=pass DB.name=postgres port=5432"))
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

func TestPolicy(t *testing.T, pol policy.Policy) {
	t.Helper()
	ctx := context.Background()

	l := testlog.New(t)
	l.SetLevel(hclog.Debug)
	pool, err := setupDatabase()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Release()

	l := testlog.New(t)
	l.SetLevel(hclog.Debug)
	resource.Provider.Logger = l
	tableCreator := provider.NewTableCreator(l)
	if err := tableCreator.CreateTable(context.Background(), conn, resource.Table, nil); err != nil {
		assert.FailNow(t, fmt.Sprintf("failed to create tables %s", resource.Table.Name), err)
	}

	if err := deleteTables(conn, resource.Table); err != nil {
		t.Fatal(err)
	}

	if err = fetch(t, &resource); err != nil {
		t.Fatal(err)
	}

	verifyNoEmptyColumns(t, resource, conn)

	if err := conn.Conn().Close(ctx); err != nil {
		t.Fatal(err)
	}

}
