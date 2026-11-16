package client

import (
	"context"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestValidateTableName(t *testing.T) {
	valid := []string{
		"logs",
		"aws_ec2_instances",
		"logs-app",
		"logs.app",
		"logs+app",
		"Logs",
		"a",
		strings.Repeat("a", 255),
	}
	for _, name := range valid {
		t.Run("valid/"+name, func(t *testing.T) {
			require.NoError(t, validateTableName(name))
		})
	}

	invalid := []string{
		"",
		".",
		"..",
		"*",
		"logs*",
		"logs,other",
		"logs,-other",
		"logs?",
		"logs/other",
		`logs\app`,
		"logs#other",
		`logs"other`,
		"logs<other",
		"logs>other",
		"logs|other",
		"logs other",
		"cluster:logs",
		"_all",
		"-logs",
		"+logs",
		strings.Repeat("a", 256),
	}
	for _, name := range invalid {
		t.Run("invalid/"+name, func(t *testing.T) {
			err := validateTableName(name)
			require.Error(t, err)
			require.Contains(t, err.Error(), strconv.Quote(name))
		})
	}
}

func TestIsTableIndex(t *testing.T) {
	cases := []struct {
		table string
		index string
		want  bool
	}{
		{"logs", "logs", true},
		{"logs", "logs-2026-01-01", true},
		{"logs", "logs-app-2026-01-01", false},
		{"logs", "logs-app", false},
		{"logs", "logs-2026-01-01-extra", false},
		{"logs", "logs-2026-1-1", false},
		{"logs", "logs-2026-13-01", false},
		{"logs", "logs-", false},
		{"logs", "logs2-2026-01-01", false},
		{"logs", "other-2026-01-01", false},
		{"logs-app", "logs-app-2026-01-01", true},
	}
	for _, tc := range cases {
		t.Run(tc.table+"/"+tc.index, func(t *testing.T) {
			require.Equal(t, tc.want, isTableIndex(tc.table, tc.index))
		})
	}
}

func TestTableIndices(t *testing.T) {
	got := tableIndices("logs", []string{"logs-app-2026-01-01", "logs-2026-01-01", "logs-app", "logs-2026-01-02"})
	require.Equal(t, []string{"logs-2026-01-01", "logs-2026-01-02"}, got)
}

func TestTableIndicesMatchesOwnIndexName(t *testing.T) {
	c := &Client{}
	table := &schema.Table{Name: "logs", Columns: schema.ColumnList{{Name: "id", Type: arrow.BinaryTypes.String}}}
	require.True(t, isTableIndex(table.Name, c.getIndexName(table, time.Now())))
}

func TestInvalidTableNameIsRejectedBeforeAnyRequest(t *testing.T) {
	const name = "logs,other*"
	ctx := context.Background()
	c := &Client{spec: &Spec{}}
	table := &schema.Table{Name: name, Columns: schema.ColumnList{{Name: "id", Type: arrow.BinaryTypes.String}}}

	t.Run("migrate", func(t *testing.T) {
		err := c.MigrateTables(ctx, message.WriteMigrateTables{{Table: table, MigrateForce: true}})
		require.ErrorContains(t, err, name)
	})
	t.Run("write", func(t *testing.T) {
		bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
		bldr.Field(0).(*array.StringBuilder).Append("a")
		err := c.WriteTableBatch(ctx, name, message.WriteInserts{{Record: bldr.NewRecordBatch()}})
		require.ErrorContains(t, err, name)
	})
	t.Run("delete stale", func(t *testing.T) {
		err := c.DeleteStale(ctx, message.WriteDeleteStales{{TableName: name, SourceName: "source", SyncTime: time.Now()}})
		require.ErrorContains(t, err, name)
	})
	t.Run("read", func(t *testing.T) {
		err := c.Read(ctx, table, make(chan arrow.RecordBatch))
		require.ErrorContains(t, err, name)
	})
}
