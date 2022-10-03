package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestWriteOverwriteDeleteStale(t *testing.T) {
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
	if err := c.Write(ctx, "simple_table", testData); err != nil {
		t.Fatalf("failed to write data: %v", err)
	}
	if err := c.Write(ctx, "simple_table", testData); err != nil {
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
	if diff := cmp.Diff(results[0], testData); diff != "" {
		t.Fatal(fmt.Errorf("unexpected results: %s", diff))
	}
}

func copyMap(m map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{}, len(m))
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func TestWriteAppend(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, getTestLogger(t), specs.Destination{
		WriteMode: specs.WriteModeAppend,
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
	testDataTwo := copyMap(testData)
	testDataTwo["_cq_id"] = uuid.New().String()
	expectedData := map[string]map[string]interface{}{
		testData["_cq_id"].(string):    testData,
		testDataTwo["_cq_id"].(string): testDataTwo,
	}

	// check migration logic
	if err := c.Drop(ctx, testTables); err != nil {
		t.Fatalf("failed to drop tables: %v", err)
	}
	if err := c.Migrate(ctx, testTables); err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}
	if err := c.Write(ctx, "simple_table", testData); err != nil {
		t.Fatalf("failed to write data: %v", err)
	}
	if err := c.Write(ctx, "simple_table", testDataTwo); err != nil {
		t.Fatalf("failed to write second data: %v", err)
	}

	var results []map[string]interface{}
	totalResults, err := selectTableAsJson(ctx, c.conn, testTable.Name, &results)
	if err != nil {
		t.Fatalf("failed to select data from test table: %v", err)
	}
	if totalResults != 1 {
		t.Fatalf("got %d, expected 1", totalResults)
	}
	if len(results) != 2 {
		t.Fatalf("got %d, expected json_agg to return list with two entries", len(results))
	}
	for _, result := range results {
		cqId, err := uuid.Parse(result["_cq_id"].(string))
		if err != nil {
			t.Fatalf("failed to parse _cq_id: %v", err)
		}
		if diff := cmp.Diff(result, expectedData[cqId.String()]); diff != "" {
			t.Fatal(diff)
		}
	}
}

func selectTableAsJson(ctx context.Context, conn *pgxpool.Pool, table string, results *[]map[string]interface{}) (uint64, error) {
	rows, err := conn.Query(ctx, "SELECT json_agg(simple_table.*) FROM "+table)
	if err != nil {
		return 0, err
	}
	totalResults := uint64(0)
	for rows.Next() {
		if err := rows.Scan(results); err != nil {
			return 0, err
		}
		totalResults++
	}
	return totalResults, nil
}
