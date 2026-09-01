package client

import (
	"context"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// Avoids having to build minCopyRows worth of rows to reach the COPY path.
func forceCopyFrom(t *testing.T) {
	t.Helper()
	previous := minCopyRows
	minCopyRows = 0
	t.Cleanup(func() { minCopyRows = previous })
}

// The only coverage of binary-format encoding per Arrow type: insertBatchExec
// negotiates formats over the extended query protocol, COPY cannot.
func TestPgPluginCopyFrom(t *testing.T) {
	forceCopyFrom(t)

	ctx := context.Background()
	p := plugin.NewPlugin("postgresql", "development", New)
	b, err := json.Marshal(&spec.Spec{
		ConnectionString: getTestConnection(),
		PgxLogLevel:      spec.LogLevel(tracelog.LogLevelError),
		UseCopyFrom:      true,
	})
	require.NoError(t, err)
	require.NoError(t, p.Init(ctx, b, plugin.NewClientOptions{}))

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{SafeMigrations: safeMigrations},
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			SkipMaps:      true,
			TimePrecision: time.Microsecond,
		}),
	)
}

func copyTestTable(t *testing.T, withPK bool) *schema.Table {
	t.Helper()
	return &schema.Table{
		Name: fmt.Sprintf("cq_copy_%d", time.Now().UnixNano()),
		Columns: []schema.Column{
			{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: withPK, NotNull: withPK},
			{Name: "name", Type: arrow.BinaryTypes.String},
			schema.CqSyncTimeColumn,
		},
	}
}

func newCopyPlugin(t *testing.T) *plugin.Plugin {
	t.Helper()
	p := plugin.NewPlugin("postgresql", "development", New, plugin.WithJSONSchema(spec.JSONSchema))
	p.SetLogger(zerolog.New(zerolog.NewTestWriter(t)).Level(zerolog.WarnLevel))
	b, err := json.Marshal(&spec.Spec{ConnectionString: getTestConnection(), UseCopyFrom: true})
	require.NoError(t, err)
	require.NoError(t, p.Init(context.Background(), b, plugin.NewClientOptions{}))
	return p
}

func copyRecord(table *schema.Table, id, name string) arrow.RecordBatch {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	defer bldr.Release()
	bldr.Field(0).(*array.StringBuilder).Append(id)
	bldr.Field(1).(*array.StringBuilder).Append(name)
	bldr.Field(2).(*array.TimestampBuilder).Append(arrow.Timestamp(time.Now().UnixMicro()))
	return bldr.NewRecordBatch()
}

func readAll(ctx context.Context, t *testing.T, p *plugin.Plugin, table *schema.Table) map[string]string {
	t.Helper()
	ch := make(chan arrow.RecordBatch)
	var readErr error
	go func() {
		defer close(ch)
		readErr = p.Read(ctx, table, ch)
	}()
	out := map[string]string{}
	for record := range ch {
		ids := record.Column(0).(*array.String)
		names := record.Column(1).(*array.String)
		for i := range int(record.NumRows()) {
			out[ids.Value(i)] = names.Value(i)
		}
	}
	require.NoError(t, readErr)
	return out
}

func TestCopyFrom_UpsertsWithoutDuplicating(t *testing.T) {
	const rows = 500
	ctx := context.Background()
	table := copyTestTable(t, true)
	p := newCopyPlugin(t)
	t.Cleanup(func() { _ = p.Close(ctx) })

	ids := make([]string, rows)
	msgs := make([]message.WriteMessage, 0, rows+1)
	msgs = append(msgs, &message.WriteMigrateTable{Table: table})
	for i := range rows {
		ids[i] = uuid.New().String()
		msgs = append(msgs, &message.WriteInsert{Record: copyRecord(table, ids[i], "first")})
	}
	require.NoError(t, p.WriteAll(ctx, msgs))

	got := readAll(ctx, t, p, table)
	require.Len(t, got, rows)
	require.Equal(t, "first", got[ids[0]])

	rewrite := make([]message.WriteMessage, 0, rows)
	for i := range rows {
		rewrite = append(rewrite, &message.WriteInsert{Record: copyRecord(table, ids[i], "second")})
	}
	require.NoError(t, p.WriteAll(ctx, rewrite))

	got = readAll(ctx, t, p, table)
	require.Len(t, got, rows, "rewriting the same primary keys must upsert, not duplicate")
	require.Equal(t, "second", got[ids[0]])
}

// Regression test for the distinct on in mergeFromStaging: without it Postgres
// rejects the whole merge with "cannot affect row a second time".
func TestCopyFrom_DuplicatePrimaryKeysInOneBatch(t *testing.T) {
	const rows = 300
	ctx := context.Background()
	table := copyTestTable(t, true)
	p := newCopyPlugin(t)
	t.Cleanup(func() { _ = p.Close(ctx) })

	id := uuid.New().String()
	msgs := make([]message.WriteMessage, 0, rows+1)
	msgs = append(msgs, &message.WriteMigrateTable{Table: table})
	for i := range rows {
		msgs = append(msgs, &message.WriteInsert{Record: copyRecord(table, id, fmt.Sprintf("v%03d", i))})
	}
	require.NoError(t, p.WriteAll(ctx, msgs))

	got := readAll(ctx, t, p, table)
	require.Len(t, got, 1)
	require.Equal(t, fmt.Sprintf("v%03d", rows-1), got[id], "the last row in the batch must win")
}

func TestCopyFrom_NoPrimaryKeyCopiesDirectly(t *testing.T) {
	const rows = 400
	ctx := context.Background()
	table := copyTestTable(t, false)
	p := newCopyPlugin(t)
	t.Cleanup(func() { _ = p.Close(ctx) })

	msgs := make([]message.WriteMessage, 0, rows+1)
	msgs = append(msgs, &message.WriteMigrateTable{Table: table})
	for range rows {
		msgs = append(msgs, &message.WriteInsert{Record: copyRecord(table, uuid.New().String(), "x")})
	}
	require.NoError(t, p.WriteAll(ctx, msgs))
	require.Len(t, readAll(ctx, t, p, table), rows)
}

// The mixed batch writer hands InsertBatch several tables at once; COPY targets
// one at a time.
func TestCopyFrom_MixedTablesInOneBatch(t *testing.T) {
	const rowsPerTable = 150
	ctx := context.Background()
	first, second := copyTestTable(t, true), copyTestTable(t, true)
	second.Name += "_b"
	p := newCopyPlugin(t)
	t.Cleanup(func() { _ = p.Close(ctx) })

	msgs := make([]message.WriteMessage, 0, 2*rowsPerTable+2)
	msgs = append(msgs,
		&message.WriteMigrateTable{Table: first},
		&message.WriteMigrateTable{Table: second},
	)
	// Interleaved so the grouping has to reorder them.
	for range rowsPerTable {
		msgs = append(msgs,
			&message.WriteInsert{Record: copyRecord(first, uuid.New().String(), "a")},
			&message.WriteInsert{Record: copyRecord(second, uuid.New().String(), "b")},
		)
	}
	require.NoError(t, p.WriteAll(ctx, msgs))

	require.Len(t, readAll(ctx, t, p, first), rowsPerTable)
	require.Len(t, readAll(ctx, t, p, second), rowsPerTable)
}

func TestCopyFrom_SmallBatchFallsBackToInsert(t *testing.T) {
	ctx := context.Background()
	table := copyTestTable(t, true)
	p := newCopyPlugin(t)
	t.Cleanup(func() { _ = p.Close(ctx) })

	msgs := make([]message.WriteMessage, 0, 2)
	msgs = append(msgs, &message.WriteMigrateTable{Table: table})
	id := uuid.New().String()
	msgs = append(msgs, &message.WriteInsert{Record: copyRecord(table, id, "only")})
	require.NoError(t, p.WriteAll(ctx, msgs))

	got := readAll(ctx, t, p, table)
	require.Len(t, got, 1)
	require.Equal(t, "only", got[id])
}

func TestUseCopyFrom(t *testing.T) {
	pgVector := &spec.PgVectorConfig{Tables: []spec.PgVectorTableConfig{{SourceTableName: "s", TargetTableName: "t"}}}

	for _, tt := range []struct {
		name   string
		spec   spec.Spec
		pgType pgType
		want   bool
	}{
		{name: "off unless asked for", spec: spec.Spec{}, pgType: pgTypePostgreSQL, want: false},
		{name: "opt in", spec: spec.Spec{UseCopyFrom: true}, pgType: pgTypePostgreSQL, want: true},
		{name: "cockroachdb", spec: spec.Spec{UseCopyFrom: true}, pgType: pgTypeCockroachDB, want: false},
		{name: "cratedb", spec: spec.Spec{UseCopyFrom: true}, pgType: pgTypeCrateDB, want: false},
		{name: "pgvector", spec: spec.Spec{UseCopyFrom: true, PgVectorConfig: pgVector}, pgType: pgTypePostgreSQL, want: false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{spec: &tt.spec, pgType: tt.pgType}
			require.Equal(t, tt.want, c.useCopyFrom())
		})
	}
}

func TestCopyGroupEligible(t *testing.T) {
	withPK := &schema.Table{
		Name:             "t",
		PkConstraintName: "t_cqpk",
		Columns:          []schema.Column{{Name: "id", PrimaryKey: true}},
	}
	noPK := &schema.Table{Name: "t", Columns: []schema.Column{{Name: "id"}}}
	noConstraint := &schema.Table{
		Name:    "t",
		Columns: []schema.Column{{Name: "id", PrimaryKey: true}},
	}

	eligible := func(table *schema.Table, pkColumns []string, rows int64) bool {
		return (&copyGroup{table: table, pkColumns: pkColumns, rows: rows}).eligible()
	}

	require.True(t, eligible(withPK, []string{"id"}, minCopyRows))
	require.False(t, eligible(withPK, []string{"id"}, minCopyRows-1), "small batches use the row-by-row path")
	require.True(t, eligible(noPK, nil, minCopyRows))
	require.False(t, eligible(noConstraint, []string{"id"}, minCopyRows),
		"a primary key with no known constraint name cannot be merged from staging")
	require.False(t, eligible(withPK, nil, minCopyRows),
		"a constraint whose columns are unknown cannot be deduped on")
}

func TestMergeFromStaging(t *testing.T) {
	table := &schema.Table{
		Name:             "aws_instances",
		PkConstraintName: "aws_instances_cqpk",
		Columns: []schema.Column{
			{Name: "_cq_id", PrimaryKey: true},
			{Name: "name"},
		},
	}
	require.Equal(t,
		`insert into "aws_instances" ("_cq_id","name") `+
			`select distinct on ("_cq_id") "_cq_id","name" from "cq_staging_1" `+
			`order by "_cq_id", ctid desc `+
			`on conflict on constraint "aws_instances_cqpk" `+
			`do update set "_cq_id"=excluded."_cq_id","name"=excluded."name"`,
		mergeFromStaging(table, "cq_staging_1", []string{"_cq_id"}))

	// The dedupe key follows the constraint, not table.PrimaryKeys().
	require.Contains(t,
		mergeFromStaging(table, "cq_staging_1", []string{"name"}),
		`select distinct on ("name")`)
}

// Deduping on the source schema's primary keys once they have drifted from the
// database's aborts the merge with "cannot affect row a second time".
func TestCopyFrom_PrimaryKeyDriftFromDatabase(t *testing.T) {
	forceCopyFrom(t)

	ctx := context.Background()
	table := copyTestTable(t, true)
	p := newCopyPlugin(t)
	t.Cleanup(func() { _ = p.Close(ctx) })

	require.NoError(t, p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: table}}))

	// No migrate message: the source moves its primary key while the database
	// keeps its constraint on id.
	drifted := &schema.Table{Name: table.Name, Columns: slices.Clone(table.Columns)}
	drifted.Columns[0].PrimaryKey = false
	drifted.Columns[1].PrimaryKey = true

	// The same id under two source primary keys: only deduping on id leaves the
	// merge one row per constraint key.
	require.NoError(t, p.WriteAll(ctx, []message.WriteMessage{
		&message.WriteInsert{Record: copyRecord(drifted, "same-id", "first")},
		&message.WriteInsert{Record: copyRecord(drifted, "same-id", "second")},
	}))

	require.Equal(t, map[string]string{"same-id": "second"}, readAll(ctx, t, p, table),
		"the last row written must win, as it does on the insert path")
}
