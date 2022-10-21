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

func getTestData() *schema.DestinationResource {
	// because data is sent over the wire encoded in json we need to use strings, numbers, objects, arrays, booleans and nulls
	// to test everything correctly
	return &schema.DestinationResource {
		TableName: "simple_table",
		Data: []interface{}{
			uuid.New().String(),
			nil,
			"test_source",
			"2022-09-02T20:57:55.449238",
			uuid.New().String(),
			true,
			float64(3),
			float64(3),
			"9a6011b7-c5ee-4b55-95a6-37ce5e02a5a0",
			"test",
			[]interface{}{"test", "test2"},
			[]interface{}{float64(1), float64(2), float64(3)},
			"2019-01-01T00:00:00",
			"01:02:03",
			map[string]interface{}{"1": float64(1), "test": "test"},
			[]interface{}{"1a6011b7-c5ee-4b55-95a6-37ce5e02a5a0", "9a6011b7-c5ee-4b55-95a6-37ce5e02a5a0"},
			"1.1.1.1",
			[]interface{}{"8.8.8.8"},
			"0.0.0.0/24",
			[]interface{}{"0.0.0.0/24", "0.0.0.0/16"},
			"00:00:00:00:00:ab",
		},
	}
}

func getTestTable() *schema.Table {
	return &schema.Table{
		Name: "simple_table",
		Columns: []schema.Column{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
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
				Name: "interval_column",
				Type: schema.TypeTimeInterval,
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
