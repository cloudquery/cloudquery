package client

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	uuid "github.com/vgarvardt/pgx-google-uuid/v4"
)
type testRow struct {
	CqId uuid.UUID
	CqParentId uuid.UUID
	CqSourceName string
	CqSyncTime time.Time
	Id uuid.UUID
	Bool bool
	Int int
	Float float64
	Uuid uuid.UUID
	String string
	StringArray []string 
	IntArray []int
	Timestamp time.Time
	Interval time.Duration
	Json map[string]interface{}
	UuidArray []*uuid.UUID
	Inet net.IP
	InetArray []net.IP
	Cidr net.IPNet
	CidrArray []net.IPNet
	Mac string
}

func TestMigrate(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, getTestLogger(t), specs.Destination{
		WriteMode: specs.WriteModeOverwriteDeleteStale,
		Spec: &Spec{
			PgxLogLevel:      LogLevelDebug,
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

	// check migration without column does nothing
	testTable.Columns = testTable.Columns[:len(testTable.Columns)-1]
	if err := c.Migrate(ctx, []*schema.Table{testTable}); err != nil {
		t.Fatalf("failed to migrate tables with missing column: %v", err)
	}

	// migrate primary key
	testTable.Columns[2].CreationOptions.PrimaryKey = true
	if err := c.Migrate(ctx, []*schema.Table{testTable}); err != nil {
		t.Fatalf("failed to migrate tables with missing column: %v", err)
	}
	testTables = schema.Tables{getTestTable()}
	resources := make(chan *schema.DestinationResource, 1)
	resources <- testData
	close(resources)
	if err := c.Write(ctx, testTables, resources); err != nil {
		t.Fatalf("failed to write data: %v", err)
	}
	rows, err := c.conn.Query(ctx, `
	SELECT
		 _cq_id,
		 _cq_parent_id,
		 _cq_source_name,
		 _cq_sync_time,
		 id,
		 bool_column,
		 int_column,
		 float_column,
		 uuid_column,
		 string_column,
		 string_array_column,
		 int_array_column,
		 timestamp_column,
		 interval_column,
		 json_column,
		 --uuid_array_column,
		 inet_column,
		 inet_array_column,
		 cidr_column,
		 cidr_array_column,
		 mac_addr_column
	FROM simple_table`)
	if err != nil {
		t.Fatal(err)
	}
	totalResults := 0
	result := testRow{}
	for rows.Next() {
		if err := rows.Scan(
			&result.CqId,
			&result.CqParentId,
			&result.CqSourceName,
			&result.CqSyncTime,
			&result.Id,
			&result.Bool,
			&result.Int,
			&result.Float,
			&result.Uuid,
			&result.String,
			&result.StringArray,
			&result.IntArray,
			&result.Timestamp,
			&result.Interval,
			&result.Json,
			// &result.UuidArray,
			&result.Inet,
			&result.InetArray,
			&result.Cidr,
			&result.CidrArray,
			&result.Mac,
			); err != nil {
			t.Fatal(err)
		}
		totalResults++
	}
	if totalResults != 1 {
		t.Fatal("expected 1 result, got", totalResults)
	}
	// if diff := cmp.Diff(results, testData.Data); diff != "" {
	// 	t.Fatal(diff)
	// }
}
