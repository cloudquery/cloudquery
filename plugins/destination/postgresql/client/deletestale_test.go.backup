package client

import (
	"context"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func TestDeleteStale(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, getTestLogger(t), specs.Destination{
		WriteMode: specs.WriteModeOverwriteDeleteStale,
		Spec: &Spec{
			ConnectionString: getTestConnection(),
			BatchSize:        1,
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
	resources := make(chan *schema.DestinationResource, 1)
	resources <- testData
	if err := c.Write(ctx, testTables, resources); err != nil {
		t.Fatalf("failed to write data: %v", err)
	}

	var results []map[string]interface{}
	totalResults, err := selectTableAsJson(ctx, c.conn, testTable.Name, &results)
	if err != nil {
		t.Fatalf("failed to select data from test table: %v", err)
	}
	if totalResults != 1 {
		t.Fatalf("got %d, expected 1", totalResults)
	}
	if len(results) != 1 {
		t.Fatalf("got %d, expected json_agg to return list with one entry", len(results))
	}
	// TODO: figure normal pgx<->go time normal conversion
	syncTime := time.Now().UTC().Add(-24 * 30 * 12 * time.Hour)
	// syncTime = syncTime.Add(time.Second * 2)
	if err := c.DeleteStale(ctx, []string{"simple_table"}, "test_source", syncTime); err != nil {
		t.Fatalf("failed to delete stale data: %v", err)
	}
	totalResults, err = selectTableAsJson(ctx, c.conn, testTable.Name, &results)
	if err != nil {
		t.Fatalf("failed to select data from test table: %v", err)
	}
	if totalResults != 1 {
		t.Fatalf("got %d, expected 1", totalResults)
	}
	if len(results) != 1 {
		t.Fatalf("got %d, expected json_agg to return list with 1 entries", len(results))
	}
	syncTime = time.Now().UTC()
	if err := c.DeleteStale(ctx, []string{"simple_table"}, "test_source", syncTime); err != nil {
		t.Fatalf("failed to delete stale data: %v", err)
	}
	totalResults, err = selectTableAsJson(ctx, c.conn, testTable.Name, &results)
	if err != nil {
		t.Fatalf("failed to select data from test table: %v", err)
	}
	if totalResults != 1 {
		t.Fatalf("got %d, expected 0", totalResults)
	}
	if len(results) != 0 {
		t.Fatalf("got %d, expected json_agg to return list with 0 entries", len(results))
	}
}
