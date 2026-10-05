package client

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v8/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func newNoConnectionClient(t *testing.T) *Client {
	t.Helper()
	c, err := New(context.Background(), zerolog.Nop(), nil, plugin.NewClientOptions{NoConnection: true})
	require.NoError(t, err)
	return c.(*Client)
}

func newConnectedClient(t *testing.T) *Client {
	t.Helper()
	specBytes, err := json.Marshal(spec.Spec{ConnectionString: getTestConnection()})
	require.NoError(t, err)
	c, err := New(context.Background(), zerolog.Nop(), specBytes, plugin.NewClientOptions{})
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

func addedNotNullColumn() plugin.TablePair {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, NotNull: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = append(newTable.Columns, schema.Column{Name: "name", Type: arrow.BinaryTypes.String, NotNull: true})
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func removedColumn(notNull bool) plugin.TablePair {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "name", Type: arrow.BinaryTypes.String, NotNull: notNull},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = newTable.Columns[:1]
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func TestNewWithNoConnectionSkipsConnection(t *testing.T) {
	c := newNoConnectionClient(t)
	require.Nil(t, c.conn)
	require.NoError(t, c.Close(context.Background()))
}

func TestAssessTablesTypeChange(t *testing.T) {
	require.Equal(t, plugin.TableFinding{
		TableName:          "datadog_monitors",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorRejectChanges,
		ForcedModeBehavior: behaviorRecreateTable,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "tags",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			OldType:            "Array(Nullable(String))",
			NewType:            "Nullable(String)",
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRecreateTable,
		}},
	}, assessTable(t, datadogMonitorsTags()))
}

func TestAssessTablesAddedNullableColumns(t *testing.T) {
	expected := plugin.TableFinding{
		TableName:          "okta_policy_rules",
		Category:           plugin.AssessCategoryAutomaticallyMigratable,
		SafeModeBehavior:   behaviorMigrateTable,
		ForcedModeBehavior: behaviorMigrateTable,
		Columns: []plugin.ColumnFinding{
			{ColumnName: "policy_id", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "Nullable(String)", SafeModeBehavior: "adds the column", ForcedModeBehavior: "adds the column"},
			{ColumnName: "actions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "Nullable(String)", SafeModeBehavior: "adds the column", ForcedModeBehavior: "adds the column"},
			{ColumnName: "conditions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "Nullable(String)", SafeModeBehavior: "adds the column", ForcedModeBehavior: "adds the column"},
		},
	}

	t.Run("default write mode", func(t *testing.T) {
		require.Equal(t, expected, assessTable(t, oktaPolicyRules()))
	})

	t.Run("append write mode", func(t *testing.T) {
		require.Equal(t, expected, assessTable(t, oktaPolicyRulesAppendMode()))
	})
}

func TestAssessTablesAddedNotNullColumn(t *testing.T) {
	require.Equal(t, plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorRejectChanges,
		ForcedModeBehavior: behaviorRecreateTable,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "name",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			NewType:            "String",
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRecreateTable,
		}},
	}, assessTable(t, addedNotNullColumn()))
}

func TestAssessTablesRemovedColumn(t *testing.T) {
	t.Run("nullable", func(t *testing.T) {
		finding := assessTable(t, removedColumn(false))
		require.Equal(t, plugin.AssessCategoryAutomaticallyMigratable, finding.Category)
		require.Equal(t, []plugin.ColumnFinding{{
			ColumnName:         "name",
			Category:           plugin.AssessCategoryAutomaticallyMigratable,
			OldType:            "Nullable(String)",
			SafeModeBehavior:   "keeps the column, new rows leave it empty",
			ForcedModeBehavior: "keeps the column, new rows leave it empty",
		}}, finding.Columns)
	})

	t.Run("not null", func(t *testing.T) {
		finding := assessTable(t, removedColumn(true))
		require.Equal(t, plugin.AssessCategoryManualMigrationRequired, finding.Category)
		require.Equal(t, []plugin.ColumnFinding{{
			ColumnName:         "name",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			OldType:            "String",
			SafeModeBehavior:   behaviorRejectChanges,
			ForcedModeBehavior: behaviorRecreateTable,
		}}, finding.Columns)
	})
}

func TestAssessTablesSortingKeyChange(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "a", Type: arrow.BinaryTypes.String, NotNull: true},
		{Name: "b", Type: arrow.BinaryTypes.String, NotNull: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[0], newTable.Columns[1] = newTable.Columns[1], newTable.Columns[0]

	pair := plugin.TablePair{Old: oldTable, New: newTable}

	t.Run("default order", func(t *testing.T) {
		require.Equal(t,
			tableFinding("test_table", plugin.AssessCategoryManualMigrationRequired, behaviorRejectChanges, behaviorRecreateTable),
			assessTable(t, pair),
		)
	})

	t.Run("order from spec", func(t *testing.T) {
		specBytes, err := json.Marshal(spec.Spec{OrderBy: []spec.OrderByStrategy{{Tables: []string{"*"}, OrderBy: []string{"a"}}}})
		require.NoError(t, err)
		c, err := New(context.Background(), zerolog.Nop(), specBytes, plugin.NewClientOptions{NoConnection: true})
		require.NoError(t, err)
		findings, err := c.(*Client).AssessTables(context.Background(), []plugin.TablePair{pair}, plugin.AssessOptions{})
		require.NoError(t, err)
		require.Equal(t, []plugin.TableFinding{
			tableFinding("test_table", plugin.AssessCategoryNoChange, behaviorNoChange, behaviorNoChange),
		}, findings)
	})
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
			expected: tableFinding("test_table", plugin.AssessCategoryNoChange, behaviorNoChange, behaviorNoChange),
		},
		{
			name:     "added table",
			pair:     plugin.TablePair{New: table},
			expected: tableFinding("test_table", plugin.AssessCategoryAutomaticallyMigratable, behaviorCreateTable, behaviorCreateTable),
		},
		{
			name:     "removed table",
			pair:     plugin.TablePair{Old: table},
			expected: tableFinding("test_table", plugin.AssessCategoryTableRemoved, behaviorKeepTable, behaviorKeepTable),
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
		"type change":              datadogMonitorsTags(),
		"added primary key column": oktaPolicyRules(),
		"added nullable columns":   oktaPolicyRulesAppendMode(),
		"added not null column":    addedNotNullColumn(),
		"removed nullable column":  removedColumn(false),
		"removed not null column":  removedColumn(true),
	}
	for name, pair := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			manualMigration := assessTable(t, pair).Category == plugin.AssessCategoryManualMigrationRequired
			pair = withUniqueTableName(pair)
			c := newConnectedClient(t)
			t.Cleanup(func() { require.NoError(t, c.dropTable(ctx, pair.New)) })

			require.NoError(t, migrateTable(c, pair.Old, false))
			require.NoError(t, c.conn.Exec(ctx, "INSERT INTO `"+pair.Old.Name+"` (`id`) VALUES ('1')"))

			require.Equal(t, manualMigration, migrateTable(c, pair.New, false) != nil, "safe migration rejects the changes")
			require.NoError(t, migrateTable(c, pair.New, true))
			var rows uint64
			require.NoError(t, c.conn.QueryRow(ctx, "SELECT count() FROM `"+pair.New.Name+"`").Scan(&rows))
			require.Equal(t, manualMigration, rows == 0, "forced migration recreates the table")
		})
	}
}

func withUniqueTableName(pair plugin.TablePair) plugin.TablePair {
	name := fmt.Sprintf("%s_%d", pair.Old.Name, time.Now().UnixNano())
	oldTable, newTable := pair.Old.Copy(nil), pair.New.Copy(nil)
	oldTable.Name, newTable.Name = name, name
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func migrateTable(c *Client, table *schema.Table, force bool) error {
	return c.MigrateTables(context.Background(), message.WriteMigrateTables{{Table: table, MigrateForce: force}})
}
