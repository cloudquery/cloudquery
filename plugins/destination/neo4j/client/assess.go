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
	behaviorKeepIndexNoPK = "keeps the index on the old primary key, new rows are created as new nodes, which can duplicate existing ones"
	behaviorKeepNodes     = "keeps the existing nodes"
)

var _ plugin.Assessor = (*Client)(nil)

func (*Client) AssessTables(_ context.Context, tables []plugin.TablePair, _ plugin.AssessOptions) ([]plugin.TableFinding, error) {
	findings := make([]plugin.TableFinding, len(tables))
	for i, pair := range tables {
		finding, err := assessTable(pair)
		if err != nil {
			return nil, err
		}
		findings[i] = finding
	}
	return findings, nil
}

func assessTable(pair plugin.TablePair) (plugin.TableFinding, error) {
	if pair.New == nil {
		return tableFinding(pair.Old.Name, plugin.AssessCategoryTableRemoved, behaviorKeepNodes), nil
	}
	if pair.Old == nil {
		category, behavior, err := indexMigration(&schema.Table{Name: pair.New.Name}, pair.New)
		return tableFinding(pair.New.Name, category, behavior), err
	}

	category, behavior, err := indexMigration(pair.Old, pair.New)
	if err != nil {
		return plugin.TableFinding{}, err
	}
	finding := tableFinding(pair.New.Name, category, behavior)
	for _, change := range pair.New.GetChanges(pair.Old) {
		if change.ColumnName == "" {
			continue
		}
		finding.Columns = append(finding.Columns, columnFinding(change, category, behavior))
	}
	return finding, nil
}

// indexMigration mirrors MigrateTables: CREATE INDEX ... IF NOT EXISTS is a no-op when an index with the same name exists,
// so a changed primary key never errors and the old index is kept in both safe and forced modes.
func indexMigration(oldTable, newTable *schema.Table) (plugin.AssessCategory, string, error) {
	oldQuery, err := createIndexQuery(oldTable)
	if err != nil {
		return plugin.AssessCategoryUnknown, "", err
	}
	newQuery, err := createIndexQuery(newTable)
	if err != nil {
		return plugin.AssessCategoryUnknown, "", err
	}
	switch {
	case oldQuery == newQuery:
		return plugin.AssessCategoryNoChange, behaviorNoChange, nil
	case oldQuery == "":
		return plugin.AssessCategoryAutomaticallyMigratable, behaviorCreateIndex, nil
	case newQuery == "":
		return plugin.AssessCategoryManualMigrationRequired, behaviorKeepIndexNoPK, nil
	default:
		return plugin.AssessCategoryManualMigrationRequired, behaviorKeepIndex, nil
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
