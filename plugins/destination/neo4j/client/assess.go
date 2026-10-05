package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	behaviorNoChange      = "makes no changes"
	behaviorNoFixedSchema = "makes no changes, properties have no fixed schema"
	behaviorCreateIndex   = "creates the primary key index"
	behaviorKeepIndex     = "keeps the index on the old primary key, new rows are merged on the new primary key"
	behaviorKeepNodes     = "keeps the existing nodes"
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
		return tableFinding(pair.Old.Name, plugin.AssessCategoryTableRemoved, behaviorKeepNodes)
	}
	if pair.Old == nil {
		category, behavior := indexMigration(&schema.Table{Name: pair.New.Name}, pair.New)
		return tableFinding(pair.New.Name, category, behavior)
	}

	category, behavior := indexMigration(pair.Old, pair.New)
	finding := tableFinding(pair.New.Name, category, behavior)
	for _, change := range pair.New.GetChanges(pair.Old) {
		if change.ColumnName == "" {
			continue
		}
		finding.Columns = append(finding.Columns, columnFinding(change, category, behavior))
	}
	return finding
}

// indexMigration mirrors MigrateTables: CREATE INDEX ... IF NOT EXISTS is a no-op when an index with the same name exists,
// so a changed primary key never errors and the old index is kept in both safe and forced modes.
func indexMigration(oldTable, newTable *schema.Table) (plugin.AssessCategory, string) {
	switch createIndexQuery(oldTable) {
	case createIndexQuery(newTable):
		return plugin.AssessCategoryNoChange, behaviorNoChange
	case "":
		return plugin.AssessCategoryAutomaticallyMigratable, behaviorCreateIndex
	default:
		return plugin.AssessCategoryManualMigrationRequired, behaviorKeepIndex
	}
}

func tableFinding(name string, category plugin.AssessCategory, behavior string) plugin.TableFinding {
	return plugin.TableFinding{
		TableName:          name,
		Category:           category,
		SafeModeBehavior:   behavior,
		ForcedModeBehavior: behavior,
	}
}

func columnFinding(change schema.TableColumnChange, indexCategory plugin.AssessCategory, indexBehavior string) plugin.ColumnFinding {
	finding := plugin.ColumnFinding{
		ColumnName:         change.ColumnName,
		Category:           plugin.AssessCategoryNoChange,
		SafeModeBehavior:   behaviorNoFixedSchema,
		ForcedModeBehavior: behaviorNoFixedSchema,
	}
	if indexCategory != plugin.AssessCategoryNoChange && (change.Previous.PrimaryKey || change.Current.PrimaryKey) {
		finding.Category = indexCategory
		finding.SafeModeBehavior = indexBehavior
		finding.ForcedModeBehavior = indexBehavior
	}
	return finding
}
