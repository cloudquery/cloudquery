package client

import (
	"context"
	"slices"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	behaviorNoChange         = "makes no changes"
	behaviorCreateIndex      = "creates the index"
	behaviorKeepIndex        = "keeps the existing index"
	behaviorUpdateIndex      = "updates the index settings"
	behaviorRejectChanges    = "rejects the changes"
	behaviorRecreateIndex    = "deletes and recreates the index, deleting existing documents"
	behaviorNewDocumentIDs   = "keeps the index, but document IDs change, so existing documents are not replaced"
	behaviorAddAttribute     = "adds the field to the filterable and sortable attributes"
	behaviorRemoveAttribute  = "keeps the field in the index settings, new documents omit it"
	behaviorUntypedAttribute = "makes no changes, Meilisearch has no fixed field types"
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
	if pair.New == nil {
		return tableFinding(pair.Old.Name, plugin.AssessCategoryTableRemoved, behaviorKeepIndex, behaviorKeepIndex)
	}

	var oldIndex *indexSchema
	if pair.Old != nil {
		oldIndex = tableIndexSchema(pair.Old)
	}
	switch planIndexMigration(oldIndex, tableIndexSchema(pair.New)) {
	case indexCreate:
		return tableFinding(pair.New.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorCreateIndex, behaviorCreateIndex)
	case indexRecreate:
		return tableFinding(pair.New.Name, plugin.AssessCategoryManualMigrationRequired, behaviorRejectChanges, behaviorRecreateIndex)
	default:
		return assessIndexUpdate(pair.Old, pair.New)
	}
}

func assessIndexUpdate(oldTable, newTable *schema.Table) plugin.TableFinding {
	oldPrimaryKeys, newPrimaryKeys := oldTable.PrimaryKeys(), newTable.PrimaryKeys()

	category, behavior := plugin.AssessCategoryNoChange, behaviorNoChange
	var columns []plugin.ColumnFinding
	for _, change := range newTable.GetChanges(oldTable) {
		column, ok := columnFinding(change, oldPrimaryKeys, newPrimaryKeys)
		if !ok {
			continue
		}
		if column.Category == plugin.AssessCategoryAutomaticallyMigratable {
			category, behavior = plugin.AssessCategoryAutomaticallyMigratable, behaviorUpdateIndex
		}
		columns = append(columns, column)
	}
	if !slices.Equal(oldPrimaryKeys, newPrimaryKeys) {
		category, behavior = plugin.AssessCategoryManualMigrationRequired, behaviorNewDocumentIDs
	}

	finding := tableFinding(newTable.Name, category, behavior, behavior)
	finding.Columns = columns
	return finding
}

func columnFinding(change schema.TableColumnChange, oldPrimaryKeys, newPrimaryKeys []string) (plugin.ColumnFinding, bool) {
	var category plugin.AssessCategory
	var behavior string
	switch change.Type {
	case schema.TableColumnChangeTypeAdd:
		category, behavior = plugin.AssessCategoryAutomaticallyMigratable, behaviorAddAttribute
	case schema.TableColumnChangeTypeRemove:
		category, behavior = plugin.AssessCategoryNoChange, behaviorRemoveAttribute
	case schema.TableColumnChangeTypeUpdate:
		category, behavior = plugin.AssessCategoryNoChange, behaviorUntypedAttribute
	default:
		return plugin.ColumnFinding{}, false
	}
	if slices.Contains(oldPrimaryKeys, change.ColumnName) != slices.Contains(newPrimaryKeys, change.ColumnName) {
		category, behavior = plugin.AssessCategoryManualMigrationRequired, behaviorNewDocumentIDs
	}
	return plugin.ColumnFinding{
		ColumnName:         change.ColumnName,
		Category:           category,
		SafeModeBehavior:   behavior,
		ForcedModeBehavior: behavior,
	}, true
}

func tableFinding(name string, category plugin.AssessCategory, safeModeBehavior, forcedModeBehavior string) plugin.TableFinding {
	return plugin.TableFinding{
		TableName:          name,
		Category:           category,
		SafeModeBehavior:   safeModeBehavior,
		ForcedModeBehavior: forcedModeBehavior,
	}
}
