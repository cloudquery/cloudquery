package client

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/resources/plugin"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/typeconv/ch/types"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func getTestConnection() string {
	if testConn := os.Getenv("CQ_DEST_CH_TEST_CONN"); len(testConn) > 0 {
		return testConn
	}

	return (&url.URL{
		User: url.UserPassword("cq", "test"),
		Host: "localhost:9000",
		Path: "cloudquery", // database
	}).String()
}

func TestPlugin(t *testing.T) {
	p := initPlugin(t)

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipUpsert: true,
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:    true,
				RemoveColumn: true,
				// MovePKToCQOnly- is only a change to the underlying PKs, and because clickhouse only supports append only mode this is not a factor
				MovePKToCQOnly: true,
			},
			SkipSpecificMigrations: plugin.Migrations{
				RemoveUniqueConstraint: true,
			},
		},
		plugin.WithTestSourceAllowNull(types.CanBeNullable),
	)
}

func initPlugin(t *testing.T) *plugin.Plugin {
	ctx := context.Background()
	p := plugin.NewPlugin("clickhouse",
		internalPlugin.Version,
		New,
		plugin.WithJSONSchema(spec.JSONSchema),
	)
	s := &spec.Spec{
		ConnectionString: getTestConnection(),
	}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	require.NoError(t, p.Init(ctx, b, plugin.NewClientOptions{}))
	return p
}

func TestMigrateCQClientIDColumnWhenSortKeyIsAlreadySet(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("clickhouse",
		internalPlugin.Version,
		New,
		plugin.WithJSONSchema(spec.JSONSchema),
	)
	s := &spec.Spec{ConnectionString: getTestConnection(), OrderBy: []spec.OrderByStrategy{
		{
			OrderBy: []string{"`_cq_source_name`", "`_cq_sync_group_id`", "`_cq_id`"},
		},
	}}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	err = p.Init(ctx, b, plugin.NewClientOptions{})
	require.NoError(t, err)

	tableName := fmt.Sprintf("cq_test_migrate_cq_client_id_column_%d", time.Now().UnixNano())
	table := &schema.Table{
		Name: tableName,
		Columns: []schema.Column{
			schema.CqIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			{
				Name:       "_cq_sync_group_id",
				Type:       arrow.BinaryTypes.String,
				NotNull:    true,
				PrimaryKey: true,
			},
		},
	}
	if err := p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: table}}); err != nil {
		t.Fatal(fmt.Errorf("failed to create table: %w", err))
	}

	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	bldr.Field(0).(*sdkTypes.UUIDBuilder).Append(uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"))
	bldr.Field(1).(*array.StringBuilder).Append("foo")
	bldr.Field(2).(*array.TimestampBuilder).Append(arrow.Timestamp(time.Now().UnixMicro()))
	bldr.Field(3).(*array.StringBuilder).Append("cq-sync-group-id")
	record := bldr.NewRecord()

	if err := p.WriteAll(ctx, []message.WriteMessage{&message.WriteInsert{
		Record: record,
	}}); err != nil {
		t.Fatal(fmt.Errorf("failed to insert record: %w", err))
	}

	tableWithCQClientIDColumn := &schema.Table{
		Name:    tableName,
		Columns: append(table.Columns, schema.CqClientIDColumn),
	}

	if err := p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: tableWithCQClientIDColumn}}); err != nil {
		t.Fatal(fmt.Errorf("failed to migrate table: %w", err))
	}

	bldr = array.NewRecordBuilder(memory.DefaultAllocator, tableWithCQClientIDColumn.ToArrowSchema())
	bldr.Field(0).(*sdkTypes.UUIDBuilder).Append(uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"))
	bldr.Field(1).(*array.StringBuilder).Append("foo")
	bldr.Field(2).(*array.TimestampBuilder).Append(arrow.Timestamp(time.Now().UnixMicro()))
	bldr.Field(3).(*array.StringBuilder).Append("cq-sync-group-id")
	bldr.Field(4).(*array.StringBuilder).Append("cq-client-id")
	record = bldr.NewRecord()

	if err := p.WriteAll(ctx, []message.WriteMessage{&message.WriteInsert{
		Record: record,
	}}); err != nil {
		t.Fatal(fmt.Errorf("failed to insert record: %w", err))
	}
}

func TestMigrateNewArrayAndMapColumns(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("clickhouse",
		internalPlugin.Version,
		New,
		plugin.WithJSONSchema(spec.JSONSchema),
	)
	s := &spec.Spec{ConnectionString: getTestConnection()}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	err = p.Init(ctx, b, plugin.NewClientOptions{})
	require.NoError(t, err)

	tableName := fmt.Sprintf("cq_test_migrate_new_array_and_map_columns_%d", time.Now().UnixNano())
	table := &schema.Table{
		Name: tableName,
		Columns: []schema.Column{
			schema.CqIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.CqClientIDColumn,
			{
				Name:       "_cq_sync_group_id",
				Type:       arrow.BinaryTypes.String,
				NotNull:    true,
				PrimaryKey: true,
			},
		},
	}
	if err := p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: table}}); err != nil {
		t.Fatal(fmt.Errorf("failed to create table: %w", err))
	}

	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	bldr.Field(0).(*sdkTypes.UUIDBuilder).Append(uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"))
	bldr.Field(1).(*array.StringBuilder).Append("foo")
	bldr.Field(2).(*array.TimestampBuilder).Append(arrow.Timestamp(time.Now().UnixMicro()))
	bldr.Field(4).(*array.StringBuilder).Append("cq-client-id")
	bldr.Field(3).(*array.StringBuilder).Append("cq-sync-group-id")
	record := bldr.NewRecord()

	if err := p.WriteAll(ctx, []message.WriteMessage{&message.WriteInsert{
		Record: record,
	}}); err != nil {
		t.Fatal(fmt.Errorf("failed to insert record: %w", err))
	}

	newColumns := schema.ColumnList{
		{
			Name: "array_column",
			Type: arrow.ListOf(arrow.BinaryTypes.String),
		},
		{
			Name: "map_column",
			Type: arrow.MapOf(arrow.BinaryTypes.String, arrow.BinaryTypes.String),
		},
	}
	tableWithCQClientIDColumn := &schema.Table{
		Name:    tableName,
		Columns: append(table.Columns, newColumns...),
	}

	if err := p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: tableWithCQClientIDColumn}}); err != nil {
		t.Fatal(fmt.Errorf("failed to migrate table: %w", err))
	}

	bldr = array.NewRecordBuilder(memory.DefaultAllocator, tableWithCQClientIDColumn.ToArrowSchema())
	bldr.Field(0).(*sdkTypes.UUIDBuilder).Append(uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"))
	bldr.Field(1).(*array.StringBuilder).Append("foo")
	bldr.Field(2).(*array.TimestampBuilder).Append(arrow.Timestamp(time.Now().UnixMicro()))
	bldr.Field(4).(*array.StringBuilder).Append("cq-client-id")
	bldr.Field(3).(*array.StringBuilder).Append("cq-sync-group-id")
	bldr.Field(5).(*array.ListBuilder).Append(true)
	bldr.Field(5).(*array.ListBuilder).ValueBuilder().(*array.StringBuilder).Append("foo")
	bldr.Field(6).(*array.MapBuilder).Append(true)
	record = bldr.NewRecord()

	if err := p.WriteAll(ctx, []message.WriteMessage{&message.WriteInsert{
		Record: record,
	}}); err != nil {
		t.Fatal(fmt.Errorf("failed to insert record: %w", err))
	}
}

func TestConcurrentSyncsSameTable(t *testing.T) {
	const syncConcurrency = 200
	ctx := context.Background()
	group, _ := errgroup.WithContext(ctx)
	randomUUIDString := uuid.New().String()
	tableName := "k8s_core_namespaces_" + randomUUIDString
	table := &schema.Table{
		Name: tableName,
		Columns: []schema.Column{
			schema.CqIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
		},
	}
	// Create the table
	p := plugin.NewPlugin("clickhouse",
		internalPlugin.Version,
		New,
		plugin.WithJSONSchema(spec.JSONSchema),
	)
	s := &spec.Spec{ConnectionString: getTestConnection()}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	err = p.Init(ctx, b, plugin.NewClientOptions{})
	require.NoError(t, err)
	if err := p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: table}}); err != nil {
		t.Fatal(fmt.Errorf("failed to create table: %w", err))
	}

	for i := range syncConcurrency {
		group.Go(func() error {
			// Simulate a sync job against the same table
			syncContext := context.Background()
			p := plugin.NewPlugin("clickhouse",
				internalPlugin.Version,
				New,
				plugin.WithJSONSchema(spec.JSONSchema),
			)
			s := &spec.Spec{ConnectionString: getTestConnection()}
			b, err := json.Marshal(s)
			require.NoError(t, err)
			err = p.Init(syncContext, b, plugin.NewClientOptions{})
			require.NoError(t, err)
			if err := p.WriteAll(syncContext, []message.WriteMessage{&message.WriteMigrateTable{Table: table}}); err != nil {
				t.Fatal(fmt.Errorf("failed to create table: %w", err))
			}

			jobIndexAsString := strconv.Itoa(i)
			randomUUIDStringWithLastCharacterReplaced := randomUUIDString[:len(randomUUIDString)-len(jobIndexAsString)] + jobIndexAsString
			bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
			bldr.Field(0).(*sdkTypes.UUIDBuilder).Append(uuid.MustParse(randomUUIDStringWithLastCharacterReplaced))
			bldr.Field(1).(*array.StringBuilder).Append("source")
			bldr.Field(2).(*array.TimestampBuilder).Append(arrow.Timestamp(time.Now().UnixMicro()))
			record := bldr.NewRecord()

			if err := p.WriteAll(syncContext, []message.WriteMessage{&message.WriteInsert{
				Record: record,
			}}); err != nil {
				t.Fatal(fmt.Errorf("failed to insert record: %w", err))
			}
			return nil
		})
	}

	require.NoError(t, group.Wait())

	ch := make(chan arrow.Record)
	go func() {
		defer close(ch)
		err = p.Read(ctx, table, ch)
	}()

	numRows := 0
	for record := range ch {
		numRows += int(record.NumRows())
	}

	require.Equal(t, syncConcurrency, numRows)
	require.NoError(t, err)
}
