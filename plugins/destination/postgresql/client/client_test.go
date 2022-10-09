package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func getTestData() map[string]interface{} {
	// because data is sent over the wire encoded in json we need to use strings, numbers, objects, arrays, booleans and nulls
	// to test everything correctly
	return map[string]interface{}{
		"_cq_id":              uuid.New().String(),
		"_cq_parent_id":       nil,
		"_cq_source_name":     "test_source",
		"_cq_sync_time":       "2022-09-02T20:57:55.449238",
		"id":                  uuid.New().String(),
		"bool_column":         true,
		"int_column":          float64(3),
		"float_column":        float64(3),
		"uuid_column":         "9a6011b7-c5ee-4b55-95a6-37ce5e02a5a0",
		"string_column":       "test",
		"string_array_column": []interface{}{"test", "test2"},
		"int_array_column":    []interface{}{float64(1), float64(2), float64(3)},
		"timestamp_column":    "2019-01-01T00:00:00",
		"json_column":         map[string]interface{}{"1": float64(1), "test": "test"},
		"uuid_array_column":   []interface{}{"1a6011b7-c5ee-4b55-95a6-37ce5e02a5a0", "9a6011b7-c5ee-4b55-95a6-37ce5e02a5a0"},
		"inet_column":         "1.1.1.1",
		"inet_array_column":   []interface{}{"8.8.8.8/0"},
		"cidr_column":         "0.0.0.0/24",
		"cidr_array_column":   []interface{}{"0.0.0.0/24", "0.0.0.0/16"},
		"mac_addr_column":     "00:00:00:00:00:ab",
	}
}

func getTestTable() *schema.Table {
	return &schema.Table{
		Name: "simple_table",
		Columns: []schema.Column{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSyncTime,
			schema.CqSourceName,
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
	}
}

func getTestLogger(t *testing.T) zerolog.Logger {
	t.Helper()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	return zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	return testConn
}

func TestInitialize(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, getTestLogger(t), specs.Destination{
		Spec: Spec{
			ConnectionString: getTestConnection(),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if client == nil {
		t.Fatal("client is nil")
	}
	if err := client.Close(ctx); err != nil {
		t.Fatal(err)
	}
	err = client.Close(ctx)
	if err == nil {
		t.Fatal("expected error when closing a closed client second time")
	}

	if err.Error() != "client already closed or not initialized" {
		t.Fatal("expected error when closing a closed client second time")
	}
}
