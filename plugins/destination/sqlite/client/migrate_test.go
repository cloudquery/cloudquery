package client

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTableInfoWithQuoteInTableName(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	c := &Client{db: db}
	t.Cleanup(func() { require.NoError(t, db.Close()) })

	const tableName = `foo'bar`
	_, err = db.Exec(`CREATE TABLE "foo'bar" (id INTEGER, name TEXT)`)
	require.NoError(t, err)

	info, err := c.getTableInfo(tableName)
	require.NoError(t, err)
	require.NotNil(t, info)
	require.Len(t, info.columns, 2)
	require.Equal(t, "id", info.columns[0].name)
	require.Equal(t, "name", info.columns[1].name)
}
