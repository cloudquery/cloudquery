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

func assessTables(t *testing.T, pairs ...plugin.TablePair) []plugin.TableFinding {
	t.Helper()
	ctx := context.Background()
	p := plugin.NewPlugin("firehose", "development", New)
	require.NoError(t, p.Init(ctx, nil, plugin.NewClientOptions{NoConnection: true}))
	findings, err := p.AssessTables(ctx, pairs, plugin.AssessOptions{})
	require.NoError(t, err)
	return findings
}

func typeChange(name string, oldType, newType arrow.DataType) plugin.TablePair {
	return plugin.TablePair{
		Old: &schema.Table{Name: name, Columns: schema.ColumnList{{Name: "tags", Type: oldType}}},
		New: &schema.Table{Name: name, Columns: schema.ColumnList{{Name: "tags", Type: newType}}},
	}
}

func TestNewWithNoConnectionSkipsAWSClient(t *testing.T) {
	c, err := New(context.Background(), zerolog.Nop(), nil, plugin.NewClientOptions{NoConnection: true})
	require.NoError(t, err)
	require.Nil(t, c.(*Client).firehoseClient)
}

func TestAssessTablesDatadogTags(t *testing.T) {
	findings := assessTables(t, typeChange("datadog_monitors", arrow.ListOf(arrow.BinaryTypes.String), types.ExtensionTypes.JSON))

	escaped := `{"_cq_table_name":"datadog_monitors","tags":["comma, \"quote\", back\\slash\nnew line\ttab é"]}`
	require.Equal(t, []plugin.TableFinding{{
		TableName: "datadog_monitors",
		Category:  plugin.AssessCategoryNoChange,
		Columns: []plugin.ColumnFinding{{
			ColumnName: "tags",
			Category:   plugin.AssessCategoryNoChange,
			Evidence: []plugin.Evidence{
				{SyntheticValue: `["env:prod"]`, Before: `{"_cq_table_name":"datadog_monitors","tags":["env:prod"]}`, After: `{"_cq_table_name":"datadog_monitors","tags":["env:prod"]}`},
				{SyntheticValue: `null`, Before: `{"_cq_table_name":"datadog_monitors","tags":null}`, After: `{"_cq_table_name":"datadog_monitors","tags":null}`},
				{SyntheticValue: `[]`, Before: `{"_cq_table_name":"datadog_monitors","tags":[]}`, After: `{"_cq_table_name":"datadog_monitors","tags":[]}`},
				{SyntheticValue: `["comma, \"quote\", back\\slash\nnew line\ttab é"]`, Before: escaped, After: escaped},
			},
		}},
	}}, findings)
}

func TestAssessTablesOutputChanged(t *testing.T) {
	findings := assessTables(t, typeChange("intervals", arrow.FixedWidthTypes.MonthDayNanoInterval, types.ExtensionTypes.JSON))

	require.Equal(t, []plugin.TableFinding{{
		TableName: "intervals",
		Category:  plugin.AssessCategoryFileSchemaChanged,
		Columns: []plugin.ColumnFinding{{
			ColumnName: "tags",
			Category:   plugin.AssessCategoryFileSchemaChanged,
			Evidence: []plugin.Evidence{
				{
					SyntheticValue: `{"months":1,"days":1,"nanoseconds":1}`,
					Before:         `{"_cq_table_name":"intervals","tags":{"months":1,"days":1,"nanoseconds":1}}`,
					After:          `{"_cq_table_name":"intervals","tags":{"days":1,"months":1,"nanoseconds":1}}`,
				},
				{SyntheticValue: `null`, Before: `{"_cq_table_name":"intervals","tags":null}`, After: `{"_cq_table_name":"intervals","tags":null}`},
			},
		}},
	}}, findings)
}

func TestAssessTablesUnableToCompare(t *testing.T) {
	findings := assessTables(t, typeChange("numbers", arrow.PrimitiveTypes.Int64, arrow.BinaryTypes.String))

	require.Equal(t, []plugin.TableFinding{{
		TableName:                "numbers",
		Category:                 plugin.AssessCategoryUnknown,
		Columns:                  []plugin.ColumnFinding{{ColumnName: "tags", Category: plugin.AssessCategoryUnknown}},
		IncompleteCoverageReason: "column tags: unable to compare: no equivalent value for int64 and utf8",
	}}, findings)
}

func TestAssessTablesSchemaChanges(t *testing.T) {
	oldTable := &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "created_on", Type: arrow.FixedWidthTypes.Timestamp_us},
	}}
	newTable := &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String},
		{Name: "policy_id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}

	findings := assessTables(t,
		plugin.TablePair{Old: oldTable, New: oldTable},
		plugin.TablePair{Old: oldTable, New: newTable},
		plugin.TablePair{New: newTable},
		plugin.TablePair{Old: oldTable},
	)

	require.Equal(t, []plugin.TableFinding{
		{TableName: "okta_policy_rules", Category: plugin.AssessCategoryNoChange},
		{
			TableName: "okta_policy_rules",
			Category:  plugin.AssessCategoryFileSchemaChanged,
			Columns: []plugin.ColumnFinding{
				{ColumnName: "policy_id", Category: plugin.AssessCategoryFileSchemaChanged},
				{ColumnName: "created_on", Category: plugin.AssessCategoryFileSchemaChanged},
			},
		},
		{TableName: "okta_policy_rules", Category: plugin.AssessCategoryFileSchemaChanged},
		{TableName: "okta_policy_rules", Category: plugin.AssessCategoryTableRemoved},
	}, findings)
}

func TestAssessTablesMixedResults(t *testing.T) {
	changed := schema.Column{Name: "interval", Type: arrow.FixedWidthTypes.MonthDayNanoInterval}
	unknown := schema.Column{Name: "count", Type: arrow.PrimitiveTypes.Int64}
	changedNew := schema.Column{Name: "interval", Type: types.ExtensionTypes.JSON}
	unknownNew := schema.Column{Name: "count", Type: arrow.BinaryTypes.String}

	for name, columns := range map[string][2]schema.ColumnList{
		"changed first": {{changed, unknown}, {changedNew, unknownNew}},
		"unknown first": {{unknown, changed}, {unknownNew, changedNew}},
	} {
		t.Run(name, func(t *testing.T) {
			findings := assessTables(t, plugin.TablePair{
				Old: &schema.Table{Name: "mixed", Columns: columns[0]},
				New: &schema.Table{Name: "mixed", Columns: columns[1]},
			})

			require.Len(t, findings, 1)
			require.Equal(t, plugin.AssessCategoryFileSchemaChanged, findings[0].Category)
			require.Equal(t, "column count: unable to compare: no equivalent value for int64 and utf8", findings[0].IncompleteCoverageReason)
		})
	}
}
