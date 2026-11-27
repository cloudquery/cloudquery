package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"path"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/duckdb/duckdb-go/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

const (
	hostileTableName  = `hostile" (id bigint); create table injected_by_table (id bigint); --`
	hostileColumnName = `value" bigint); create table injected_by_column (id bigint); --`
)

func TestSanitizeID(t *testing.T) {
	require.Equal(t, `"plain"`, sanitizeID("plain"))
	require.Equal(t, `"with space; semicolon"`, sanitizeID("with space; semicolon"))
	require.Equal(t, `"with""quote"`, sanitizeID(`with"quote`))
	require.Equal(t, `"""; drop table t; --"`, sanitizeID(`"; drop table t; --`))
	require.Equal(t, []string{`"a"`, `"b""c"`}, sanitized([]string{"a", `b"c`}))
}

func TestSanitizeLiteral(t *testing.T) {
	require.Equal(t, `'plain'`, sanitizeLiteral("plain"))
	require.Equal(t, `'it''s'`, sanitizeLiteral("it's"))
}

func TestHostileIdentifiers(t *testing.T) {
	listType := arrow.ListOf(arrow.BinaryTypes.String)
	cases := []struct {
		name     string
		columns  []schema.Column
		wantRows int64
	}{
		{
			name: "no primary key",
			columns: []schema.Column{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: hostileColumnName, Type: arrow.BinaryTypes.String},
			},
			wantRows: 2,
		},
		{
			name: "primary key",
			columns: []schema.Column{
				{Name: hostileColumnName, Type: arrow.BinaryTypes.String, PrimaryKey: true},
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			},
			wantRows: 1,
		},
		{
			name: "primary key with list column",
			columns: []schema.Column{
				{Name: hostileColumnName, Type: arrow.BinaryTypes.String, PrimaryKey: true},
				{Name: "tags", Type: listType},
			},
			wantRows: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			tempDB := path.Join(t.TempDir(), "test_hostile_identifiers.duckdb") + "?threads=1"
			p := newTestPlugin(ctx, t, tempDB)
			table := &schema.Table{Name: hostileTableName, Columns: tc.columns}

			writeHostileRow(ctx, t, p, table)
			writeHostileRow(ctx, t, p, table)

			require.Equal(t, tc.wantRows, readRowCount(ctx, t, p, table))

			connector, err := duckdb.NewConnector(tempDB, nil)
			require.NoError(t, err)
			defer connector.Close()
			db := sql.OpenDB(connector)
			defer db.Close()

			require.Equal(t, []string{hostileTableName}, queryStrings(ctx, t, db,
				"SELECT table_name FROM information_schema.tables ORDER BY table_name"))
			require.Equal(t, schema.ColumnList(tc.columns).Names(), queryStrings(ctx, t, db,
				"SELECT column_name FROM information_schema.columns WHERE table_name = ? ORDER BY ordinal_position", hostileTableName))
		})
	}
}

func newTestPlugin(ctx context.Context, t *testing.T, connectionString string) *plugin.Plugin {
	t.Helper()
	p := plugin.NewPlugin("duckdb", "development", New)
	specBytes, err := json.Marshal(Spec{ConnectionString: connectionString, Debug: true})
	require.NoError(t, err)
	p.SetLogger(zerolog.New(zerolog.NewTestWriter(t)).Level(zerolog.DebugLevel))
	require.NoError(t, p.Init(ctx, specBytes, plugin.NewClientOptions{}))
	t.Cleanup(func() {
		if err := p.Close(ctx); err != nil {
			t.Logf("failed to close plugin: %v", err)
		}
	})
	return p
}

func writeHostileRow(ctx context.Context, t *testing.T, p *plugin.Plugin, table *schema.Table) {
	t.Helper()
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	for i, col := range table.Columns {
		switch b := bldr.Field(i).(type) {
		case *array.Int64Builder:
			b.Append(1)
		case *array.StringBuilder:
			b.Append("foo")
		case *array.ListBuilder:
			b.Append(true)
			b.ValueBuilder().(*array.StringBuilder).Append("bar")
		default:
			t.Fatalf("unsupported column %s", col.Name)
		}
	}

	msgs := make(chan message.WriteMessage, 2)
	msgs <- &message.WriteMigrateTable{Table: table}
	msgs <- &message.WriteInsert{Record: bldr.NewRecordBatch()}
	close(msgs)
	require.NoError(t, p.Write(ctx, msgs))
}

func readRowCount(ctx context.Context, t *testing.T, p *plugin.Plugin, table *schema.Table) int64 {
	t.Helper()
	records := make(chan arrow.RecordBatch, 10)
	require.NoError(t, p.Read(ctx, table, records))
	close(records)
	var count int64
	for record := range records {
		count += record.NumRows()
	}
	return count
}

func queryStrings(ctx context.Context, t *testing.T, db *sql.DB, query string, args ...any) []string {
	t.Helper()
	rows, err := db.QueryContext(ctx, query, args...)
	require.NoError(t, err)
	defer rows.Close()
	var values []string
	for rows.Next() {
		var value string
		require.NoError(t, rows.Scan(&value))
		values = append(values, value)
	}
	require.NoError(t, rows.Err())
	return values
}
