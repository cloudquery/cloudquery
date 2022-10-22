package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/cqtypes"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v4"
)

type testTypeTestCase struct {
	name string
	columnType schema.ValueType
	typ schema.CQType
	expected interface{}
}

var testTypeTestCases = []testTypeTestCase{
	{
		name: "string",
		columnType: schema.TypeString,
		typ: &cqtypes.Text{
			String: "test",
			Status: cqtypes.Present,
		},
		expected: "test",
	},
	{
		name: "uuid",
		columnType: schema.TypeUUID,
		typ: &cqtypes.UUID{
			Bytes: [16]byte{1},
			Status: cqtypes.Present,
		},
		expected: [16]byte{1},
	},
	{
		name: "int",
		columnType: schema.TypeInt,
		typ: &cqtypes.Int8{
			Int: 1,
			Status: cqtypes.Present,
		},
		expected: int64(1),
	},
	{
		name: "bool",
		columnType: schema.TypeBool,
		typ: &cqtypes.Bool{
			Bool: true,
			Status: cqtypes.Present,
		},
		expected: true,
	},
	{
		name: "json",
		columnType: schema.TypeJSON,
		typ: &cqtypes.JSON{
			Bytes: []byte(`{"test": "test"}`),
			Status: cqtypes.Present,
		},
		expected: map[string]any{"test": "test"},
	},
}

func TestWriteTypes(t *testing.T) {
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
	for _, tc := range testTypeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			tables := []*schema.Table{
				{
					Name: "test_table",
					Columns: schema.ColumnList{
						schema.CqIDColumn,
						{
							Name: tc.name,
							Type: tc.columnType,
						},
					},
				},
			}
			if err := c.Drop(ctx, tables); err != nil {
				t.Fatal(err)
			}
			if err := c.Migrate(ctx, tables); err != nil {
				t.Fatal(err)
			}
			testData := &schema.DestinationResource{
				TableName: tables[0].Name,
				Data: schema.CQTypes{&cqtypes.UUID{
					Bytes: [16]byte{1},
					Status: cqtypes.Present,
				}, tc.typ},
			}
			resources := make(chan *schema.DestinationResource, 1)
			resources <- testData
			close(resources)
			if err := c.Write(ctx, tables, resources); err != nil {
				t.Fatal(err)
			}

			rows, err := c.conn.Query(ctx, fmt.Sprintf("SELECT %s FROM test_table", pgx.Identifier{tc.name}.Sanitize()))
			if err != nil {
				t.Fatal(err)
			}
			totalResults := uint64(0)
			var res interface{}
			for rows.Next() {
				if err := rows.Scan(&res); err != nil {
					t.Fatal(err)
				}
				totalResults++
			}
			if totalResults != 1 {
				t.Fatalf("expected 1 result got %d", totalResults)
			}
			if diff := cmp.Diff(res, tc.expected); diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

// func TestWriteOverwriteDeleteStale(t *testing.T) {
// 	ctx := context.Background()
// 	client, err := New(ctx, getTestLogger(t), specs.Destination{
// 		WriteMode: specs.WriteModeOverwriteDeleteStale,
// 		Spec: &Spec{
// 			PgxLogLevel: 		LogLevelTrace,
// 			ConnectionString: getTestConnection(),
// 			BatchSize:        1,
// 		},
// 	})
// 	if err != nil {
// 		t.Fatalf("failed to initialize client: %v", err)
// 	}
// 	c := client.(*Client)

// 	testTable := getTestTable()
// 	testTables := []*schema.Table{testTable}
// 	testData := getTestData()

// 	// check migration logic
// 	if err := c.Drop(ctx, testTables); err != nil {
// 		t.Fatalf("failed to drop tables: %v", err)
// 	}
// 	if err := c.Migrate(ctx, testTables); err != nil {
// 		t.Fatalf("failed to migrate tables: %v", err)
// 	}
// 	resources := make(chan *schema.DestinationResource, 1)
// 	resources <- testData
// 	if err := c.Write(ctx, testTables, resources); err != nil {
// 		t.Fatalf("failed to write data: %v", err)
// 	}
// 	resources <- testData
// 	if err := c.Write(ctx, testTables, resources); err != nil {
// 		t.Fatalf("failed to write data: %v", err)
// 	}

// 	var results []map[string]interface{}
// 	totalResults, err := selectTableAsJson(ctx, c.conn, testTable.Name, &results)
// 	if err != nil {
// 		t.Fatalf("failed to select data from test table: %v", err)
// 	}
// 	if totalResults != 1 {
// 		t.Fatalf("got %d, expected 1", totalResults)
// 	}
// 	if len(results) != 1 {
// 		t.Fatalf("got %d, expected json_agg to return list with one entry", len(results))
// 	}
// 	if diff := cmp.Diff(results[0], testData); diff != "" {
// 		t.Fatal(fmt.Errorf("unexpected results: %s", diff))
// 	}
// }


// func TestWriteAppend(t *testing.T) {
// 	ctx := context.Background()
// 	client, err := New(ctx, getTestLogger(t), specs.Destination{
// 		WriteMode: specs.WriteModeAppend,
// 		Spec: &Spec{
// 			ConnectionString: getTestConnection(),
// 			BatchSize:        1,
// 		},
// 	})
// 	if err != nil {
// 		t.Fatalf("failed to initialize client: %v", err)
// 	}
// 	c := client.(*Client)

// 	testTable := getTestTable()
// 	testTables := []*schema.Table{testTable}
// 	testData := getTestData()
// 	testDataTwo := getTestData()
// 	testDataTwo.Data[0] = uuid.New().String()
// 	expectedData := map[string][]interface{}{
// 		testData.Data[0].(string):    testData.Data,
// 		testDataTwo.Data[0].(string): testDataTwo.Data,
// 	}

// 	// check migration logic
// 	if err := c.Drop(ctx, testTables); err != nil {
// 		t.Fatalf("failed to drop tables: %v", err)
// 	}
// 	if err := c.Migrate(ctx, testTables); err != nil {
// 		t.Fatalf("failed to migrate tables: %v", err)
// 	}
// 	resources := make(chan *schema.DestinationResource, 1)
// 	resources <- testData
// 	if err := c.Write(ctx, testTables, resources); err != nil {
// 		t.Fatalf("failed to write data: %v", err)
// 	}
// 	resources <- testDataTwo
// 	if err := c.Write(ctx, testTables, resources); err != nil {
// 		t.Fatalf("failed to write second data: %v", err)
// 	}

// 	var results []map[string]interface{}
// 	totalResults, err := selectTableAsJson(ctx, c.conn, testTable.Name, &results)
// 	if err != nil {
// 		t.Fatalf("failed to select data from test table: %v", err)
// 	}
// 	if totalResults != 1 {
// 		t.Fatalf("got %d, expected 1", totalResults)
// 	}
// 	if len(results) != 2 {
// 		t.Fatalf("got %d, expected json_agg to return list with two entries", len(results))
// 	}
// 	for _, result := range results {
// 		cqId, err := uuid.Parse(result["_cq_id"].(string))
// 		if err != nil {
// 			t.Fatalf("failed to parse _cq_id: %v", err)
// 		}
// 		if diff := cmp.Diff(result, expectedData[cqId.String()]); diff != "" {
// 			t.Fatal(diff)
// 		}
// 	}
// }

// func selectTableAsJson(ctx context.Context, conn *pgxpool.Pool, table string, results *[]map[string]interface{}) (uint64, error) {
// 	rows, err := conn.Query(ctx, "SELECT json_agg(simple_table.*) FROM "+table)
// 	if err != nil {
// 		return 0, err
// 	}
// 	totalResults := uint64(0)
// 	for rows.Next() {
// 		if err := rows.Scan(results); err != nil {
// 			return 0, err
// 		}
// 		totalResults++
// 	}
// 	return totalResults, nil
// }
