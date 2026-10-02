package client

import (
	"context"
	"encoding/json"
	"path/filepath"
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

func newFileClient(t *testing.T) *Client {
	t.Helper()
	spec, err := json.Marshal(Spec{ConnectionString: filepath.Join(t.TempDir(), "test.db")})
	require.NoError(t, err)
	c, err := New(context.Background(), zerolog.Nop(), spec, plugin.NewClientOptions{})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, c.Close(context.Background())) })
	return c.(*Client)
}

func assessTable(t *testing.T, pair plugin.TablePair) plugin.TableFinding {
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

func datadogMonitorsTags() plugin.TablePair {
	oldTable := &schema.Table{Name: "datadog_monitors", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "tags", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].Type = types.ExtensionTypes.JSON
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func stringToIntegerColumn() plugin.TablePair {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "count", Type: arrow.BinaryTypes.String},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].Type = arrow.PrimitiveTypes.Int64
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func oktaPolicyRules() plugin.TablePair {
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

func oktaPolicyRulesAppendMode() plugin.TablePair {
	pair := oktaPolicyRules()
	return plugin.TablePair{Old: withoutPrimaryKeys(pair.Old), New: withoutPrimaryKeys(pair.New)}
}

func removedColumn() plugin.TablePair {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "name", Type: arrow.BinaryTypes.String},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = newTable.Columns[:1]
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func TestNewWithNoConnectionSkipsDB(t *testing.T) {
	require.Nil(t, newNoConnectionClient(t).db)
}

func TestAssessTablesTypeChangeWithSameSqliteType(t *testing.T) {
	require.Equal(t, plugin.TableFinding{
		TableName:          "datadog_monitors",
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "tags",
			Category:           plugin.AssessCategoryNoChange,
			OldType:            "text",
			NewType:            "text",
			SafeModeBehavior:   behaviorNoChange,
			ForcedModeBehavior: behaviorNoChange,
		}},
	}, assessTable(t, datadogMonitorsTags()))
}

func TestAssessTablesTypeChange(t *testing.T) {
	require.Equal(t, plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorRejectChanges,
		ForcedModeBehavior: behaviorRecreateTable,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "count",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			OldType:            "text",
			NewType:            "integer",
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRecreateTable,
		}},
	}, assessTable(t, stringToIntegerColumn()))
}

func TestAssessTablesAddedPrimaryKeyColumn(t *testing.T) {
	t.Run("default write mode", func(t *testing.T) {
		require.Equal(t, plugin.TableFinding{
			TableName:          "okta_policy_rules",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRecreateTable,
			Columns: []plugin.ColumnFinding{
				{ColumnName: "policy_id", Category: plugin.AssessCategoryManualMigrationRequired, NewType: "text", SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateTable},
				{ColumnName: "actions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "text", SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateTable},
				{ColumnName: "conditions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "text", SafeModeBehavior: behaviorRejectChanges, ForcedModeBehavior: behaviorRecreateTable},
			},
		}, assessTable(t, oktaPolicyRules()))
	})

	t.Run("append write mode", func(t *testing.T) {
		require.Equal(t, plugin.TableFinding{
			TableName:          "okta_policy_rules",
			Category:           plugin.AssessCategoryAutomaticallyMigratable,
			SafeModeBehavior:   behaviorMigrateTable,
			ForcedModeBehavior: behaviorMigrateTable,
			Columns: []plugin.ColumnFinding{
				{ColumnName: "policy_id", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "text", SafeModeBehavior: "adds the column", ForcedModeBehavior: "adds the column"},
				{ColumnName: "actions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "text", SafeModeBehavior: "adds the column", ForcedModeBehavior: "adds the column"},
				{ColumnName: "conditions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "text", SafeModeBehavior: "adds the column", ForcedModeBehavior: "adds the column"},
			},
		}, assessTable(t, oktaPolicyRulesAppendMode()))
	})
}

func TestAssessTablesRemovedColumn(t *testing.T) {
	finding := assessTable(t, removedColumn())
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

func TestAssessTablesMatchesMigrate(t *testing.T) {
	tests := map[string]plugin.TablePair{
		"type change with same sqlite type": datadogMonitorsTags(),
		"type change":                       stringToIntegerColumn(),
		"added primary key column":          oktaPolicyRules(),
		"added nullable columns":            oktaPolicyRulesAppendMode(),
		"removed column":                    removedColumn(),
	}
	for name, pair := range tests {
		t.Run(name, func(t *testing.T) {
			manualMigration := assessTable(t, pair).Category == plugin.AssessCategoryManualMigrationRequired
			c := newFileClient(t)
			require.NoError(t, migrateTable(c, pair.Old, false))
			_, err := c.db.Exec(`insert into "` + pair.Old.Name + `" ("id") values ('1')`)
			require.NoError(t, err)

			require.Equal(t, manualMigration, migrateTable(c, pair.New, false) != nil, "safe migration rejects the changes")
			require.NoError(t, migrateTable(c, pair.New, true))
			var rows int
			require.NoError(t, c.db.QueryRow(`select count(*) from "`+pair.New.Name+`"`).Scan(&rows))
			require.Equal(t, manualMigration, rows == 0, "forced migration recreates the table")
		})
	}
}

func migrateTable(c *Client, table *schema.Table, force bool) error {
	return c.MigrateTables(context.Background(), message.WriteMigrateTables{{Table: table, MigrateForce: force}})
}
