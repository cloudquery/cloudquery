package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
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

func withoutPrimaryKeys(table *schema.Table) *schema.Table {
	result := table.Copy(nil)
	for i := range result.Columns {
		result.Columns[i].PrimaryKey = false
	}
	return result
}

func TestInitWithNoConnectionSendsNoRequests(t *testing.T) {
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests.Add(1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	specBytes, err := json.Marshal(&Spec{Addresses: []string{server.URL}})
	require.NoError(t, err)

	p := plugin.NewPlugin("elasticsearch", "development", New)
	require.NoError(t, p.Init(context.Background(), specBytes, plugin.NewClientOptions{NoConnection: true}))

	findings, err := p.AssessTables(context.Background(), []plugin.TablePair{{New: &schema.Table{Name: "test_table"}}}, plugin.AssessOptions{})
	require.NoError(t, err)
	require.Equal(t, plugin.AssessCategoryAutomaticallyMigratable, findings[0].Category)
	require.Zero(t, requests.Load())
}

func TestNewWithNoConnectionSkipsClients(t *testing.T) {
	c := newNoConnectionClient(t)
	require.Nil(t, c.client)
	require.Nil(t, c.typedClient)
	require.Error(t, c.Close(context.Background()))
}

func TestAssessTablesListToJSONKeepsTextMapping(t *testing.T) {
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
		ForcedModeBehavior: behaviorDeleteIndices,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "tags",
			Category:           plugin.AssessCategoryNoChange,
			OldType:            `{"type":"text"}`,
			NewType:            `{"type":"text"}`,
			SafeModeBehavior:   behaviorNoChange,
			ForcedModeBehavior: behaviorDeleteIndices,
		}},
	}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesMappingChange(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "count", Type: arrow.BinaryTypes.String},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].Type = arrow.PrimitiveTypes.Int64

	require.Equal(t, plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorKeepOldMapping,
		ForcedModeBehavior: behaviorDeleteIndices,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "count",
			Category:           plugin.AssessCategoryManualMigrationRequired,
			OldType:            `{"type":"text"}`,
			NewType:            `{"type":"long"}`,
			SafeModeBehavior:   behaviorKeepOldMapping,
			ForcedModeBehavior: behaviorDeleteIndices,
		}},
	}, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
}

func TestAssessTablesConstraintChangeKeepsMapping(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "name", Type: arrow.BinaryTypes.String, NotNull: true, Unique: true},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].NotNull = false
	newTable.Columns[1].Unique = false

	finding := assessTable(t, plugin.TablePair{Old: oldTable, New: newTable})
	require.Equal(t, plugin.AssessCategoryNoChange, finding.Category)
	require.Equal(t, []plugin.ColumnFinding{{
		ColumnName:         "name",
		Category:           plugin.AssessCategoryNoChange,
		OldType:            `{"type":"text"}`,
		NewType:            `{"type":"text"}`,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorDeleteIndices,
	}}, finding.Columns)
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
	expected := plugin.TableFinding{
		TableName:          "okta_policy_rules",
		Category:           plugin.AssessCategoryAutomaticallyMigratable,
		SafeModeBehavior:   behaviorUpdateTemplate,
		ForcedModeBehavior: behaviorDeleteIndices,
		Columns: []plugin.ColumnFinding{
			{ColumnName: "policy_id", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: `{"type":"text"}`, SafeModeBehavior: behaviorAddField, ForcedModeBehavior: behaviorDeleteIndices},
			{ColumnName: "actions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: `{"type":"text"}`, SafeModeBehavior: behaviorAddField, ForcedModeBehavior: behaviorDeleteIndices},
			{ColumnName: "conditions", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: `{"type":"text"}`, SafeModeBehavior: behaviorAddField, ForcedModeBehavior: behaviorDeleteIndices},
		},
	}

	t.Run("default write mode", func(t *testing.T) {
		require.Equal(t, expected, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
	})

	t.Run("append write mode", func(t *testing.T) {
		require.Equal(t, expected, assessTable(t, plugin.TablePair{Old: withoutPrimaryKeys(oldTable), New: withoutPrimaryKeys(newTable)}))
	})
}

func TestAssessTablesIndexNameChange(t *testing.T) {
	oldTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[0].PrimaryKey = true

	expected := plugin.TableFinding{
		TableName:          "test_table",
		Category:           plugin.AssessCategoryManualMigrationRequired,
		SafeModeBehavior:   behaviorWriteToNewIndices,
		ForcedModeBehavior: behaviorWriteToNewIndices,
		Columns: []plugin.ColumnFinding{{
			ColumnName:         "id",
			Category:           plugin.AssessCategoryNoChange,
			OldType:            `{"type":"text"}`,
			NewType:            `{"type":"text"}`,
			SafeModeBehavior:   behaviorWriteToNewIndices,
			ForcedModeBehavior: behaviorWriteToNewIndices,
		}},
	}

	t.Run("primary key added", func(t *testing.T) {
		require.Equal(t, expected, assessTable(t, plugin.TablePair{Old: oldTable, New: newTable}))
	})

	t.Run("primary key removed", func(t *testing.T) {
		require.Equal(t, expected, assessTable(t, plugin.TablePair{Old: newTable, New: oldTable}))
	})
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
		OldType:            `{"type":"text"}`,
		SafeModeBehavior:   behaviorRemoveField,
		ForcedModeBehavior: behaviorDeleteIndices,
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
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryNoChange, SafeModeBehavior: behaviorNoChange, ForcedModeBehavior: behaviorDeleteIndices},
		},
		{
			name:     "added table",
			pair:     plugin.TablePair{New: table},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryAutomaticallyMigratable, SafeModeBehavior: behaviorCreateTemplate, ForcedModeBehavior: behaviorCreateTemplate},
		},
		{
			name:     "removed table",
			pair:     plugin.TablePair{Old: table},
			expected: plugin.TableFinding{TableName: "test_table", Category: plugin.AssessCategoryTableRemoved, SafeModeBehavior: behaviorKeepIndices, ForcedModeBehavior: behaviorKeepIndices},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, assessTable(t, tc.pair))
		})
	}
}
