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

func assessTable(t *testing.T, pair plugin.TablePair) plugin.TableFinding {
	t.Helper()
	findings, err := newNoConnectionClient(t).AssessTables(context.Background(), []plugin.TablePair{pair}, plugin.AssessOptions{})
	require.NoError(t, err)
	require.Len(t, findings, 1)
	return findings[0]
}

func TestNewWithNoConnectionSkipsBigQueryClient(t *testing.T) {
	require.Nil(t, newNoConnectionClient(t).client)
}

func TestAssessTablesTypeChange(t *testing.T) {
	oldTable := &schema.Table{Name: "datadog_monitors", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "tags", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].Type = types.ExtensionTypes.JSON

	require.Equal(t, plugin.TableFinding{
		TableName:          "datadog_monitors",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorRejectChanges,
		ForcedModeBehavior: behaviorRejectChanges,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "tags",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			OldType:            "ARRAY<STRING>",
			NewType:            "JSON",
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRejectChanges,
		}},
	}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesAddedColumns(t *testing.T) {
	oldTable := &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = append(newTable.Columns,
		schema.Column{Name: "policy_id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		schema.Column{Name: "actions", Type: types.ExtensionTypes.JSON},
	)

	require.Equal(t, plugin.TableFinding{
		TableName:          "okta_policy_rules",
		Category:           plugin.AssessCategoryAutomaticallyMigratable,
		SafeModeBehavior:   behaviorMigrateTable,
		ForcedModeBehavior: behaviorMigrateTable,
		Columns: []plugin.ColumnFinding{
			{ColumnName: "policy_id", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "STRING", SafeModeBehavior: behaviorAddColumn, ForcedModeBehavior: behaviorAddColumn},
			{ColumnName: "actions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "JSON", SafeModeBehavior: behaviorAddColumn, ForcedModeBehavior: behaviorAddColumn},
		},
	}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesRemovedColumn(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "name", Type: arrow.BinaryTypes.String, NotNull: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = newTable.Columns[:1]

	finding := assessTable(t, plugin.TablePair{Old: oldTable, New: newTable})
	require.Equal(t, plugin.AssessCategoryAutomaticallyMigratable, finding.Category)
	require.Equal(t, []plugin.ColumnFinding{{
		ColumnName:         "name",
		Category:           plugin.AssessCategoryAutomaticallyMigratable,
		OldType:            "STRING",
		SafeModeBehavior:   behaviorKeepColumn,
		ForcedModeBehavior: behaviorKeepColumn,
	}}, finding.Columns)
}

func TestAssessTablesDefinitionChangeWithSameType(t *testing.T) {
	tests := []struct {
		name    string
		oldType arrow.DataType
		newType arrow.DataType
		oldName string
		newName string
	}{
		{
			name:    "mode",
			oldType: arrow.BinaryTypes.String,
			newType: arrow.ListOf(arrow.BinaryTypes.String),
			oldName: "STRING",
			newName: "ARRAY<STRING>",
		},
		{
			name:    "nested field",
			oldType: arrow.StructOf(arrow.Field{Name: "a", Type: arrow.BinaryTypes.String}),
			newType: arrow.StructOf(arrow.Field{Name: "a", Type: arrow.BinaryTypes.String}, arrow.Field{Name: "b", Type: arrow.PrimitiveTypes.Int64}),
			oldName: "RECORD<a STRING>",
			newName: "RECORD<a STRING, b INTEGER>",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{{Name: "value", Type: tc.oldType}}}
			newTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{{Name: "value", Type: tc.newType}}}

			require.Equal(t, plugin.TableFinding{
				TableName:          "test_table",
				Category:           plugin.AssessCategoryManualMigrationRequired,
				SafeModeBehavior:   behaviorMigrateTable,
				ForcedModeBehavior: behaviorMigrateTable,
				Columns: []plugin.ColumnFinding{{
					ColumnName:         "value",
					Category:           plugin.AssessCategoryManualMigrationRequired,
					OldType:            tc.oldName,
					NewType:            tc.newName,
					SafeModeBehavior:   behaviorKeepColumnDefinition,
					ForcedModeBehavior: behaviorKeepColumnDefinition,
				}},
			}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
		})
	}
}

func TestAssessTablesWholeTable(t *testing.T) {
	table := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	withoutPrimaryKey := table.Copy(nil)
	withoutPrimaryKey.Columns[0].PrimaryKey = false

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
			name:     "primary key change only",
			pair:     plugin.TablePair{Old: table, New: withoutPrimaryKey},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorNoChange, ForcedModeBehavior: behaviorNoChange},
		},
		{
			name:     "added table",
			pair:     plugin.TablePair{New: table},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryAutomaticallyMigratable, SafeModeBehavior: behaviorCreateTable, ForcedModeBehavior: behaviorCreateTable},
		},
		{
			name:     "removed table",
			pair:     plugin.TablePair{Old: table},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryTableRemoved, SafeModeBehavior: behaviorKeepTable, ForcedModeBehavior: behaviorKeepTable},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, assessTable(t, tc.pair))
		})
	}
}
