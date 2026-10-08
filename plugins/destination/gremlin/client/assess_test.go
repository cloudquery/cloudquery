package client

import (
	"context"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func newNoConnectionClient(t *testing.T) *Client {
	t.Helper()
	c, err := New(context.Background(), zerolog.Nop(), nil, plugin.NewClientOptions{NoConnection: true})
	require.NoError(t, err)
	return c.(*Client)
}

func assessTables(t *testing.T, options plugin.AssessOptions, pairs ...plugin.TablePair) []plugin.TableFinding {
	t.Helper()
	p := plugin.NewPlugin("gremlin", "development", New)
	require.NoError(t, p.Init(context.Background(), nil, plugin.NewClientOptions{NoConnection: true}))
	findings, err := p.AssessTables(context.Background(), pairs, options)
	require.NoError(t, err)
	return findings
}

func noChangeColumn(name string) plugin.ColumnFinding {
	return plugin.ColumnFinding{
		ColumnName:         name,
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
	}
}

func TestNewWithNoConnectionSkipsConnection(t *testing.T) {
	c := newNoConnectionClient(t)
	require.Nil(t, c.client)
	require.Nil(t, c.writer)
	require.Error(t, c.Close(context.Background()))
}

func TestAssessTablesTypeChange(t *testing.T) {
	oldTable := &schema.Table{Name: "datadog_monitors", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "tags", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].Type = types.ExtensionTypes.JSON

	expected := []plugin.TableFinding{{
		TableName:          "datadog_monitors",
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
		Columns:            []plugin.ColumnFinding{noChangeColumn("tags")},
	}}
	for _, force := range []bool{false, true} {
		require.Equal(t, expected, assessTables(t, plugin.AssessOptions{MigrateForce: force}, plugin.TablePair{Old: oldTable, New: newTable}))
		require.NoError(t, newNoConnectionClient(t).MigrateTables(context.Background(), message.WriteMigrateTables{{Table: newTable, MigrateForce: force}}))
	}
}

func TestAssessTablesColumnChanges(t *testing.T) {
	oldTable := &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "name", Type: arrow.BinaryTypes.String},
	}}
	newTable := &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "policy_id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "actions", Type: types.ExtensionTypes.JSON},
	}}

	require.Equal(t, []plugin.TableFinding{{
		TableName:          "okta_policy_rules",
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
		Columns: []plugin.ColumnFinding{
			noChangeColumn("policy_id"),
			noChangeColumn("actions"),
			noChangeColumn("name"),
		},
	}}, assessTables(t, plugin.AssessOptions{}, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesWholeTable(t *testing.T) {
	table := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}

	tests := []struct {
		name     string
		pair     plugin.TablePair
		expected plugin.TableFinding
	}{
		{
			name:     "no change",
			pair:     plugin.TablePair{Old: table, New: table},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorNoChange, ForcedModeBehavior: behaviorNoChange},
		},
		{
			name:     "added table",
			pair:     plugin.TablePair{New: table},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorNoChange, ForcedModeBehavior: behaviorNoChange},
		},
		{
			name:     "removed table",
			pair:     plugin.TablePair{Old: table},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryTableRemoved, SafeModeBehavior: behaviorKeepVertices, ForcedModeBehavior: behaviorKeepVertices},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, []plugin.TableFinding{tc.expected}, assessTables(t, plugin.AssessOptions{}, tc.pair))
		})
	}
}
