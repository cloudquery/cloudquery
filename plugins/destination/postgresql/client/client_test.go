package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/stretchr/testify/require"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5442/postgres?sslmode=disable"
	}
	return testConn
}

func getConnection(ctx context.Context) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(getTestConnection())
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string %w", err)
	}
	pgxConfig.ConnConfig.RuntimeParams["timezone"] = "UTC"
	return pgxpool.NewWithConfig(ctx, pgxConfig)
}

func getIndexesForTable(ctx context.Context, tableName string) (map[string]string, error) {
	conn, err := getConnection(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "SELECT indexname,indexdef FROM pg_indexes WHERE tablename = $1", tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to list indexes %w", err)
	}
	defer rows.Close()
	indexes := make(map[string]string)
	for rows.Next() {
		var indexName, indexDef string
		if err := rows.Scan(&indexName, &indexDef); err != nil {
			return nil, fmt.Errorf("failed to scan index %w", err)
		}
		indexes[indexName] = indexDef
	}
	return indexes, nil
}

var safeMigrations = plugin.SafeMigrations{
	AddColumn:              true,
	AddColumnNotNull:       false,
	RemoveColumn:           true,
	RemoveColumnNotNull:    false,
	RemoveUniqueConstraint: true,
	MovePKToCQOnly:         true,
}

func TestPgPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("postgresql", "development", New)
	s := &spec.Spec{
		ConnectionString: getTestConnection(),
		PgxLogLevel:      spec.LogLevel(tracelog.LogLevelTrace),
	}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	err = p.Init(ctx, b, plugin.NewClientOptions{})
	if err != nil {
		t.Fatal(err)
	}
	testOpts := schema.TestSourceOptions{
		SkipMaps:      true,
		TimePrecision: time.Microsecond, // only us precision supported by time cols
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipDeleteRecord: true,
			SafeMigrations:   safeMigrations,
		},
		plugin.WithTestDataOptions(testOpts),
	)
}

func TestCreateIndexesPluginNewTable(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("postgresql", "development", New)
	s := &spec.Spec{
		ConnectionString:         getTestConnection(),
		PgxLogLevel:              spec.LogLevel(tracelog.LogLevelTrace),
		CreatePerformanceIndexes: true,
	}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	err = p.Init(ctx, b, plugin.NewClientOptions{})
	require.NoError(t, err)

	tableName := fmt.Sprintf("cq_test_create_indexes_%d", time.Now().UnixNano())
	table := &schema.Table{
		Name: tableName,
		Columns: []schema.Column{
			{Name: "_cq_id", Type: types.ExtensionTypes.UUID, PrimaryKey: true, NotNull: true, Unique: true},
			{Name: "_cq_source_name", Type: arrow.BinaryTypes.String},
			{Name: "_cq_sync_time", Type: arrow.FixedWidthTypes.Timestamp_us},
		},
	}

	err = p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: table}})
	require.NoError(t, err)

	indexes, err := getIndexesForTable(ctx, tableName)
	require.NoError(t, err)
	require.Len(t, indexes, 2)
	require.Equal(t, fmt.Sprintf(`CREATE UNIQUE INDEX %[1]s_cqpk ON public.%[1]s USING btree (_cq_id)`, tableName), indexes[fmt.Sprintf("%s_cqpk", tableName)])
	require.Equal(t, fmt.Sprintf(`CREATE INDEX %[1]s_cqpi ON public.%[1]s USING btree (_cq_source_name, _cq_sync_time)`, tableName), indexes[fmt.Sprintf("%s_cqpi", tableName)])
}

func TestCreateIndexesPluginExistingTable(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("postgresql", "development", New)
	s := &spec.Spec{
		ConnectionString:         getTestConnection(),
		PgxLogLevel:              spec.LogLevel(tracelog.LogLevelTrace),
		CreatePerformanceIndexes: false,
	}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	err = p.Init(ctx, b, plugin.NewClientOptions{})
	require.NoError(t, err)

	tableName := fmt.Sprintf("cq_test_create_indexes_%d", time.Now().UnixNano())
	table := &schema.Table{
		Name: tableName,
		Columns: []schema.Column{
			{Name: "_cq_id", Type: types.ExtensionTypes.UUID, PrimaryKey: true, NotNull: true, Unique: true},
			{Name: "_cq_source_name", Type: arrow.BinaryTypes.String},
			{Name: "_cq_sync_time", Type: arrow.FixedWidthTypes.Timestamp_us},
		},
	}
	if err := p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: table}}); err != nil {
		t.Fatal(fmt.Errorf("failed to create table: %w", err))
	}

	indexes, err := getIndexesForTable(ctx, tableName)
	require.NoError(t, err)
	require.Len(t, indexes, 1)
	require.Equal(t, fmt.Sprintf(`CREATE UNIQUE INDEX %[1]s_cqpk ON public.%[1]s USING btree (_cq_id)`, tableName), indexes[fmt.Sprintf("%s_cqpk", tableName)])

	s.CreatePerformanceIndexes = true
	b, err = json.Marshal(s)
	require.NoError(t, err)
	err = p.Init(ctx, b, plugin.NewClientOptions{})
	require.NoError(t, err)

	err = p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: table}})
	require.NoError(t, err)

	indexes, err = getIndexesForTable(ctx, tableName)
	require.NoError(t, err)
	require.Len(t, indexes, 2)
	require.Equal(t, fmt.Sprintf(`CREATE UNIQUE INDEX %[1]s_cqpk ON public.%[1]s USING btree (_cq_id)`, tableName), indexes[fmt.Sprintf("%s_cqpk", tableName)])
	require.Equal(t, fmt.Sprintf(`CREATE INDEX %[1]s_cqpi ON public.%[1]s USING btree (_cq_source_name, _cq_sync_time)`, tableName), indexes[fmt.Sprintf("%s_cqpi", tableName)])
}
