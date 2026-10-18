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

func TestNewWithNoConnectionSkipsDB(t *testing.T) {
	c := newNoConnectionClient(t)
	require.Nil(t, c.db)
	require.Error(t, c.Close(context.Background()))
}

func TestAssessTablesListToJSONIsStoredUnchanged(t *testing.T) {
	oldTable := &schema.Table{Name: "datadog_monitors", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "tags", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].Type = types.ExtensionTypes.JSON

	require.Equal(t, plugin.TableFinding{
		TableName:          "datadog_monitors",
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "tags",
			Category:           plugin.AssessCategoryNoChange,
			OldType:            "variant",
			NewType:            "variant",
			SafeModeBehavior:   behaviorNoChange,
			ForcedModeBehavior: behaviorNoChange,
		}},
	}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesTypeChange(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "count", Type: arrow.PrimitiveTypes.Int64},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].Type = arrow.BinaryTypes.String

	require.Equal(t, plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorRejectChanges,
		ForcedModeBehavior: behaviorRecreateTable,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "count",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			OldType:            "number",
			NewType:            "text",
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRecreateTable,
		}},
	}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesAddedColumn(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}

	t.Run("nullable", func(t *testing.T) {
		newTable := oldTable.Copy(nil)
		newTable.Columns = append(newTable.Columns, schema.Column{Name: "name", Type: arrow.BinaryTypes.String})

		require.Equal(t, plugin.TableFinding{
			TableName:          "test_table",
			Category:           plugin.AssessCategoryAutomaticallyMigratable,
			SafeModeBehavior:   behaviorMigrateTable,
			ForcedModeBehavior: behaviorMigrateTable,
			Columns: []plugin.ColumnFinding{{
				ColumnName:         "name",
				Category:           plugin.AssessCategoryAutomaticallyMigratable,
				NewType:            "text",
				SafeModeBehavior:   "adds the column",
				ForcedModeBehavior: "adds the column",
			}},
		}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
	})

	t.Run("not null", func(t *testing.T) {
		newTable := oldTable.Copy(nil)
		newTable.Columns = append(newTable.Columns, schema.Column{Name: "name", Type: arrow.BinaryTypes.String, NotNull: true})

		require.Equal(t, plugin.TableFinding{
			TableName:          "test_table",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRecreateTable,
			Columns: []plugin.ColumnFinding{{
				ColumnName:         "name",
				Category:           plugin.AssessCategoryManualMigrationRequired,
				NewType:            "text",
				SafeModeBehavior:   behaviorRejectChanges,
				ForcedModeBehavior: behaviorRecreateTable,
			}},
		}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
	})
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

	require.Equal(t, plugin.TableFinding{
		TableName:          "okta_policy_rules",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorMigrateFails,
		ForcedModeBehavior: behaviorMigrateFails,
		Columns: []plugin.ColumnFinding{
			{ColumnName: "policy_id", Category: plugin.AssessCategoryManualMigrationRequired, NewType: "text", SafeModeBehavior: behaviorMigrateFails, ForcedModeBehavior: behaviorMigrateFails},
			{ColumnName: "actions", Category: plugin.AssessCategoryManualMigrationRequired, NewType: "variant", SafeModeBehavior: behaviorMigrateFails, ForcedModeBehavior: behaviorMigrateFails},
			{ColumnName: "conditions", Category: plugin.AssessCategoryManualMigrationRequired, NewType: "variant", SafeModeBehavior: behaviorMigrateFails, ForcedModeBehavior: behaviorMigrateFails},
		},
	}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesPrimaryKeyConstraintChanges(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true, NotNull: true},
		{Name: "region", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}

	t.Run("primary key and not null dropped", func(t *testing.T) {
		newTable := oldTable.Copy(nil)
		newTable.Columns[0].PrimaryKey = false
		newTable.Columns[0].NotNull = false

		finding := assessTable(t, plugin.TablePair{Old: oldTable, New: newTable})
		require.Equal(t, plugin.AssessCategoryAutomaticallyMigratable, finding.Category)
		require.Equal(t, "updates the primary key, drops the not null constraint", finding.Columns[0].SafeModeBehavior)
	})

	t.Run("primary key column removed", func(t *testing.T) {
		newTable := oldTable.Copy(nil)
		newTable.Columns = newTable.Columns[:1]

		finding := assessTable(t, plugin.TablePair{Old: oldTable, New: newTable})
		require.Equal(t, plugin.AssessCategoryAutomaticallyMigratable, finding.Category)
		require.Equal(t, "keeps the column, new rows leave it empty, drops the primary key", finding.Columns[0].SafeModeBehavior)
	})
}

func TestAssessTablesMixedChangesRecreateTable(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "count", Type: arrow.PrimitiveTypes.Int64},
		{Name: "tags", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].Type = arrow.BinaryTypes.String
	newTable.Columns[2].Type = types.ExtensionTypes.JSON
	newTable.Columns = append(newTable.Columns, schema.Column{Name: "name", Type: arrow.BinaryTypes.String})

	require.Equal(t, plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorRejectChanges,
		ForcedModeBehavior: behaviorRecreateTable,
		Columns: []plugin.ColumnFinding{
			{ColumnName: "count", Category: plugin.AssessCategoryManualMigrationRequired, OldType: "number", NewType: "text", SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateTable},
			{ColumnName: "name", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "text", SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateTable},
			{ColumnName: "tags", Category: plugin.AssessCategoryNoChange, OldType: "variant", NewType: "variant", SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateTable},
		},
	}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesDroppedNotNull(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "name", Type: arrow.BinaryTypes.String, NotNull: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].NotNull = false

	require.Equal(t, plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryAutomaticallyMigratable,
		SafeModeBehavior:   behaviorMigrateTable,
		ForcedModeBehavior: behaviorMigrateTable,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "name",
			Category:           plugin.AssessCategoryAutomaticallyMigratable,
			OldType:            "text",
			NewType:            "text",
			SafeModeBehavior:   "drops the not null constraint",
			ForcedModeBehavior: "drops the not null constraint",
		}},
	}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesRemovedColumn(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "name", Type: arrow.BinaryTypes.String},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = newTable.Columns[:1]

	finding := assessTable(t, plugin.TablePair{Old: oldTable, New: newTable})
	require.Equal(t, plugin.AssessCategoryAutomaticallyMigratable, finding.Category)
	require.Equal(t, []plugin.ColumnFinding{{
		ColumnName:         "name",
		Category:           plugin.AssessCategoryAutomaticallyMigratable,
		OldType:            "text",
		SafeModeBehavior:   "keeps the column, new rows leave it empty",
		ForcedModeBehavior: "keeps the column, new rows leave it empty",
	}}, finding.Columns)
}

func TestAssessTablesWholeTable(t *testing.T) {
	table := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	upperCaseTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "ID", Type: arrow.BinaryTypes.String, PrimaryKey: true},
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
			name:     "column name case change",
			pair:     plugin.TablePair{Old: upperCaseTable, New: table},
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
