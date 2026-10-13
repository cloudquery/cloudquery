package client

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var _ plugin.Assessor = (*Client)(nil)

func (*Client) AssessTables(_ context.Context, tables []plugin.TablePair, _ plugin.AssessOptions) ([]plugin.TableFinding, error) {
	findings := make([]plugin.TableFinding, len(tables))
	for i, pair := range tables {
		findings[i] = assessTable(pair)
	}
	return findings, nil
}

func assessTable(pair plugin.TablePair) plugin.TableFinding {
	if pair.Old == nil {
		return plugin.TableFinding{TableName: pair.New.Name, Category: plugin.AssessCategoryFileSchemaChanged}
	}
	if pair.New == nil {
		return plugin.TableFinding{TableName: pair.Old.Name, Category: plugin.AssessCategoryTableRemoved}
	}

	finding := plugin.TableFinding{TableName: pair.New.Name}
	var unknownReasons []string
	for _, newColumn := range pair.New.Columns {
		oldColumn := pair.Old.Columns.Get(newColumn.Name)
		if oldColumn == nil {
			finding.Columns = append(finding.Columns, plugin.ColumnFinding{ColumnName: newColumn.Name, Category: plugin.AssessCategoryFileSchemaChanged})
			continue
		}
		if arrow.TypeEqual(oldColumn.Type, newColumn.Type) {
			continue
		}
		column, err := assessTypeChange(pair.New.Name, *oldColumn, newColumn)
		if err != nil {
			unknownReasons = append(unknownReasons, fmt.Sprintf("column %s: %v", newColumn.Name, err))
		}
		finding.Columns = append(finding.Columns, column)
	}
	for _, oldColumn := range pair.Old.Columns {
		if pair.New.Columns.Get(oldColumn.Name) == nil {
			finding.Columns = append(finding.Columns, plugin.ColumnFinding{ColumnName: oldColumn.Name, Category: plugin.AssessCategoryFileSchemaChanged})
		}
	}

	finding.Category = tableCategory(finding.Columns)
	finding.IncompleteCoverageReason = strings.Join(unknownReasons, "; ")
	return finding
}

func assessTypeChange(tableName string, oldColumn, newColumn schema.Column) (plugin.ColumnFinding, error) {
	finding := plugin.ColumnFinding{ColumnName: newColumn.Name, Category: plugin.AssessCategoryUnknown}
	oldField, newField := oldColumn.ToArrowField(), newColumn.ToArrowField()
	pairs, err := schema.SyntheticPairs(oldField, newField)
	if err != nil {
		return finding, err
	}
	oldRecord, newRecord, err := schema.SyntheticRecords(oldField, newField, pairs)
	if err != nil {
		return finding, err
	}
	defer oldRecord.Release()
	defer newRecord.Release()

	outputChanged := false
	evidence := make([]plugin.Evidence, len(pairs))
	for i, pair := range pairs {
		before, err := marshalRow(tableName, oldRecord, i)
		if err != nil {
			return finding, err
		}
		after, err := marshalRow(tableName, newRecord, i)
		if err != nil {
			return finding, err
		}
		outputChanged = outputChanged || !bytes.Equal(before, after)
		evidence[i] = plugin.Evidence{SyntheticValue: pair.Value, Before: string(before), After: string(after)}
	}

	finding.Category = plugin.AssessCategoryNoChange
	if outputChanged {
		finding.Category = plugin.AssessCategoryFileSchemaChanged
	}
	finding.Evidence = evidence
	return finding, nil
}

func tableCategory(columns []plugin.ColumnFinding) plugin.AssessCategory {
	category := plugin.AssessCategoryNoChange
	for _, column := range columns {
		switch column.Category {
		case plugin.AssessCategoryFileSchemaChanged:
			return plugin.AssessCategoryFileSchemaChanged
		case plugin.AssessCategoryUnknown:
			category = plugin.AssessCategoryUnknown
		}
	}
	return category
}
