package client

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

const (
	ordinaryTable = "aws_ec2_instances"
	hostileTable  = `otherdb.sch."victim"; drop table x`
	hostileQuoted = `"OTHERDB.SCH.""VICTIM""; DROP TABLE X"`
	hostileColumn = `a"b.c; d`
	hostileColSQL = `"A""B.C; D"`
)

func TestQuoteIdentifier(t *testing.T) {
	cases := map[string]string{
		ordinaryTable:        `"AWS_EC2_INSTANCES"`,
		"OTHERDB.SCH.VICTIM": `"OTHERDB.SCH.VICTIM"`,
		hostileTable:         hostileQuoted,
		`"`:                  `""""`,
	}
	for name, want := range cases {
		require.Equal(t, want, sanitizeTable(name))
		require.Equal(t, want, sanitizeColumn(name))
	}
}

func TestMigrationStatementsQuoteIdentifiers(t *testing.T) {
	for _, tc := range []struct {
		table, quotedTable string
		col, quotedCol     string
		quotedUnique       string
		quotedPK           string
	}{
		{ordinaryTable, `"AWS_EC2_INSTANCES"`, "arn", `"ARN"`, `"AWS_EC2_INSTANCES_ARN_KEY"`, `"AWS_EC2_INSTANCES_PK"`},
		{hostileTable, hostileQuoted, hostileColumn, hostileColSQL, `"OTHERDB.SCH.""VICTIM""; DROP TABLE X_A""B.C; D_KEY"`, `"OTHERDB.SCH.""VICTIM""; DROP TABLE X_PK"`},
	} {
		c, rec := newRecordingClient(t)
		ctx := context.Background()
		column := schema.Column{Name: tc.col, Type: arrow.BinaryTypes.String, NotNull: true, Unique: true, PrimaryKey: true}
		table := &schema.Table{Name: tc.table, Columns: schema.ColumnList{column}}
		uniques := constraintMap{}
		uniques[strings.ToUpper(tc.table)] = []constInfo{{tableName: tc.table, colName: tc.col, constName: uniqueConstraintName(tc.table, tc.col), keySeq: 1}}

		require.NoError(t, c.createTableIfNotExist(ctx, table))
		require.NoError(t, c.addColumn(ctx, tc.table, column))
		require.NoError(t, c.alterColumnDropNotNull(ctx, tc.table, column))
		require.NoError(t, c.alterColumnDropUniqueConstraint(ctx, table, uniques, schema.TableColumnChange{ColumnName: tc.col}))
		require.NoError(t, c.alterTableAddPK(ctx, tc.table, []string{tc.col}))
		require.NoError(t, c.alterTableDropPK(ctx, tc.table))
		require.NoError(t, c.dropTable(ctx, tc.table))

		require.Equal(t, []recordedStatement{
			{query: "create table if not exists " + tc.quotedTable + " (" + tc.quotedCol + " text not null constraint " + tc.quotedUnique + " unique, constraint " + tc.quotedPK + " primary key (" + tc.quotedCol + "))"},
			{query: "alter table " + tc.quotedTable + " add column " + tc.quotedCol + " text not null constraint " + tc.quotedUnique + " unique"},
			{query: "alter table " + tc.quotedTable + " alter column " + tc.quotedCol + " drop not null"},
			{query: "alter table " + tc.quotedTable + " drop unique (" + tc.quotedCol + ")"},
			{query: "alter table " + tc.quotedTable + " add constraint " + tc.quotedPK + " primary key(" + tc.quotedCol + ")"},
			{query: "alter table " + tc.quotedTable + " drop primary key"},
			{query: "drop table " + tc.quotedTable},
		}, rec.statements())
	}
}

func TestWriteStatementsQuoteIdentifiers(t *testing.T) {
	for _, tc := range []struct {
		table, quotedTable     string
		col, quotedCol, colKey string
	}{
		{ordinaryTable, `"AWS_EC2_INSTANCES"`, "arn", `"ARN"`, `'arn'`},
		{hostileTable, hostileQuoted, `a"b'c\d`, `"A""B'C\D"`, `'a"b\'c\\d'`},
	} {
		c, rec := newRecordingClient(t)
		ctx := context.Background()
		table := &schema.Table{Name: tc.table, Columns: schema.ColumnList{{Name: tc.col, Type: arrow.BinaryTypes.String, PrimaryKey: true}}}
		f, err := os.CreateTemp(t.TempDir(), stageFilePattern(tc.table))
		require.NoError(t, err)
		require.NoError(t, f.Close())
		file := filepath.Base(f.Name())

		require.NoError(t, c.copyIntoTable(ctx, table, f))
		require.NoError(t, c.mergeIntoTable(ctx, table, f))

		require.Equal(t, []recordedStatement{
			{query: "copy into " + tc.quotedTable + " from '@cq_plugin_stage' files=('" + file + ".gz') on_error = ABORT_STATEMENT file_format = (format_name = cq_plugin_json_format) match_by_column_name = case_insensitive"},
			{query: "MERGE INTO " + tc.quotedTable + " AS dest USING (SELECT $1[" + tc.colKey + "]::text as " + tc.quotedCol + " FROM @cq_plugin_stage/" + file + ".gz t) AS source ON source." + tc.quotedCol + "=dest." + tc.quotedCol + " WHEN MATCHED THEN  UPDATE SET " + tc.quotedCol + "=source." + tc.quotedCol + "  WHEN NOT MATCHED THEN INSERT (" + tc.quotedCol + ") VALUES (source." + tc.quotedCol + ")"},
		}, rec.statements())
	}
}

func TestStageFilePatternKeepsOnlySafeCharacters(t *testing.T) {
	require.Equal(t, "aws_ec2_instances.json.*", stageFilePattern(ordinaryTable))
	require.Equal(t, "a___b_c__d__e_f.json.*", stageFilePattern(`a') b.c; d/"e\f`))
}

func TestDeleteStaleQuotesTable(t *testing.T) {
	syncTime := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	for table, quoted := range map[string]string{ordinaryTable: `"AWS_EC2_INSTANCES"`, hostileTable: hostileQuoted} {
		c, rec := newRecordingClient(t)

		require.NoError(t, c.DeleteStale(context.Background(), message.WriteDeleteStales{{TableName: table, SourceName: "src", SyncTime: syncTime}}))

		require.Equal(t, []recordedStatement{{
			query: "delete from " + quoted + ` where "_CQ_SOURCE_NAME" = ? and "_CQ_SYNC_TIME" < CONVERT_TIMEZONE('UTC','UTC', TO_TIMESTAMP(?))`,
			args:  []driver.Value{"src", syncTime},
		}}, rec.statements())
	}
}

func TestReadQuotesIdentifiers(t *testing.T) {
	for _, tc := range []struct{ table, col, want string }{
		{ordinaryTable, "arn", `SELECT "ARN" FROM "AWS_EC2_INSTANCES"`},
		{hostileTable, hostileColumn, "SELECT " + hostileColSQL + " FROM " + hostileQuoted},
	} {
		c, rec := newRecordingClient(t)
		table := &schema.Table{Name: tc.table, Columns: schema.ColumnList{{Name: tc.col, Type: arrow.BinaryTypes.String}}}

		require.NoError(t, c.Read(context.Background(), table, make(chan arrow.RecordBatch)))

		require.Equal(t, []recordedStatement{{query: tc.want}}, rec.statements())
	}
}

type recordedStatement struct {
	query string
	args  []driver.Value
}

type statementRecorder struct {
	mu       sync.Mutex
	recorded []recordedStatement
}

func newRecordingClient(t *testing.T) (*Client, *statementRecorder) {
	t.Helper()

	rec := &statementRecorder{}
	db := sql.OpenDB(rec)
	t.Cleanup(func() { require.NoError(t, db.Close()) })
	return &Client{db: db}, rec
}

func (r *statementRecorder) statements() []recordedStatement {
	r.mu.Lock()
	defer r.mu.Unlock()

	return append([]recordedStatement(nil), r.recorded...)
}

func (r *statementRecorder) record(query string, named []driver.NamedValue) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.recorded = append(r.recorded, recordedStatement{query: query, args: argValues(named)})
}

func argValues(named []driver.NamedValue) []driver.Value {
	if len(named) == 0 {
		return nil
	}
	values := make([]driver.Value, len(named))
	for i, arg := range named {
		values[i] = arg.Value
	}
	return values
}

func (r *statementRecorder) Connect(context.Context) (driver.Conn, error) { return r, nil }
func (*statementRecorder) Driver() driver.Driver                          { return nil }
func (*statementRecorder) Prepare(string) (driver.Stmt, error)            { return nil, driver.ErrSkip }
func (*statementRecorder) Close() error                                   { return nil }
func (*statementRecorder) Begin() (driver.Tx, error)                      { return nil, driver.ErrSkip }

func (r *statementRecorder) ExecContext(_ context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	r.record(query, args)
	return driver.RowsAffected(0), nil
}

func (r *statementRecorder) QueryContext(_ context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	r.record(query, args)
	return &fakeRows{}, nil
}
