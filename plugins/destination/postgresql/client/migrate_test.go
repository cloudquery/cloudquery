package client

import (
	"context"
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/google/go-cmp/cmp"
)

func TestMigrate(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, getTestLogger(t), specs.Destination{
		WriteMode: specs.WriteModeOverwriteDeleteStale,
		Spec: &Spec{
			ConnectionString: getTestConnection(),
		},
	})
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}
	c := client.(*Client)
	testTable := getTestTable()
	testTables := []*schema.Table{testTable}
	testData := getTestData()

	// check migration logic
	if err := c.Drop(ctx, testTables); err != nil {
		t.Fatalf("failed to drop tables: %v", err)
	}
	if err := c.Migrate(ctx, testTables); err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}
	isExist, err := c.isTableExistSQL(ctx, testTable.Name)
	if err != nil {
		t.Fatalf("failed to check if table exists: %v", err)
	}
	if !isExist {
		t.Fatalf("failed to migrate table. table %s doesn't exist", testTable.Name)
	}
	// test that calling migrate twice works
	if err := c.Migrate(ctx, testTables); err != nil {
		t.Fatalf("failed to migrate tables second time: %v", err)
	}

	if err := c.Migrate(ctx, []*schema.Table{testTable}); err != nil {
		t.Fatalf("failed to migrate tables with different column: %v", err)
	}
	// check migration without column does nothing
	testTable.Columns = testTable.Columns[:len(testTable.Columns)-1]
	if err := c.Migrate(ctx, []*schema.Table{testTable}); err != nil {
		t.Fatalf("failed to migrate tables with missing column: %v", err)
	}

	if err := c.Write(ctx, "simple_table", testData); err != nil {
		t.Fatalf("failed to write data: %v", err)
	}
	var results []map[string]interface{}
	rows, err := c.conn.Query(ctx, "SELECT json_agg(simple_table.*) FROM simple_table")
	if err != nil {
		t.Fatal(err)
	}
	totalResults := 0
	for rows.Next() {
		if err := rows.Scan(&results); err != nil {
			t.Fatal(err)
		}
		totalResults++
	}
	if totalResults != 1 {
		t.Fatal("expected 1 result, got", totalResults)
	}
	if len(results) != 1 {
		t.Fatal("expected json_agg to return list with one entry, got", len(results))
	}
	if diff := cmp.Diff(results[0], testData); diff != "" {
		t.Fatal(diff)
	}
}
