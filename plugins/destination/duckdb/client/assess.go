package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	behaviorNoChange      = "makes no changes"
	behaviorCreateTable   = "creates the table"
	behaviorKeepTable     = "keeps the existing table"
	behaviorMigrateTable  = "applies the changes to the existing table"
	behaviorRejectChanges = "rejects the changes"
	behaviorRecreateTable = "drops and recreates the table, deleting existing rows"
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
		return tableFinding(pair.New.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorCreateTable, behaviorCreateTable)
	}
	if pair.New == nil {
		return tableFinding(pair.Old.Name, plugin.AssessCategoryTableRemoved, behaviorKeepTable, behaviorKeepTable)
	}

	oldTable := normalizeTable(pair.Old)
	newTable := normalizeTable(pair.New)
	changes := newTable.GetChanges(oldTable)
	if len(changes) == 0 {
		return tableFinding(newTable.Name, plugin.AssessCategoryNoChange, behaviorNoChange, behaviorNoChange)
	}

	tableAutoMigratable := canAutoMigrate(changes)
	finding := tableFinding(newTable.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorMigrateTable, behaviorMigrateTable)
	if !tableAutoMigratable {
		finding = tableFinding(newTable.Name, plugin.AssessCategoryManualMigrationRequired, behaviorRejectChanges, behaviorRecreateTable)
	}
	for _, change := range changes {
		if change.ColumnName == "" {
			continue
		}
		finding.Columns = append(finding.Columns, columnFinding(change, oldTable, newTable, tableAutoMigratable))
	}
	return finding
}

func tableFinding(name string, category plugin.AssessCategory, safeModeBehavior, forcedModeBehavior string) plugin.TableFinding {
	return plugin.TableFinding{
		TableName:          name,
		Category:           category,
		SafeModeBehavior:   safeModeBehavior,
		ForcedModeBehavior: forcedModeBehavior,
	}
}

func columnFinding(change schema.TableColumnChange, oldTable, newTable *schema.Table, tableAutoMigratable bool) plugin.ColumnFinding {
	finding := plugin.ColumnFinding{
		ColumnName: change.ColumnName,
		OldType:    duckDBColumnType(oldTable, change.ColumnName),
		NewType:    duckDBColumnType(newTable, change.ColumnName),
	}
	if tableAutoMigratable || canAutoMigrate([]schema.TableColumnChange{change}) {
		finding.Category = plugin.AssessCategoryAutomaticallyMigratable
	} else {
		finding.Category = plugin.AssessCategoryManualMigrationRequired
	}
	if tableAutoMigratable {
		finding.SafeModeBehavior = autoMigrationBehavior(change)
		finding.ForcedModeBehavior = finding.SafeModeBehavior
	} else {
		finding.SafeModeBehavior = behaviorRejectChanges
		finding.ForcedModeBehavior = behaviorRecreateTable
	}
	return finding
}

func duckDBColumnType(table *schema.Table, columnName string) string {
	column := table.Columns.Get(columnName)
	if column == nil {
		return ""
	}
	return duckDBType(*column)
}

func autoMigrationBehavior(change schema.TableColumnChange) string {
	if change.Type == schema.TableColumnChangeTypeAdd {
		return "adds the column"
	}
	return "keeps the column, new rows leave it empty"
}
