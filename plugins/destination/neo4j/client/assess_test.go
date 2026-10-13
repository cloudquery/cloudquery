package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func newNoConnectionClient(t *testing.T) *Client {
	t.Helper()
	c, err := New(context.Background(), zerolog.Nop(), nil, plugin.NewClientOptions{NoConnection: true})
	require.NoError(t, err)
	return c.(*Client)
}

func assessPair(t *testing.T, pair plugin.TablePair) plugin.TableFinding {
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

func datadogMonitorsTagsChange() plugin.TablePair {
	oldTable := &schema.Table{Name: "datadog_monitors", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "tags", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].Type = types.ExtensionTypes.JSON
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func oktaPolicyRulesPrimaryKeyAdded() plugin.TablePair {
	oldTable := &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns = append(newTable.Columns,
		schema.Column{Name: "policy_id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		schema.Column{Name: "actions", Type: types.ExtensionTypes.JSON},
	)
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func TestNewWithNoConnectionSkipsDriver(t *testing.T) {
	require.Nil(t, newNoConnectionClient(t).client)
}

func TestAssessTablesTypeChange(t *testing.T) {
	require.Equal(t, plugin.TableFinding{
		TableName:          "datadog_monitors",
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "tags",
			Category:           plugin.AssessCategoryNoChange,
			SafeModeBehavior:   behaviorNoFixedSchema,
			ForcedModeBehavior: behaviorNoFixedSchema,
		}},
	}, assessPair(t, datadogMonitorsTagsChange()))
}

func TestAssessTablesAddedPrimaryKeyColumn(t *testing.T) {
	pair := oktaPolicyRulesPrimaryKeyAdded()

	t.Run("default write mode", func(t *testing.T) {
		require.Equal(t, plugin.TableFinding{
			TableName:          "okta_policy_rules",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			SafeModeBehavior:   behaviorKeepIndex,
			ForcedModeBehavior: behaviorKeepIndex,
			Columns: []plugin.ColumnFinding{
				{ColumnName: "policy_id", Category: plugin.AssessCategoryManualMigrationRequired, SafeModeBehavior: behaviorKeepIndex, ForcedModeBehavior: behaviorKeepIndex},
				{ColumnName: "actions", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorNoFixedSchema, ForcedModeBehavior: behaviorNoFixedSchema},
			},
		}, assessPair(t, pair))
	})

	t.Run("append write mode", func(t *testing.T) {
		require.Equal(t, plugin.TableFinding{
			TableName:          "okta_policy_rules",
			Category:           plugin.AssessCategoryNoChange,
			SafeModeBehavior:   behaviorNoChange,
			ForcedModeBehavior: behaviorNoChange,
			Columns: []plugin.ColumnFinding{
				{ColumnName: "policy_id", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorNoFixedSchema, ForcedModeBehavior: behaviorNoFixedSchema},
				{ColumnName: "actions", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorNoFixedSchema, ForcedModeBehavior: behaviorNoFixedSchema},
			},
		}, assessPair(t, plugin.TablePair{Old: withoutPrimaryKeys(pair.Old), New: withoutPrimaryKeys(pair.New)}))
	})
}

func TestAssessTablesPrimaryKeyRemoved(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}

	finding := assessPair(t, plugin.TablePair{Old: oldTable, New: withoutPrimaryKeys(oldTable)})
	require.Equal(t, plugin.AssessCategoryManualMigrationRequired, finding.Category)
	require.Equal(t, []plugin.ColumnFinding{{
		ColumnName:         "id",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorKeepIndexNoPK,
		ForcedModeBehavior: behaviorKeepIndexNoPK,
	}}, finding.Columns)
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
			SafeModeBehavior:   behaviorNoFixedSchema,
			ForcedModeBehavior: behaviorNoFixedSchema,
		}},
	}, assessPair(t, plugin.TablePair{Old: oldTable, New: newTable}))
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
			name:     "added table with primary key",
			pair:     plugin.TablePair{New: table},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryAutomaticallyMigratable, SafeModeBehavior: behaviorCreateIndex, ForcedModeBehavior: behaviorCreateIndex},
		},
		{
			name:     "added table without primary key",
			pair:     plugin.TablePair{New: withoutPrimaryKeys(table)},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorNoChange, ForcedModeBehavior: behaviorNoChange},
		},
		{
			name:     "removed table",
			pair:     plugin.TablePair{Old: table},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryTableRemoved, SafeModeBehavior: behaviorKeepNodes, ForcedModeBehavior: behaviorKeepNodes},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, assessPair(t, tc.pair))
		})
	}
}

func newConnectedClient(t *testing.T) *Client {
	t.Helper()
	s := &Spec{
		Username:         getenv("CQ_DEST_NEO4J_USERNAME", defaultUsername),
		Password:         getenv("CQ_DEST_NEO4J_PASSWORD", defaultPassword),
		ConnectionString: getenv("CQ_DEST_NEO4J_CONNECTION_STRING", defaultConnectionString),
	}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	c, err := New(context.Background(), zerolog.Nop(), b, plugin.NewClientOptions{})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, c.Close(context.Background())) })
	return c.(*Client)
}

func indexedProperties(t *testing.T, c *Client, table *schema.Table) []string {
	t.Helper()
	ctx := context.Background()
	sess := c.Session(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer sess.Close(ctx)

	result, err := sess.Run(ctx, `SHOW INDEXES YIELD name, properties WHERE name = $name RETURN properties`, map[string]any{"name": indexName(table)})
	require.NoError(t, err)
	records, err := result.Collect(ctx)
	require.NoError(t, err)
	if len(records) == 0 {
		return nil
	}
	properties, _ := records[0].Get("properties")
	var names []string
	for _, p := range properties.([]any) {
		names = append(names, p.(string))
	}
	return names
}

func dropIndex(t *testing.T, c *Client, table *schema.Table) {
	t.Helper()
	ctx := context.Background()
	sess := c.Session(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer sess.Close(ctx)
	_, err := sess.Run(ctx, `DROP INDEX `+indexName(table)+` IF EXISTS;`, map[string]any{})
	require.NoError(t, err)
}

func renamed(pair plugin.TablePair, name string) plugin.TablePair {
	oldTable, newTable := pair.Old.Copy(nil), pair.New.Copy(nil)
	oldTable.Name, newTable.Name = name, name
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func TestAssessTablesMatchesMigrate(t *testing.T) {
	ctx := context.Background()
	c := newConnectedClient(t)
	withoutPK := oktaPolicyRulesPrimaryKeyAdded()
	withoutPK.Old = withoutPrimaryKeys(withoutPK.Old)

	tests := []struct {
		name     string
		pair     plugin.TablePair
		category plugin.AssessCategory
	}{
		{name: "datadog tags type change", pair: datadogMonitorsTagsChange(), category: plugin.AssessCategoryNoChange},
		{name: "okta primary key added", pair: oktaPolicyRulesPrimaryKeyAdded(), category: plugin.AssessCategoryManualMigrationRequired},
		{name: "primary key added to table without one", pair: withoutPK, category: plugin.AssessCategoryAutomaticallyMigratable},
	}
	for _, tc := range tests {
		for _, force := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s force=%t", tc.name, force), func(t *testing.T) {
				pair := renamed(tc.pair, fmt.Sprintf("assess_%d", time.Now().UnixNano()))
				t.Cleanup(func() { dropIndex(t, c, pair.New) })
				finding := assessPair(t, pair)
				require.Equal(t, tc.category, finding.Category)

				require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: pair.Old}}))
				require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: pair.New, MigrateForce: force}}))

				expected := pair.Old.PrimaryKeys()
				if finding.Category == plugin.AssessCategoryAutomaticallyMigratable {
					expected = pair.New.PrimaryKeys()
				}
				require.Equal(t, expected, indexedProperties(t, c, pair.New))
			})
		}
	}
}
