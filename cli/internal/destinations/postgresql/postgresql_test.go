package postgresql

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/google/go-cmp/cmp"
	"github.com/rs/zerolog"
)

var createTablesTests = []*schema.Table{
	{
		Name:    "empty_table",
		Columns: nil,
	},
	{
		Name: "simple_table",
		Columns: schema.ColumnList{
			{
				Name: "id",
				Type: schema.TypeUUID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "bool_column",
				Type: schema.TypeBool,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "int_column",
				Type: schema.TypeInt,
			},
			{
				Name: "float_column",
				Type: schema.TypeFloat,
			},
			{
				Name: "uuid_column",
				Type: schema.TypeUUID,
			},
			{
				Name: "string_column",
				Type: schema.TypeString,
			},
			{
				Name: "string_array_column",
				Type: schema.TypeStringArray,
			},
			{
				Name: "int_array_column",
				Type: schema.TypeIntArray,
			},
			{
				Name: "timestamp_column",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "json_column",
				Type: schema.TypeJSON,
			},
			{
				Name: "uuid_array_column",
				Type: schema.TypeUUIDArray,
			},
			{
				Name: "inet_column",
				Type: schema.TypeInet,
			},
			{
				Name: "inet_array_column",
				Type: schema.TypeInetArray,
			},
			{
				Name: "cidr_column",
				Type: schema.TypeCIDR,
			},
			{
				Name: "cidr_array_column",
				Type: schema.TypeCIDRArray,
			},
			{
				Name: "mac_addr_column",
				Type: schema.TypeMacAddr,
			},
		},
	},
}

func TestPostgreSqlCreateTables(t *testing.T) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	ctx := context.Background()
	c := NewClient(l)
	if err := c.Initialize(ctx,
		specs.Destination{
			Spec: &PostgreSqlSpec{
				ConnectionString: "postgres://postgres:pass@localhost:5432/postgres",
				PgxLogLevel:      LogLevelInfo,
			},
		},
	); err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}

	// check migration logic
	if err := c.Drop(ctx, createTablesTests); err != nil {
		t.Fatalf("failed to drop tables: %v", err)
	}
	if err := c.Migrate(ctx, createTablesTests); err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}
	// test that calling migrate twice works
	if err := c.Migrate(ctx, createTablesTests); err != nil {
		t.Fatalf("failed to migrate tables second time: %v", err)
	}
	// check table migration
	createTablesTests[1].Columns = append(createTablesTests[1].Columns, schema.Column{
		Name: "mac_addr_array_column",
		Type: schema.TypeMacAddrArray,
	})
	if err := c.Migrate(ctx, createTablesTests); err != nil {
		t.Fatalf("failed to migrate tables with different column: %v", err)
	}
	// check migration without column does nothing
	createTablesTests[1].Columns = createTablesTests[1].Columns[:len(createTablesTests[1].Columns)-1]
	if err := c.Migrate(ctx, createTablesTests); err != nil {
		t.Fatalf("failed to migrate tables with missing column: %v", err)
	}
	createTablesTests[1].Columns[3].CreationOptions.PrimaryKey = true
	if err := c.Migrate(ctx, createTablesTests); err != nil {
		t.Fatalf("failed to migrate tables with different pk: %v", err)
	}

	data := map[string]interface{}{
		"id":                    "9a6011b7-c5ee-4b55-95a6-37ce5e02a5a0",
		"bool_column":           true,
		"int_column":            float64(3),
		"float_column":          float64(3.3),
		"uuid_column":           "9a6011b7-c5ee-4b55-95a6-37ce5e02a5a0",
		"string_column":         "test",
		"string_array_column":   []interface{}{"test", "test2"},
		"int_array_column":      []interface{}{float64(1), float64(2), float64(3)},
		"timestamp_column":      "2019-01-01T00:00:00",
		"json_column":           map[string]interface{}{"1": float64(1), "test": "test"},
		"uuid_array_column":     []interface{}{"1a6011b7-c5ee-4b55-95a6-37ce5e02a5a0", "9a6011b7-c5ee-4b55-95a6-37ce5e02a5a0"},
		"inet_column":           "1.1.1.1",
		"inet_array_column":     []interface{}{"8.8.8.8/0"},
		"cidr_column":           "0.0.0.0/24",
		"cidr_array_column":     []interface{}{"0.0.0.0/24", "0.0.0.0/16"},
		"mac_addr_column":       "00:00:00:00:00:ab",
		"mac_addr_array_column": nil,
	}
	if err := c.Write(ctx, "simple_table", data); err != nil {
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
	if diff := cmp.Diff(results[0], data); diff != "" {
		t.Fatal(diff)
	}
}
