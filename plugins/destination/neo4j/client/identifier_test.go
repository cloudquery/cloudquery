package client

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
	"github.com/stretchr/testify/require"
)

const (
	hostileLabel       = "a` b) DETACH DELETE n //"
	quotedHostileLabel = "`a`` b) DETACH DELETE n //`"
	hostileKey         = "k` x) RETURN 1 //"
	quotedHostileKey   = "`k`` x) RETURN 1 //`"
)

func TestQuoteIdentifier(t *testing.T) {
	cases := map[string]string{
		"aws_ec2_instances":                  "`aws_ec2_instances`",
		"":                                   "``",
		"a`b":                                "`a``b`",
		"a``b":                               "`a````b`",
		"`":                                  "````",
		"a b)":                               "`a b)`",
		"a" + backslash + "b":                "`a" + backslash + "u005Cb`",
		backslash + backslash:                "`" + backslash + "u005C" + backslash + "u005C`",
		backslash + "`":                      "`" + backslash + "u005C```",
		"a" + backslash + "u0060b":           "`a" + backslash + "u005Cu0060b`",
		"x" + backslash + "uu0060":           "`x" + backslash + "u005Cuu0060`",
		"x" + backslash + "uuu0060":          "`x" + backslash + "u005Cuuu0060`",
		"a" + backslash + "u0060" + "`b":     "`a" + backslash + "u005Cu0060``b`",
		"a" + backslash + "u005C" + "u0060b": "`a" + backslash + "u005Cu005Cu0060b`",
	}
	for name, want := range cases {
		quoted, err := quoteIdentifier(name)
		require.NoError(t, err)
		require.Equal(t, want, quoted, "name %q", name)
	}
}

func statement(t *testing.T) func(string, error) string {
	return func(stmt string, err error) string {
		t.Helper()
		require.NoError(t, err)
		return stmt
	}
}

const (
	unicodeEscapeLabel       = "x" + backslash + "uu0060) SET t.injected = true //"
	quotedUnicodeEscapeLabel = "`x" + backslash + "u005Cuu0060) SET t.injected = true //`"
	unicodeEscapeKey         = "a" + backslash + "u0060b"
	quotedUnicodeEscapeKey   = "`a" + backslash + "u005Cu0060b`"
)

func TestStatementsEscapeUnicodeEscapes(t *testing.T) {
	s := statement(t)
	require.Equal(t, "UNWIND $rows AS row CREATE (t:"+quotedUnicodeEscapeLabel+") SET t = row", s(writeStatement(unicodeEscapeLabel, nil)))
	require.Equal(t, "UNWIND $rows AS row MERGE (t:`pk` {"+quotedUnicodeEscapeKey+": row."+quotedUnicodeEscapeKey+"}) SET t = row", s(writeStatement("pk", []string{unicodeEscapeKey})))
	require.Equal(t, "MATCH (t:"+quotedUnicodeEscapeLabel+") RETURN t", s(readStatement(unicodeEscapeLabel)))
	require.Equal(t, "MATCH (n:"+quotedUnicodeEscapeLabel+") WHERE n._cq_source_name = $cq_source_name AND n._cq_sync_time < $cq_sync_time DETACH DELETE n", s(deleteStaleStatement(unicodeEscapeLabel)))

	table := &schema.Table{Name: unicodeEscapeLabel, Columns: schema.ColumnList{{Name: unicodeEscapeKey, PrimaryKey: true}}}
	quotedIndex := "`_cq_index_x" + backslash + "u005Cuu0060) SET t.injected = true //`"
	require.Equal(t, "CREATE INDEX "+quotedIndex+" IF NOT EXISTS FOR (n:"+quotedUnicodeEscapeLabel+") ON (n."+quotedUnicodeEscapeKey+");", s(createIndexQuery(table)))
	require.Equal(t, "DROP INDEX "+quotedIndex+" IF EXISTS;", s(dropIndexQuery(table)))
}

func requireOnlyQuotedIdentifiers(t *testing.T, stmt string, quoted ...string) {
	t.Helper()
	for _, identifier := range quoted {
		require.Contains(t, stmt, identifier)
		stmt = strings.ReplaceAll(stmt, identifier, "")
	}
	require.NotContains(t, stmt, "`")
	require.NotContains(t, stmt, "DETACH DELETE n //")
	require.NotContains(t, stmt, "RETURN 1")
}

func TestWriteStatementQuotesIdentifiers(t *testing.T) {
	s := statement(t)
	create := s(writeStatement(hostileLabel, nil))
	require.Equal(t, "UNWIND $rows AS row CREATE (t:"+quotedHostileLabel+") SET t = row", create)
	requireOnlyQuotedIdentifiers(t, create, quotedHostileLabel)

	merge := s(writeStatement(hostileLabel, []string{hostileKey, "id"}))
	require.Equal(t, "UNWIND $rows AS row MERGE (t:"+quotedHostileLabel+" {"+quotedHostileKey+": row."+quotedHostileKey+", `id`: row.`id`}) SET t = row", merge)
	requireOnlyQuotedIdentifiers(t, merge, quotedHostileLabel, quotedHostileKey, "`id`")
}

func TestDeleteStaleStatementQuotesLabel(t *testing.T) {
	s := statement(t)
	stmt := s(deleteStaleStatement(hostileLabel))
	require.Equal(t, "MATCH (n:"+quotedHostileLabel+") WHERE n._cq_source_name = $cq_source_name AND n._cq_sync_time < $cq_sync_time DETACH DELETE n", stmt)
	requireOnlyQuotedIdentifiers(t, strings.TrimSuffix(stmt, " DETACH DELETE n"), quotedHostileLabel)
}

func TestReadStatementQuotesLabel(t *testing.T) {
	s := statement(t)
	stmt := s(readStatement(hostileLabel))
	require.Equal(t, "MATCH (t:"+quotedHostileLabel+") RETURN t", stmt)
	requireOnlyQuotedIdentifiers(t, stmt, quotedHostileLabel)
}

func hostileTable() *schema.Table {
	return &schema.Table{
		Name: hostileLabel,
		Columns: schema.ColumnList{
			{Name: hostileKey, PrimaryKey: true},
			{Name: "id", PrimaryKey: true},
		},
	}
}

func TestIndexQueriesQuoteIdentifiers(t *testing.T) {
	s := statement(t)
	quotedIndex := "`_cq_index_a`` b) DETACH DELETE n //`"

	create := s(createIndexQuery(hostileTable()))
	require.Equal(t, "CREATE INDEX "+quotedIndex+" IF NOT EXISTS FOR (n:"+quotedHostileLabel+") ON (n."+quotedHostileKey+", n.`id`);", create)
	requireOnlyQuotedIdentifiers(t, create, quotedIndex, quotedHostileLabel, quotedHostileKey, "`id`")

	drop := s(dropIndexQuery(hostileTable()))
	require.Equal(t, "DROP INDEX "+quotedIndex+" IF EXISTS;", drop)
	requireOnlyQuotedIdentifiers(t, drop, quotedIndex)
}

func TestIndexNameUnchangedForOrdinaryTable(t *testing.T) {
	s := statement(t)
	table := &schema.Table{Name: "aws_ec2_instances", Columns: schema.ColumnList{{Name: "arn", PrimaryKey: true}}}
	require.Equal(t, "_cq_index_aws_ec2_instances", indexName(table))
	require.Equal(t, "CREATE INDEX `_cq_index_aws_ec2_instances` IF NOT EXISTS FOR (n:`aws_ec2_instances`) ON (n.`arn`);", s(createIndexQuery(table)))
}

func stringTable(name string, primaryKey string) *schema.Table {
	return &schema.Table{Name: name, Columns: schema.ColumnList{{Name: primaryKey, Type: arrow.BinaryTypes.String, PrimaryKey: true}}}
}

func singleStringRecord(table *schema.Table, value string) arrow.RecordBatch {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	defer bldr.Release()
	bldr.Field(0).(*array.StringBuilder).Append(value)
	return bldr.NewRecordBatch()
}

func runQuery(t *testing.T, c *Client, query string, params map[string]any) []*neo4j.Record {
	t.Helper()
	ctx := context.Background()
	sess := c.Session(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer sess.Close(ctx)
	result, err := sess.Run(ctx, query, params)
	require.NoError(t, err)
	records, err := result.Collect(ctx)
	require.NoError(t, err)
	return records
}

func removeTable(t *testing.T, c *Client, table *schema.Table) {
	s := statement(t)
	runQuery(t, c, "MATCH (n) WHERE $label IN labels(n) DETACH DELETE n", map[string]any{"label": table.Name})
	runQuery(t, c, s(dropIndexQuery(table)), nil)
}

func migrateWriteRead(t *testing.T, c *Client, table *schema.Table, value string) []arrow.RecordBatch {
	t.Helper()
	ctx := context.Background()
	t.Cleanup(func() { removeTable(t, c, table) })
	require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: table}}))
	require.NoError(t, c.WriteTableBatch(ctx, table.Name, message.WriteInserts{{Record: singleStringRecord(table, value)}}))
	read := make(chan arrow.RecordBatch, 10)
	require.NoError(t, c.Read(ctx, table, read))
	close(read)
	var records []arrow.RecordBatch
	for record := range read {
		records = append(records, record)
	}
	return records
}

func requireSingleValue(t *testing.T, records []arrow.RecordBatch, want string) {
	t.Helper()
	require.Len(t, records, 1)
	require.Equal(t, want, records[0].Column(0).(*array.String).Value(0))
}

func TestUnicodeEscapeInTableNameIsNotInjected(t *testing.T) {
	c := newConnectedClient(t)
	table := stringTable(unicodeEscapeLabel, "id")

	requireSingleValue(t, migrateWriteRead(t, c, table, "row-1"), "row-1")

	require.Empty(t, runQuery(t, c, "MATCH (n) WHERE n.injected IS NOT NULL RETURN n", nil))
	labels := runQuery(t, c, "MATCH (n) WHERE n.id = $id RETURN labels(n) AS labels", map[string]any{"id": "row-1"})
	require.Len(t, labels, 1)
	got, _ := labels[0].Get("labels")
	require.Equal(t, []any{unicodeEscapeLabel}, got)
}

func TestUnicodeEscapeInPrimaryKeyNameKeepsValue(t *testing.T) {
	c := newConnectedClient(t)
	table := stringTable(fmt.Sprintf("pk_escape_%d", time.Now().UnixNano()), unicodeEscapeKey)

	requireSingleValue(t, migrateWriteRead(t, c, table, "pk-1"), "pk-1")
	require.Equal(t, []string{unicodeEscapeKey}, indexedProperties(t, c, table))
}

const nulByte = string(rune(0))

func keyValueTable(name string) *schema.Table {
	return &schema.Table{Name: name, Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "value", Type: arrow.BinaryTypes.String},
	}}
}

func keyValueRecord(table *schema.Table, id, value string) arrow.RecordBatch {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	defer bldr.Release()
	bldr.Field(0).(*array.StringBuilder).Append(id)
	bldr.Field(1).(*array.StringBuilder).Append(value)
	return bldr.NewRecordBatch()
}

func TestNULInTableNameLeavesExistingTableUntouched(t *testing.T) {
	ctx := context.Background()
	c := newConnectedClient(t)
	suffix := fmt.Sprintf("_%d", time.Now().UnixNano())
	existing := keyValueTable("ab" + suffix)
	aliased := keyValueTable(nulLabel + suffix)
	t.Cleanup(func() { removeTable(t, c, existing) })

	require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: existing}}))
	require.NoError(t, c.WriteTableBatch(ctx, existing.Name, message.WriteInserts{{Record: keyValueRecord(existing, "row-1", "original")}}))

	migrateErr := c.MigrateTables(ctx, message.WriteMigrateTables{{Table: aliased}})
	writeErr := c.WriteTableBatch(ctx, aliased.Name, message.WriteInserts{{Record: keyValueRecord(aliased, "row-1", "overwritten")}})

	values := runQuery(t, c, "MATCH (n) WHERE $label IN labels(n) RETURN n.value AS value", map[string]any{"label": existing.Name})
	require.Len(t, values, 1)
	got, _ := values[0].Get("value")
	require.Equal(t, "original", got)
	require.ErrorIs(t, migrateErr, errNULInIdentifier)
	require.ErrorIs(t, writeErr, errNULInIdentifier)
}

const nulLabel = "a" + nulByte + "b"

func TestStatementsRejectNUL(t *testing.T) {
	_, err := quoteIdentifier(nulLabel)
	require.ErrorIs(t, err, errNULInIdentifier)
	_, err = writeStatement(nulLabel, nil)
	require.ErrorIs(t, err, errNULInIdentifier)
	_, err = writeStatement("t", []string{"id", nulLabel})
	require.ErrorIs(t, err, errNULInIdentifier)
	_, err = readStatement(nulLabel)
	require.ErrorIs(t, err, errNULInIdentifier)
	_, err = deleteStaleStatement(nulLabel)
	require.ErrorIs(t, err, errNULInIdentifier)
	_, err = createIndexQuery(stringTable(nulLabel, "id"))
	require.ErrorIs(t, err, errNULInIdentifier)
	_, err = createIndexQuery(&schema.Table{Name: nulLabel})
	require.ErrorIs(t, err, errNULInIdentifier)
	_, err = createIndexQuery(stringTable("t", nulLabel))
	require.ErrorIs(t, err, errNULInIdentifier)
	_, err = dropIndexQuery(stringTable(nulLabel, "id"))
	require.ErrorIs(t, err, errNULInIdentifier)
}

func TestEntryPointsRejectNULBeforeOpeningSession(t *testing.T) {
	ctx := context.Background()
	c := &Client{}
	tables := map[string]*schema.Table{
		"table name":       stringTable(nulLabel, "id"),
		"primary key name": stringTable("t", nulLabel),
	}
	for name, table := range tables {
		t.Run(name, func(t *testing.T) {
			require.ErrorIs(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: stringTable("ok", "id")}, {Table: table}}), errNULInIdentifier)
			require.ErrorIs(t, c.WriteTableBatch(ctx, table.Name, message.WriteInserts{{Record: singleStringRecord(table, "v")}}), errNULInIdentifier)
			_, err := c.AssessTables(ctx, []plugin.TablePair{{New: table}}, plugin.AssessOptions{})
			require.ErrorIs(t, err, errNULInIdentifier)
		})
	}
	require.ErrorIs(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: &schema.Table{Name: nulLabel}}}), errNULInIdentifier)
	require.ErrorIs(t, c.Read(ctx, stringTable(nulLabel, "id"), make(chan arrow.RecordBatch)), errNULInIdentifier)
	require.ErrorIs(t, c.DeleteStale(ctx, message.WriteDeleteStales{{TableName: "ok"}, {TableName: nulLabel}}), errNULInIdentifier)
}
