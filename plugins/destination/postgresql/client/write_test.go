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

type testTypeTestCase struct {
	name string
	typ schema.CQType
}

var testTypeTestCases = []testTypeTestCase{
	{
		name: "string",
		typ: &schema.String{
			String: "test",
			Valid: true,
		},
	},
	{
		name: "uuid",
		typ: schema.NewMustUUID(uuid.New()),
	},
	{
		name: "int",
		typ: &schema.Int64{
			Int64: 1,
			Valid: true,
		},
	},
	{
		name: "bool",
		typ: &schema.Bool{
			Bool: true,
			Valid: true,
		},
	},
	{
		name: "json",
		typ: &schema.Json{
			Json: []byte(`{"test": "test"}`),
			Valid: true,
		},
	},
}

func TestWriteTypes(t *testing.T) {
	for _, tc := range testTypeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			client, err := New(ctx, getTestLogger(t), specs.Destination{
				WriteMode: specs.WriteModeOverwriteDeleteStale,
				Spec: &Spec{
					PgxLogLevel: 		LogLevelTrace,
					ConnectionString: getTestConnection(),
					BatchSize:        1,
				},
			})
			if err != nil {
				t.Fatalf("failed to initialize client: %v", err)
			}
			c := client.(*Client)
			tables := []*schema.Table{
				{
					Name: "test_table",
					Columns: schema.ColumnList{
						{
							Name: "test_column",
							Type: tc.typ.Type(),
						},
					},
				},
			}
			if err := c.Migrate(ctx, tables); err != nil {
				t.Fatalf("failed to migrate tables: %v", err)
			}
			testData := &schema.DestinationResource{
				TableName: tables[0].Name,
				Data: schema.CQTypes{tc.typ},
			}
			resources := make(chan *schema.DestinationResource, 1)
			resources <- testData

			if err := c.Write(ctx, tables, resources); err != nil {
				t.Fatalf("failed to write data: %v", err)
			}

			rows, err := c.conn.Query(ctx, "SELECT test_column FROM test_table")
			if err != nil {
				t.Fatal(err)
			}
			totalResults := uint64(0)
			for rows.Next() {
				
				if err := rows.Scan(results); err != nil {
					return 0, err
				}
				totalResults++
			}

		})
	}
}

func TestWriteOverwriteDeleteStale(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, getTestLogger(t), specs.Destination{
		WriteMode: specs.WriteModeOverwriteDeleteStale,
		Spec: &Spec{
			PgxLogLevel: 		LogLevelTrace,
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
	if diff := cmp.Diff(results[0], testData); diff != "" {
		t.Fatal(fmt.Errorf("unexpected results: %s", diff))
	}
}


func TestWriteAppend(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, getTestLogger(t), specs.Destination{
		WriteMode: specs.WriteModeAppend,
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
	testDataTwo := getTestData()
	testDataTwo.Data[0] = uuid.New().String()
	expectedData := map[string][]interface{}{
		testData.Data[0].(string):    testData.Data,
		testDataTwo.Data[0].(string): testDataTwo.Data,
	}

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
	resources <- testDataTwo
	if err := c.Write(ctx, testTables, resources); err != nil {
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
