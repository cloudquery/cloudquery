package client

import (
	"context"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
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

func assessOneTable(t *testing.T, pair plugin.TablePair) plugin.TableFinding {
	t.Helper()
	findings, err := newNoConnectionClient(t).AssessTables(context.Background(), []plugin.TablePair{pair}, plugin.AssessOptions{})
	require.NoError(t, err)
	require.Len(t, findings, 1)
	return findings[0]
}

func withoutPrimaryKeys(table *schema.Table) *schema.Table {
	result := table.Copy(nil)
	for i := range result.Columns {
		result.Columns[i].PrimaryKey = false
	}
	return result
}

func noChangeColumn(name string) plugin.ColumnFinding {
	return plugin.ColumnFinding{ColumnName: name, Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorNoChange, ForcedModeBehavior: behaviorNoChange}
}

func TestNewWithNoConnectionSkipsClient(t *testing.T) {
	c := newNoConnectionClient(t)
	require.Nil(t, c.client)
	require.Nil(t, c.writer)
	require.Error(t, c.Close(context.Background()))
}

func datadogTagsPair() plugin.TablePair {
	oldTable := &schema.Table{Name: "datadog_monitors", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "tags", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].Type = types.ExtensionTypes.JSON
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func oktaPolicyIDPair() plugin.TablePair {
	oldTable := &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = append(newTable.Columns,
		schema.Column{Name: "policy_id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		schema.Column{Name: "actions", Type: types.ExtensionTypes.JSON},
		schema.Column{Name: "conditions", Type: types.ExtensionTypes.JSON},
	)
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func withoutPrimaryKeysPair(pair plugin.TablePair) plugin.TablePair {
	return plugin.TablePair{Old: withoutPrimaryKeys(pair.Old), New: withoutPrimaryKeys(pair.New)}
}

func TestAssessTablesListToJSONStoresNoSchema(t *testing.T) {
	require.Equal(t, plugin.TableFinding{
		TableName:          "datadog_monitors",
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
		Columns:            []plugin.ColumnFinding{noChangeColumn("tags")},
	}, assessOneTable(t, datadogTagsPair()))
}

func TestAssessTablesAddedPrimaryKeyColumn(t *testing.T) {
	t.Run("default write mode", func(t *testing.T) {
		require.Equal(t, plugin.TableFinding{
			TableName:          "okta_policy_rules",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRecreateIndex,
			Columns: []plugin.ColumnFinding{
				{ColumnName: "policy_id", Category: plugin.AssessCategoryManualMigrationRequired, SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateIndex},
				noChangeColumn("actions"),
				noChangeColumn("conditions"),
			},
		}, assessOneTable(t, oktaPolicyIDPair()))
	})

	t.Run("append write mode", func(t *testing.T) {
		require.Equal(t, plugin.TableFinding{
			TableName:          "okta_policy_rules",
			Category:           plugin.AssessCategoryNoChange,
			SafeModeBehavior:   behaviorNoChange,
			ForcedModeBehavior: behaviorNoChange,
			Columns:            []plugin.ColumnFinding{noChangeColumn("policy_id"), noChangeColumn("actions"), noChangeColumn("conditions")},
		}, assessOneTable(t, withoutPrimaryKeysPair(oktaPolicyIDPair())))
	})
}

func TestAssessTablesPrimaryKeyIndex(t *testing.T) {
	withPrimaryKey := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "name", Type: arrow.BinaryTypes.String},
	}}
	withCQIDPrimaryKey := withoutPrimaryKeys(withPrimaryKey)
	withCQIDPrimaryKey.Columns = append(withCQIDPrimaryKey.Columns, schema.Column{Name: schema.CqIDColumn.Name, Type: types.ExtensionTypes.UUID, PrimaryKey: true})

	tests := []struct {
		name     string
		pair     plugin.TablePair
		expected plugin.TableFinding
	}{
		{
			name: "primary key added",
			pair: plugin.TablePair{Old: withoutPrimaryKeys(withPrimaryKey), New: withPrimaryKey},
			expected: plugin.TableFinding{
				TableName: "test_table", Category: plugin.AssessCategoryAutomaticallyMigratable, SafeModeBehavior: behaviorCreateIndex, ForcedModeBehavior: behaviorCreateIndex,
				Columns: []plugin.ColumnFinding{{ColumnName: "id", Category: plugin.AssessCategoryAutomaticallyMigratable, SafeModeBehavior: behaviorCreateIndex, ForcedModeBehavior: behaviorCreateIndex}},
			},
		},
		{
			name: "primary key removed",
			pair: plugin.TablePair{Old: withPrimaryKey, New: withoutPrimaryKeys(withPrimaryKey)},
			expected: plugin.TableFinding{
				TableName: "test_table", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorKeepIndex, ForcedModeBehavior: behaviorKeepIndex,
				Columns: []plugin.ColumnFinding{{ColumnName: "id", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorKeepIndex, ForcedModeBehavior: behaviorKeepIndex}},
			},
		},
		{
			name: "primary key moved to _cq_id",
			pair: plugin.TablePair{Old: withPrimaryKey, New: withCQIDPrimaryKey},
			expected: plugin.TableFinding{
				TableName: "test_table", Category: plugin.AssessCategoryManualMigrationRequired, SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateIndex,
				Columns: []plugin.ColumnFinding{
					{ColumnName: "id", Category: plugin.AssessCategoryManualMigrationRequired, SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateIndex},
					{ColumnName: schema.CqIDColumn.Name, Category: plugin.AssessCategoryManualMigrationRequired, SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateIndex},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, assessOneTable(t, tc.pair))
		})
	}
}

func TestAssessTablesColumnChangesStoreNoSchema(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "count", Type: arrow.PrimitiveTypes.Int64},
		{Name: "removed", Type: arrow.BinaryTypes.String},
	}}
	newTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "count", Type: arrow.BinaryTypes.String},
		{Name: "added", Type: arrow.BinaryTypes.String},
	}}

	require.Equal(t, plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
		Columns:            []plugin.ColumnFinding{noChangeColumn("id"), noChangeColumn("count"), noChangeColumn("added"), noChangeColumn("removed")},
	}, assessOneTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
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
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryAutomaticallyMigratable, SafeModeBehavior: behaviorCreateIndex, ForcedModeBehavior: behaviorCreateIndex},
		},
		{
			name:     "added table without primary key",
			pair:     plugin.TablePair{New: withoutPrimaryKeys(table)},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryAutomaticallyMigratable, SafeModeBehavior: behaviorCreateOnWrite, ForcedModeBehavior: behaviorCreateOnWrite},
		},
		{
			name:     "removed table",
			pair:     plugin.TablePair{Old: table},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryTableRemoved, SafeModeBehavior: behaviorKeepCollection, ForcedModeBehavior: behaviorKeepCollection},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, assessOneTable(t, tc.pair))
		})
	}
}
