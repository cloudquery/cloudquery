package client

import (
	"context"
	"errors"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// recordingExecPool captures every SQL string sent through Exec so tests can
// assert on identifier quoting without a live database.
type recordingExecPool struct {
	MockDBPool
	execSQL []string
}

func (p *recordingExecPool) Exec(_ context.Context, sql string, _ ...any) (pgconn.CommandTag, error) {
	p.execSQL = append(p.execSQL, sql)
	return pgconn.CommandTag{}, nil
}

func TestDropTableSanitizesTableName(t *testing.T) {
	ctx := context.Background()
	pool := &recordingExecPool{}
	c := &Client{conn: pool, logger: zerolog.Nop()}

	malicious := `weird"; drop table users; --`
	require.NoError(t, c.dropTable(ctx, malicious))
	require.Len(t, pool.execSQL, 1)
	require.Equal(t, "drop table "+pgx.Identifier{malicious}.Sanitize(), pool.execSQL[0])
	require.NotContains(t, pool.execSQL[0], "drop table "+malicious)

	require.NoError(t, c.dropTable(ctx, "normal_table"))
	require.Equal(t, `drop table "normal_table"`, pool.execSQL[1])
}

func TestAddColumnSanitizesTableAndColumnName(t *testing.T) {
	ctx := context.Background()
	pool := &recordingExecPool{}
	c := &Client{conn: pool, logger: zerolog.Nop()}

	maliciousTable := `evil"; drop table users; --`
	maliciousColumn := `col"; drop table users; --`
	col := schema.Column{Name: maliciousColumn, Type: arrow.BinaryTypes.String}
	require.NoError(t, c.addColumn(ctx, maliciousTable, col))
	require.Len(t, pool.execSQL, 1)
	want := "alter table " + pgx.Identifier{maliciousTable}.Sanitize() +
		" add column " + pgx.Identifier{maliciousColumn}.Sanitize() + " " + c.SchemaTypeToPg(col.Type)
	require.Equal(t, want, pool.execSQL[0])
}

func TestUpsertCrateDBSanitizesPrimaryKeyNames(t *testing.T) {
	c := &Client{}
	maliciousPK := `id") on conflict ("pg_class`
	table := &schema.Table{
		Name: "docs",
		Columns: schema.ColumnList{
			{Name: maliciousPK, Type: arrow.BinaryTypes.String, PrimaryKey: true},
			{Name: "body", Type: arrow.BinaryTypes.String},
		},
	}

	sql := c.upsertCrateDB(table)
	require.Contains(t, sql, " on conflict ("+pgx.Identifier{maliciousPK}.Sanitize()+") ")
	require.NotContains(t, sql, " on conflict ("+maliciousPK+") ")
}

// fakeTx implements pgx.Tx for finalizeTx tests; all methods except Commit and
// Rollback panic through the embedded nil interface, which is fine because
// finalizeTx never calls them.
type fakeTx struct {
	pgx.Tx
	committed  bool
	rolledBack bool
	commitErr  error
}

func (f *fakeTx) Commit(context.Context) error {
	f.committed = true
	return f.commitErr
}

func (f *fakeTx) Rollback(context.Context) error {
	f.rolledBack = true
	return nil
}

func TestFinalizeTx(t *testing.T) {
	ctx := context.Background()
	c := &Client{logger: zerolog.Nop()}

	t.Run("nil tx propagates error without panicking", func(t *testing.T) {
		wantErr := errors.New("begin transaction failed")
		// migrateToCQID reassigns tx from a failed BeginTx in the CockroachDB
		// path; the deferred rollback used to panic on the nil tx.
		err := c.finalizeTx(ctx, nil, "some_table", wantErr)
		require.ErrorIs(t, err, wantErr)
	})

	t.Run("nil tx with nil error is a no-op", func(t *testing.T) {
		require.NoError(t, c.finalizeTx(ctx, nil, "some_table", nil))
	})

	t.Run("commits on success", func(t *testing.T) {
		tx := &fakeTx{}
		require.NoError(t, c.finalizeTx(ctx, tx, "some_table", nil))
		require.True(t, tx.committed)
		require.False(t, tx.rolledBack)
	})

	t.Run("commit failure falls through to rollback", func(t *testing.T) {
		commitErr := errors.New("commit failed")
		tx := &fakeTx{commitErr: commitErr}
		err := c.finalizeTx(ctx, tx, "some_table", nil)
		require.ErrorIs(t, err, commitErr)
		require.True(t, tx.committed)
		require.True(t, tx.rolledBack)
	})

	t.Run("prior error triggers rollback", func(t *testing.T) {
		priorErr := errors.New("exec failed")
		tx := &fakeTx{}
		err := c.finalizeTx(ctx, tx, "some_table", priorErr)
		require.ErrorIs(t, err, priorErr)
		require.False(t, tx.committed)
		require.True(t, tx.rolledBack)
	})
}
