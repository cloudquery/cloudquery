package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	behaviorNoChange          = "makes no changes"
	behaviorCreateTemplate    = "creates the index template"
	behaviorReplaceIndices    = "deletes existing indices that match the index name and creates the index template, deleting their documents"
	behaviorKeepIndices       = "keeps the existing indices"
	behaviorUpdateTemplate    = "updates the index template"
	behaviorAddField          = "adds the field to the index template"
	behaviorRemoveField       = "removes the field from the index template, existing documents keep it"
	behaviorKeepOldMapping    = "updates the index template, existing indices keep the old mapping"
	behaviorDeleteIndices     = "deletes the existing indices and replaces the index template, deleting existing documents"
	behaviorWriteToNewIndices = "updates the index template and writes to new indices, existing indices are kept"
	behaviorDeleteNewIndices  = "deletes existing indices that match the new index name, keeps indices with the old name, updates the index template and writes to new indices"
)

var _ plugin.Assessor = (*Client)(nil)

func (c *Client) AssessTables(_ context.Context, tables []plugin.TablePair, _ plugin.AssessOptions) ([]plugin.TableFinding, error) {
	findings := make([]plugin.TableFinding, len(tables))
	for i, pair := range tables {
		finding, err := c.assessTable(pair)
		if err != nil {
			return nil, err
		}
		findings[i] = finding
	}
	return findings, nil
}

func (c *Client) assessTable(pair plugin.TablePair) (plugin.TableFinding, error) {
	if pair.Old == nil {
		return tableFinding(pair.New.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorCreateTemplate, behaviorReplaceIndices), nil
	}
	if pair.New == nil {
		return tableFinding(pair.Old.Name, plugin.AssessCategoryTableRemoved, behaviorKeepIndices, behaviorKeepIndices), nil
	}

	columns, err := columnFindings(pair)
	if err != nil {
		return plugin.TableFinding{}, err
	}
	finding := tableFinding(pair.New.Name, tableCategory(columns), "", behaviorDeleteIndices)
	indexChanged := c.getIndexNamePattern(pair.Old) != c.getIndexNamePattern(pair.New)
	switch {
	case indexChanged:
		finding.Category = plugin.AssessCategoryManualMigrationRequired
		finding.SafeModeBehavior = behaviorWriteToNewIndices
		finding.ForcedModeBehavior = behaviorDeleteNewIndices
	case finding.Category == plugin.AssessCategoryNoChange:
		finding.SafeModeBehavior = behaviorNoChange
	case finding.Category == plugin.AssessCategoryAutomaticallyMigratable:
		finding.SafeModeBehavior = behaviorUpdateTemplate
	default:
		finding.SafeModeBehavior = behaviorKeepOldMapping
	}
	for i := range columns {
		columns[i].ForcedModeBehavior = finding.ForcedModeBehavior
		if indexChanged {
			columns[i].SafeModeBehavior = finding.SafeModeBehavior
		}
	}
	finding.Columns = columns
	return finding, nil
}

func tableFinding(name string, category plugin.AssessCategory, safeModeBehavior, forcedModeBehavior string) plugin.TableFinding {
	return plugin.TableFinding{
		TableName:          name,
		Category:           category,
		SafeModeBehavior:   safeModeBehavior,
		ForcedModeBehavior: forcedModeBehavior,
	}
}

func tableCategory(columns []plugin.ColumnFinding) plugin.AssessCategory {
	category := plugin.AssessCategoryNoChange
	for _, column := range columns {
		switch column.Category {
		case plugin.AssessCategoryManualMigrationRequired:
			return column.Category
		case plugin.AssessCategoryAutomaticallyMigratable:
			category = column.Category
		}
	}
	return category
}

func columnFindings(pair plugin.TablePair) ([]plugin.ColumnFinding, error) {
	oldMappings, err := fieldMappings(pair.Old)
	if err != nil {
		return nil, err
	}
	newMappings, err := fieldMappings(pair.New)
	if err != nil {
		return nil, err
	}
	var findings []plugin.ColumnFinding
	seen := make(map[string]bool)
	for _, change := range pair.New.GetChanges(pair.Old) {
		if change.ColumnName == "" || seen[change.ColumnName] {
			continue
		}
		seen[change.ColumnName] = true
		findings = append(findings, columnFinding(change.ColumnName, oldMappings, newMappings))
	}
	return findings, nil
}

func columnFinding(columnName string, oldMappings, newMappings map[string]string) plugin.ColumnFinding {
	oldMapping, inOld := oldMappings[columnName]
	newMapping, inNew := newMappings[columnName]
	finding := plugin.ColumnFinding{
		ColumnName: columnName,
		OldType:    oldMapping,
		NewType:    newMapping,
	}
	switch {
	case !inOld:
		finding.Category = plugin.AssessCategoryAutomaticallyMigratable
		finding.SafeModeBehavior = behaviorAddField
	case !inNew:
		finding.Category = plugin.AssessCategoryAutomaticallyMigratable
		finding.SafeModeBehavior = behaviorRemoveField
	case oldMapping == newMapping:
		finding.Category = plugin.AssessCategoryNoChange
		finding.SafeModeBehavior = behaviorNoChange
	default:
		finding.Category = plugin.AssessCategoryManualMigrationRequired
		finding.SafeModeBehavior = behaviorKeepOldMapping
	}
	return finding
}

func fieldMappings(table *schema.Table) (map[string]string, error) {
	properties := indexProperties(table)
	mappings := make(map[string]string, len(properties))
	for name, property := range properties {
		b, err := json.Marshal(property)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal mapping of column %s in table %s: %w", name, table.Name, err)
		}
		mappings[name] = string(b)
	}
	return mappings, nil
}
