package client

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

const (
	backtickTableName    = "t` DROP COLUMN `x"
	quotedBacktickTable  = "`t`` DROP COLUMN ``x`"
	backtickColumnName   = "c` = 1 OR `y"
	quotedBacktickColumn = "`c`` = 1 OR ``y`"
)

func backtickTable() *schema.Table {
	return &schema.Table{
		Name: backtickTableName,
		Columns: schema.ColumnList{
			{Name: backtickColumnName, Type: arrow.BinaryTypes.String, PrimaryKey: true},
			{Name: "plain", Type: arrow.PrimitiveTypes.Int64, Unique: true},
		},
	}
}

func TestIdentifier(t *testing.T) {
	cases := map[string]string{
		"plain":           "`plain`",
		"":                "``",
		"with space":      "`with space`",
		"a`b":             "`a``b`",
		"`":               "````",
		"a``b":            "`a````b`",
		backtickTableName: quotedBacktickTable,
	}
	for name, want := range cases {
		require.Equal(t, want, identifier(name))
	}
}

func TestQueriesQuoteNamesWithBackticks(t *testing.T) {
	table := backtickTable()
	column := &table.Columns[0]
	c := &Client{maxIndexLength: 768}

	cases := map[string]struct {
		got  string
		want string
	}{
		"create table": {
			got: c.createTableQuery(table),
			want: "CREATE TABLE " + quotedBacktickTable + " (\n  " +
				quotedBacktickColumn + " text,\n  `plain` bigint UNIQUE,\n   PRIMARY KEY (" + quotedBacktickColumn + "(768))\n) CHARACTER SET utf8mb4;",
		},
		"add column": {
			got:  addColumnQuery(table, column),
			want: "ALTER TABLE " + quotedBacktickTable + " ADD COLUMN " + quotedBacktickColumn + " text;",
		},
		"drop index": {
			got:  dropIndexQuery(table, column),
			want: "ALTER TABLE " + quotedBacktickTable + " DROP INDEX " + quotedBacktickColumn + ";",
		},
		"drop table": {
			got:  dropTableQuery(table),
			want: "DROP TABLE " + quotedBacktickTable,
		},
		"insert": {
			got:  getInsertQueryBuild(table, 1, false).String(),
			want: "INSERT INTO " + quotedBacktickTable + " (" + quotedBacktickColumn + ", `plain`) VALUES (?,?)",
		},
		"upsert": {
			got: getInsertQueryBuild(table, 1, true).String(),
			want: "INSERT INTO " + quotedBacktickTable + " (" + quotedBacktickColumn + ", `plain`) VALUES (?,?) ON DUPLICATE KEY UPDATE " +
				quotedBacktickColumn + " = VALUES(" + quotedBacktickColumn + "), `plain` = VALUES(`plain`)",
		},
		"delete stale": {
			got:  deleteStaleQuery(backtickTableName),
			want: "delete from " + quotedBacktickTable + " where `_cq_source_name` = ? and `_cq_sync_time` < ?",
		},
		"read": {
			got:  readQuery(table),
			want: "SELECT " + quotedBacktickColumn + ", `plain` FROM " + quotedBacktickTable,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.want, tc.got)
		})
	}
}
