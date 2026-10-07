package client

import (
	"context"
	"slices"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	behaviorNoChange       = "makes no changes"
	behaviorCreateIndex    = "creates the unique primary key index"
	behaviorCreateOnWrite  = "creates the collection on the first write"
	behaviorKeepCollection = "keeps the existing collection"
	behaviorKeepIndex      = "keeps the existing unique primary key index"
	behaviorRejectChanges  = "rejects the changes"
	behaviorRecreateIndex  = "drops and recreates the unique primary key index, keeping existing documents"
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
		behavior := behaviorCreateOnWrite
		if len(primaryKeyIndexKeys(pair.New)) > 0 {
			behavior = behaviorCreateIndex
		}
		return tableFinding(pair.New.Name, plugin.AssessCategoryAutomaticallyMigratable, behavior, behavior)
	}
	if pair.New == nil {
		return tableFinding(pair.Old.Name, plugin.AssessCategoryTableRemoved, behaviorKeepCollection, behaviorKeepCollection)
	}

	finding := primaryKeyIndexFinding(pair.Old, pair.New)
	seen := make(map[string]bool)
	for _, change := range pair.New.GetChanges(pair.Old) {
		if change.ColumnName == "" || seen[change.ColumnName] {
			continue
		}
		seen[change.ColumnName] = true
		finding.Columns = append(finding.Columns, columnFinding(change.ColumnName, pair, finding))
	}
	return finding
}

func primaryKeyIndexFinding(oldTable, newTable *schema.Table) plugin.TableFinding {
	oldKeys, newKeys := primaryKeyIndexKeys(oldTable), primaryKeyIndexKeys(newTable)
	switch {
	case slices.Equal(oldKeys, newKeys):
		return tableFinding(newTable.Name, plugin.AssessCategoryNoChange, behaviorNoChange, behaviorNoChange)
	case len(newKeys) == 0:
		return tableFinding(newTable.Name, plugin.AssessCategoryNoChange, behaviorKeepIndex, behaviorKeepIndex)
	case len(oldKeys) == 0:
		return tableFinding(newTable.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorCreateIndex, behaviorCreateIndex)
	default:
		return tableFinding(newTable.Name, plugin.AssessCategoryManualMigrationRequired, behaviorRejectChanges, behaviorRecreateIndex)
	}
}

func tableFinding(name string, category plugin.AssessCategory, safeModeBehavior, forcedModeBehavior string) plugin.TableFinding {
	return plugin.TableFinding{
		TableName:          name,
		Category:           category,
		SafeModeBehavior:   safeModeBehavior,
		ForcedModeBehavior: forcedModeBehavior,
	}
}

func columnFinding(columnName string, pair plugin.TablePair, table plugin.TableFinding) plugin.ColumnFinding {
	finding := plugin.ColumnFinding{
		ColumnName:         columnName,
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
	}
	if isPrimaryKey(pair.Old, columnName) || isPrimaryKey(pair.New, columnName) {
		finding.Category = table.Category
		finding.SafeModeBehavior = table.SafeModeBehavior
		finding.ForcedModeBehavior = table.ForcedModeBehavior
	}
	return finding
}

func isPrimaryKey(table *schema.Table, columnName string) bool {
	column := table.Columns.Get(columnName)
	return column != nil && column.PrimaryKey
}
