package client

import (
	"context"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/meilisearch/v2/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func assessTables(t *testing.T, pairs ...plugin.TablePair) []plugin.TableFinding {
	t.Helper()
	ctx := context.Background()
	p := plugin.NewPlugin("meilisearch", internalPlugin.Version, New)
	require.NoError(t, p.Init(ctx, nil, plugin.NewClientOptions{NoConnection: true}))
	findings, err := p.AssessTables(ctx, pairs, plugin.AssessOptions{})
	require.NoError(t, err)
	require.Len(t, findings, len(pairs))
	return findings
}

func assessPair(t *testing.T, pair plugin.TablePair) plugin.TableFinding {
	t.Helper()
	return assessTables(t, pair)[0]
}

func withoutPrimaryKeys(table *schema.Table) *schema.Table {
	result := table.Copy(nil)
	for i := range result.Columns {
		result.Columns[i].PrimaryKey = false
	}
	return result
}

func TestNewWithNoConnectionSkipsConnection(t *testing.T) {
	c, err := New(context.Background(), zerolog.Nop(), nil, plugin.NewClientOptions{NoConnection: true})
	require.NoError(t, err)
	require.Nil(t, c.(*Client).Meilisearch)
	require.Error(t, c.Close(context.Background()))
}

func TestAssessTablesListToJSONHasNoFixedType(t *testing.T) {
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
			SafeModeBehavior:   behaviorUntypedAttribute,
			ForcedModeBehavior: behaviorUntypedAttribute,
		}},
	}, assessPair(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesAddedColumnKeepsIndexSettings(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = append(newTable.Columns, schema.Column{Name: "name", Type: arrow.BinaryTypes.String, NotNull: true})

	require.Equal(t, plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "name",
			Category:           plugin.AssessCategoryNoChange,
			SafeModeBehavior:   behaviorAddAttribute,
			ForcedModeBehavior: behaviorAddAttribute,
		}},
	}, assessPair(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesRemovedColumn(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "name", Type: arrow.BinaryTypes.String},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = newTable.Columns[:1]

	require.Equal(t, plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "name",
			Category:           plugin.AssessCategoryNoChange,
			SafeModeBehavior:   behaviorRemoveAttribute,
			ForcedModeBehavior: behaviorRemoveAttribute,
		}},
	}, assessPair(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesAddedPrimaryKeyColumn(t *testing.T) {
	oldTable := &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = append(newTable.Columns,
		schema.Column{Name: "policy_id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		schema.Column{Name: "actions", Type: types.ExtensionTypes.JSON},
	)

	t.Run("default write mode", func(t *testing.T) {
		require.Equal(t, plugin.TableFinding{
			TableName:          "okta_policy_rules",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			SafeModeBehavior:   behaviorNewDocumentIDs,
			ForcedModeBehavior: behaviorNewDocumentIDs,
			Columns: []plugin.ColumnFinding{
				{ColumnName: "policy_id", Category: plugin.AssessCategoryManualMigrationRequired, SafeModeBehavior: behaviorNewDocumentIDs, ForcedModeBehavior: behaviorNewDocumentIDs},
				{ColumnName: "actions", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorAddAttribute, ForcedModeBehavior: behaviorAddAttribute},
			},
		}, assessPair(t, plugin.TablePair{Old: oldTable, New: newTable}))
	})

	t.Run("append write mode", func(t *testing.T) {
		require.Equal(t, plugin.TableFinding{
			TableName:          "okta_policy_rules",
			Category:           plugin.AssessCategoryNoChange,
			SafeModeBehavior:   behaviorNoChange,
			ForcedModeBehavior: behaviorNoChange,
			Columns: []plugin.ColumnFinding{
				{ColumnName: "policy_id", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorAddAttribute, ForcedModeBehavior: behaviorAddAttribute},
				{ColumnName: "actions", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorAddAttribute, ForcedModeBehavior: behaviorAddAttribute},
			},
		}, assessPair(t, plugin.TablePair{Old: withoutPrimaryKeys(oldTable), New: withoutPrimaryKeys(newTable)}))
	})
}

func TestAssessTablesPrimaryKeyOrderChange(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "a", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "b", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	newTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{oldTable.Columns[1], oldTable.Columns[0]}}

	require.Equal(t, plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorNewDocumentIDs,
		ForcedModeBehavior: behaviorNewDocumentIDs,
	}, assessPair(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesWholeTable(t *testing.T) {
	table := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}

	require.Equal(t, []plugin.TableFinding{
		{TableName: "test_table", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorNoChange, ForcedModeBehavior: behaviorNoChange},
		{TableName: "test_table", Category: plugin.AssessCategoryAutomaticallyMigratable, SafeModeBehavior: behaviorCreateIndex, ForcedModeBehavior: behaviorCreateIndex},
		{TableName: "test_table", Category: plugin.AssessCategoryTableRemoved, SafeModeBehavior: behaviorKeepIndex, ForcedModeBehavior: behaviorKeepIndex},
	}, assessTables(t,
		plugin.TablePair{Old: table, New: table},
		plugin.TablePair{New: table},
		plugin.TablePair{Old: table},
	))
}

func TestPlanIndexMigration(t *testing.T) {
	need := &indexSchema{UID: "test_table", PrimaryKey: hashColumnName}

	require.Equal(t, indexCreate, planIndexMigration(nil, need))
	require.Equal(t, indexUpdate, planIndexMigration(&indexSchema{UID: "test_table", PrimaryKey: hashColumnName}, need))
	require.Equal(t, indexRecreate, planIndexMigration(&indexSchema{UID: "test_table", PrimaryKey: "id"}, need))
}
