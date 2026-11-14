package client

import (
	"context"
	"database/sql"
	"testing"

	"github.com/duckdb/duckdb-go/v2"
	"github.com/stretchr/testify/require"
)

func TestGetTableInfoWithQuoteInTableName(t *testing.T) {
	ctx := context.Background()
	connector, err := duckdb.NewConnector("?threads=1", useUTC)
	require.NoError(t, err)
	c := &Client{db: sql.OpenDB(connector)}
	t.Cleanup(func() { require.NoError(t, c.db.Close()) })

	const tableName = `foo'bar`
	_, err = c.db.ExecContext(ctx, `CREATE TABLE "foo'bar" (id BIGINT, name VARCHAR)`)
	require.NoError(t, err)

	info, err := c.getTableInfo(ctx, tableName)
	require.NoError(t, err)
	require.NotNil(t, info)
	require.Len(t, info.Columns, 2)
	require.Equal(t, "id", info.Columns[0].Name)
	require.Equal(t, "name", info.Columns[1].Name)
}
