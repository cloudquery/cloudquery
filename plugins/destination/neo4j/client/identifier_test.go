package client

import (
	"strings"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/schema"
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
		"aws_ec2_instances": "`aws_ec2_instances`",
		"":                  "``",
		"a`b":               "`a``b`",
		"a``b":              "`a````b`",
		"`":                 "````",
		"a b)":              "`a b)`",
		"a\x00b":            "`ab`",
		`a\u0060b`:          "`a``b`",
		`a\u0060` + "`b":    "`a````b`",
	}
	for name, want := range cases {
		require.Equal(t, want, quoteIdentifier(name), "name %q", name)
	}
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
	create := writeStatement(hostileLabel, nil)
	require.Equal(t, "UNWIND $rows AS row CREATE (t:"+quotedHostileLabel+") SET t = row", create)
	requireOnlyQuotedIdentifiers(t, create, quotedHostileLabel)

	merge := writeStatement(hostileLabel, []string{hostileKey, "id"})
	require.Equal(t, "UNWIND $rows AS row MERGE (t:"+quotedHostileLabel+" {"+quotedHostileKey+": row."+quotedHostileKey+", `id`: row.`id`}) SET t = row", merge)
	requireOnlyQuotedIdentifiers(t, merge, quotedHostileLabel, quotedHostileKey, "`id`")
}

func TestDeleteStaleStatementQuotesLabel(t *testing.T) {
	stmt := deleteStaleStatement(hostileLabel)
	require.Equal(t, "MATCH (n:"+quotedHostileLabel+") WHERE n._cq_source_name = $cq_source_name AND n._cq_sync_time < $cq_sync_time DETACH DELETE n", stmt)
	requireOnlyQuotedIdentifiers(t, strings.TrimSuffix(stmt, " DETACH DELETE n"), quotedHostileLabel)
}

func TestReadStatementQuotesLabel(t *testing.T) {
	stmt := readStatement(hostileLabel)
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
	quotedIndex := "`_cq_index_a`` b) DETACH DELETE n //`"

	create := createIndexQuery(hostileTable())
	require.Equal(t, "CREATE INDEX "+quotedIndex+" IF NOT EXISTS FOR (n:"+quotedHostileLabel+") ON (n."+quotedHostileKey+", n.`id`);", create)
	requireOnlyQuotedIdentifiers(t, create, quotedIndex, quotedHostileLabel, quotedHostileKey, "`id`")

	drop := dropIndexQuery(hostileTable())
	require.Equal(t, "DROP INDEX "+quotedIndex+" IF EXISTS;", drop)
	requireOnlyQuotedIdentifiers(t, drop, quotedIndex)
}

func TestIndexNameUnchangedForOrdinaryTable(t *testing.T) {
	table := &schema.Table{Name: "aws_ec2_instances", Columns: schema.ColumnList{{Name: "arn", PrimaryKey: true}}}
	require.Equal(t, "_cq_index_aws_ec2_instances", indexName(table))
	require.Equal(t, "CREATE INDEX `_cq_index_aws_ec2_instances` IF NOT EXISTS FOR (n:`aws_ec2_instances`) ON (n.`arn`);", createIndexQuery(table))
}
