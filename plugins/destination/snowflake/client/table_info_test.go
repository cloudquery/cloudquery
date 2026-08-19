package client

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTableInfoAllTables(t *testing.T) {
	cols := []fakeColumn{
		{name: "_cq_id", dataType: "text"},
		{name: "name", dataType: "text", nullable: true},
	}
	c, fake := newFakeSnowflake(t, map[string][]fakeColumn{
		"TABLE_A": cols,
		"TABLE_B": cols,
		"TABLE_C": cols,
	})

	got, _, err := c.getTableInfo(context.Background(), []string{"table_c", "table_a", "table_b"})
	require.NoError(t, err)

	// A wide-row VALUES clause matched only TABLE_A, leaving the rest to look non-existent.
	require.ElementsMatch(t, []string{"TABLE_A", "TABLE_B", "TABLE_C"}, got.TableNames())
	for _, table := range got {
		require.Len(t, table.Columns, len(cols), "table %s", table.Name)
	}
	require.Equal(t, [][]string{{"TABLE_A", "TABLE_B", "TABLE_C"}}, fake.matched())
}

func TestGetTableInfoBatches(t *testing.T) {
	const numTables = 450 // > 2 batches of the 200-table limit

	catalog := make(map[string][]fakeColumn, numTables)
	tableNames := make([]string, 0, numTables)
	for i := range numTables {
		name := fmt.Sprintf("table_%03d", i)
		tableNames = append(tableNames, name)
		catalog[strings.ToUpper(name)] = []fakeColumn{{name: "_cq_id", dataType: "text"}}
	}
	c, fake := newFakeSnowflake(t, catalog)

	got, _, err := c.getTableInfo(context.Background(), tableNames)
	require.NoError(t, err)
	require.Len(t, got, numTables)

	var sizes []int
	for _, batch := range fake.matched() {
		sizes = append(sizes, len(batch))
	}
	require.Equal(t, []int{200, 200, 50}, sizes)
}

func TestGetTableInfoNoTables(t *testing.T) {
	c, fake := newFakeSnowflake(t, nil)

	got, _, err := c.getTableInfo(context.Background(), nil)
	require.NoError(t, err)
	require.Empty(t, got)
	require.Empty(t, fake.matched())
}

type fakeColumn struct {
	name     string
	dataType string
	nullable bool
}

// fakeSnowflake answers the three queries getTableInfo issues, reproducing Snowflake's
// `SELECT COLUMN1 FROM VALUES ...` semantics.
type fakeSnowflake struct {
	catalog map[string][]fakeColumn // uppercase table name vs. its columns

	mu      sync.Mutex
	matches [][]string // names COLUMN1 yielded, one entry per table-info query
}

func (f *fakeSnowflake) matched() [][]string {
	f.mu.Lock()
	defer f.mu.Unlock()

	return append([][]string(nil), f.matches...)
}

var (
	fakeSnowflakeMu sync.Mutex
	fakeSnowflakes  = map[string]*fakeSnowflake{}
)

func init() { sql.Register("snowflake_fake", fakeSnowflakeDriver{}) }

func newFakeSnowflake(t *testing.T, catalog map[string][]fakeColumn) (*Client, *fakeSnowflake) {
	t.Helper()

	fake := &fakeSnowflake{catalog: catalog}
	fakeSnowflakeMu.Lock()
	fakeSnowflakes[t.Name()] = fake
	fakeSnowflakeMu.Unlock()
	t.Cleanup(func() {
		fakeSnowflakeMu.Lock()
		delete(fakeSnowflakes, t.Name())
		fakeSnowflakeMu.Unlock()
	})

	db, err := sql.Open("snowflake_fake", t.Name())
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, db.Close()) })

	c := &Client{db: db}
	c.spec.SetDefaults() // MigrateConcurrency, without which the errgroup limit is 0
	return c, fake
}

type fakeSnowflakeDriver struct{}

func (fakeSnowflakeDriver) Open(dsn string) (driver.Conn, error) {
	fakeSnowflakeMu.Lock()
	defer fakeSnowflakeMu.Unlock()

	fake, ok := fakeSnowflakes[dsn]
	if !ok {
		return nil, fmt.Errorf("no fake Snowflake registered for %q", dsn)
	}
	return &fakeSnowflakeConn{fake: fake}, nil
}

type fakeSnowflakeConn struct {
	fake *fakeSnowflake
}

func (*fakeSnowflakeConn) Prepare(string) (driver.Stmt, error) { return nil, driver.ErrSkip }
func (*fakeSnowflakeConn) Close() error                        { return nil }
func (*fakeSnowflakeConn) Begin() (driver.Tx, error)           { return nil, driver.ErrSkip }

func (c *fakeSnowflakeConn) QueryContext(_ context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	switch query {
	case sqlShowPrimaryKeys, sqlShowUniques:
		return &fakeRows{cols: []string{"table_name", "column_name", "constraint_name", "key_sequence"}}, nil
	}

	if !strings.HasPrefix(query, sqlTableInfoStart) || !strings.HasSuffix(query, sqlTableInfoEnd) {
		return nil, fmt.Errorf("unexpected query: %s", query)
	}
	values := strings.TrimSuffix(strings.TrimPrefix(query, sqlTableInfoStart), sqlTableInfoEnd)

	wanted, err := selectColumn1(values, args)
	if err != nil {
		return nil, err
	}

	var (
		matched []string
		rows    [][]driver.Value
	)
	for _, name := range wanted {
		cols, ok := c.fake.catalog[strings.ToUpper(name)]
		if !ok {
			continue
		}
		matched = append(matched, strings.ToUpper(name))
		for _, col := range cols {
			nullable := "NO"
			if col.nullable {
				nullable = "YES"
			}
			rows = append(rows, []driver.Value{strings.ToUpper(name), col.name, col.dataType, nullable})
		}
	}

	c.fake.mu.Lock()
	c.fake.matches = append(c.fake.matches, matched)
	c.fake.mu.Unlock()

	return &fakeRows{
		cols: []string{"table_name", "column_name", "data_type", "is_nullable"},
		rows: rows,
	}, nil
}

// selectColumn1 evaluates `SELECT COLUMN1 FROM VALUES <clause>`: each parenthesized group is one
// row and only its first value reaches COLUMN1, so `(?,?,?)` yields one name and `(?),(?),(?)` three.
func selectColumn1(clause string, args []driver.NamedValue) ([]string, error) {
	rows := strings.Split(strings.TrimSuffix(strings.TrimPrefix(clause, "("), ")"), "),(")
	column1 := make([]string, 0, len(rows))
	arg := 0
	for _, row := range rows {
		width := strings.Count(row, "?")
		if width == 0 || arg+width > len(args) {
			return nil, fmt.Errorf("malformed VALUES clause %q for %d args", clause, len(args))
		}
		name, ok := args[arg].Value.(string)
		if !ok {
			return nil, fmt.Errorf("expected a string arg at position %d, got %T", arg, args[arg].Value)
		}
		column1 = append(column1, name)
		arg += width
	}
	if arg != len(args) {
		return nil, fmt.Errorf("VALUES clause %q consumes %d of the %d bound args", clause, arg, len(args))
	}
	return column1, nil
}

type fakeRows struct {
	cols []string
	rows [][]driver.Value
	next int
}

func (r *fakeRows) Columns() []string { return r.cols }
func (*fakeRows) Close() error        { return nil }

func (r *fakeRows) Next(dest []driver.Value) error {
	if r.next >= len(r.rows) {
		return io.EOF
	}
	copy(dest, r.rows[r.next])
	r.next++
	return nil
}
