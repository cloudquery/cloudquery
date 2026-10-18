package client

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func newNoConnectionClient(t *testing.T, spec []byte) *Client {
	t.Helper()
	c, err := New(context.Background(), zerolog.Nop(), spec, plugin.NewClientOptions{NoConnection: true})
	require.NoError(t, err)
	return c.(*Client)
}

func assessTablePair(t *testing.T, pair plugin.TablePair) plugin.TableFinding {
	t.Helper()
	findings, err := newNoConnectionClient(t, nil).AssessTables(context.Background(), []plugin.TablePair{pair}, plugin.AssessOptions{})
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

func TestNewWithNoConnectionOpensNoDatabaseFile(t *testing.T) {
	databaseFile := filepath.Join(t.TempDir(), "test.duckdb")
	spec, err := json.Marshal(Spec{ConnectionString: databaseFile})
	require.NoError(t, err)

	c := newNoConnectionClient(t, spec)

	require.Nil(t, c.connector)
	require.Nil(t, c.db)
	require.NoFileExists(t, databaseFile)
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
		ForcedModeBehavior: behaviorRecreateTable,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "tags",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			OldType:            "varchar[]",
			NewType:            "json",
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRecreateTable,
		}},
	}, assessTablePair(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesAddedPrimaryKeyColumn(t *testing.T) {
	oldTable := &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = append(newTable.Columns,
		schema.Column{Name: "policy_id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		schema.Column{Name: "actions", Type: types.ExtensionTypes.JSON},
		schema.Column{Name: "conditions", Type: types.ExtensionTypes.JSON},
	)

	t.Run("default write mode", func(t *testing.T) {
		require.Equal(t, plugin.TableFinding{
			TableName:          "okta_policy_rules",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRecreateTable,
			Columns: []plugin.ColumnFinding{
				{ColumnName: "policy_id", Category: plugin.AssessCategoryManualMigrationRequired, NewType: "varchar", SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateTable},
				{ColumnName: "actions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "json", SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateTable},
				{ColumnName: "conditions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "json", SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateTable},
			},
		}, assessTablePair(t, plugin.TablePair{Old: oldTable, New: newTable}))
	})

	t.Run("append write mode", func(t *testing.T) {
		require.Equal(t, plugin.TableFinding{
			TableName:          "okta_policy_rules",
			Category:           plugin.AssessCategoryAutomaticallyMigratable,
			SafeModeBehavior:   behaviorMigrateTable,
			ForcedModeBehavior: behaviorMigrateTable,
			Columns: []plugin.ColumnFinding{
				{ColumnName: "policy_id", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "varchar", SafeModeBehavior: "adds the column", ForcedModeBehavior: "adds the column"},
				{ColumnName: "actions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "json", SafeModeBehavior: "adds the column", ForcedModeBehavior: "adds the column"},
				{ColumnName: "conditions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "json", SafeModeBehavior: "adds the column", ForcedModeBehavior: "adds the column"},
			},
		}, assessTablePair(t, plugin.TablePair{Old: withoutPrimaryKeys(oldTable), New: withoutPrimaryKeys(newTable)}))
	})
}

func TestAssessTablesRemovedColumn(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "name", Type: arrow.BinaryTypes.String},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = newTable.Columns[:1]

	finding := assessTablePair(t, plugin.TablePair{Old: oldTable, New: newTable})
	require.Equal(t, plugin.AssessCategoryAutomaticallyMigratable, finding.Category)
	require.Equal(t, []plugin.ColumnFinding{{
		ColumnName:         "name",
		Category:           plugin.AssessCategoryAutomaticallyMigratable,
		OldType:            "varchar",
		SafeModeBehavior:   "keeps the column, new rows leave it empty",
		ForcedModeBehavior: "keeps the column, new rows leave it empty",
	}}, finding.Columns)
}

func TestAssessTablesRemovedUniqueConstraint(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, Unique: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[0].Unique = false

	require.Equal(t, plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorRejectChanges,
		ForcedModeBehavior: behaviorRecreateTable,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "id",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			OldType:            "varchar",
			NewType:            "varchar",
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRecreateTable,
		}},
	}, assessTablePair(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesWholeTable(t *testing.T) {
	table := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "count", Type: arrow.PrimitiveTypes.Uint8},
	}}
	sameDuckDBType := table.Copy(nil)
	sameDuckDBType.Columns[1].Type = arrow.PrimitiveTypes.Uint16

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
			name:     "type change to the same duckdb type",
			pair:     plugin.TablePair{Old: table, New: sameDuckDBType},
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
			require.Equal(t, tc.expected, assessTablePair(t, tc.pair))
		})
	}
}
