package queries

import (
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestSanitizeID(t *testing.T) {
	cases := []struct {
		name     string
		parts    []string
		expected string
	}{
		{name: "plain", parts: []string{"cq", "table"}, expected: `[cq].[table]`},
		{name: "closing bracket", parts: []string{"cq", "a]b"}, expected: `[cq].[a]]b]`},
		{name: "closing bracket in every part", parts: []string{"c]q", "t]"}, expected: `[c]]q].[t]]]`},
		{name: "opening bracket", parts: []string{"a[b"}, expected: `[a[b]`},
		{name: "NUL", parts: []string{"a\x00b"}, expected: `[ab]`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, sanitizeID(tc.parts...))
		})
	}
}

func bracketTable() *schema.Table {
	return &schema.Table{
		Name: "tbl] x",
		Columns: schema.ColumnList{
			schema.Column{Name: "id] x", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true, NotNull: true},
			schema.Column{Name: "val] x", Type: new(arrow.StringType)},
		},
	}
}

func TestCreateTableWithClosingBrackets(t *testing.T) {
	const expected = `CREATE TABLE [s]] x].[tbl]] x] (
  [id]] x] bigint NOT NULL,
  [val]] x] nvarchar(max),
  CONSTRAINT [tbl]] x_cqpk] PRIMARY KEY (
  [id]] x]
  )
);`
	require.Equal(t, expected, CreateTable("s] x", bracketTable()))
}

func TestDeleteStaleWithClosingBrackets(t *testing.T) {
	const expected = `DELETE FROM [s]] x].[tbl]] x] WHERE [_cq_source_name] = @sourceName AND [_cq_sync_time] < @syncTime;`
	query, _ := DeleteStale("s] x", &message.WriteDeleteStale{
		TableName:  "tbl] x",
		SourceName: "src",
		SyncTime:   time.Now(),
	})
	require.Equal(t, expected, query)
}

func TestTVPAddTypeWithClosingBrackets(t *testing.T) {
	const expected = `CREATE TYPE [s]] x].[cq_tbl_tbl]] x] AS TABLE (
  [id]] x] bigint NOT NULL,
  [val]] x] nvarchar(max)
);`
	require.Equal(t, expected, TVPAddType("s] x", bracketTable()))
}

func TestTVPAddProcWithClosingBrackets(t *testing.T) {
	const expected = `CREATE PROCEDURE [s]] x].[cq_proc_tbl]] x] @TVP [s]] x].[cq_tbl_tbl]] x] READONLY
AS
BEGIN
 SET NOCOUNT ON;
 UPDATE [tgt] WITH (UPDLOCK)
 SET
  [tgt].[val]] x] = [src].[val]] x]
 FROM [s]] x].[tbl]] x] AS [tgt]
 INNER JOIN @TVP AS [src]
 ON
  [tgt].[id]] x] = [src].[id]] x]
;

 INSERT [s]] x].[tbl]] x] (
  [id]] x],
  [val]] x]
 ) SELECT
  [src].[id]] x],
  [src].[val]] x]
 FROM @TVP AS [src]
 LEFT JOIN [s]] x].[tbl]] x] AS [tgt] ON (
  [tgt].[id]] x] = [src].[id]] x]
 ) WHERE (
  [tgt].[id]] x] IS NULL
);
END;`
	require.Equal(t, expected, TVPAddProc("s] x", bracketTable()))
}

func TestTVPQueryWithNonIdentifierColumnNames(t *testing.T) {
	table := bracketTable()
	builder := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	defer builder.Release()
	builder.Field(0).(*array.Int64Builder).Append(1)
	builder.Field(1).(*array.StringBuilder).Append("v")
	record := builder.NewRecordBatch()
	defer record.Release()

	query, params, err := TVPQuery("s] x", table, []arrow.RecordBatch{record})
	require.NoError(t, err)
	require.Equal(t, `exec [s]] x].[cq_proc_tbl]] x] @TVP;`, query)
	require.Len(t, params, 1)
}
