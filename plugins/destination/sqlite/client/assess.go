package client

import (
	"context"
	"slices"

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

func (c *Client) AssessTables(_ context.Context, tables []plugin.TablePair, _ plugin.AssessOptions) ([]plugin.TableFinding, error) {
	findings := make([]plugin.TableFinding, len(tables))
	for i, pair := range tables {
		findings[i] = c.assessTable(pair)
	}
	return findings, nil
}

func (c *Client) assessTable(pair plugin.TablePair) plugin.TableFinding {
	if pair.Old == nil {
		return tableFinding(pair.New.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorCreateTable, behaviorCreateTable)
	}
	if pair.New == nil {
		return tableFinding(pair.Old.Name, plugin.AssessCategoryTableRemoved, behaviorKeepTable, behaviorKeepTable)
	}

	oldTable := c.storedTable(pair.Old)
	newTable := c.normalizeTable(pair.New)
	changes := newTable.GetChanges(oldTable)
	tableAutoMigratable := canAutoMigrate(changes)

	var finding plugin.TableFinding
	switch {
	case len(changes) == 0:
		finding = tableFinding(newTable.Name, plugin.AssessCategoryNoChange, behaviorNoChange, behaviorNoChange)
	case tableAutoMigratable:
		finding = tableFinding(newTable.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorMigrateTable, behaviorMigrateTable)
	default:
		finding = tableFinding(newTable.Name, plugin.AssessCategoryManualMigrationRequired, behaviorRejectChanges, behaviorRecreateTable)
	}
	for _, change := range changes {
		if change.ColumnName == "" {
			continue
		}
		finding.Columns = append(finding.Columns, c.columnFinding(change, oldTable, newTable, tableAutoMigratable))
	}
	for _, change := range pair.New.GetChanges(pair.Old) {
		if change.Type != schema.TableColumnChangeTypeUpdate || hasColumnChange(changes, change.ColumnName) {
			continue
		}
		finding.Columns = append(finding.Columns, c.sameTypeColumnFinding(change.ColumnName, oldTable, newTable, tableAutoMigratable))
	}
	return finding
}

func (c *Client) storedTable(table *schema.Table) *schema.Table {
	stored := &schema.Table{Name: table.Name}
	for _, column := range c.normalizeTable(table).Columns {
		stored.Columns = append(stored.Columns, c.sqliteColumn(column.Name, c.arrowTypeToSqliteStr(column.Type), column.PrimaryKey, column.NotNull))
	}
	return stored
}

func hasColumnChange(changes []schema.TableColumnChange, columnName string) bool {
	return slices.ContainsFunc(changes, func(change schema.TableColumnChange) bool {
		return change.ColumnName == columnName
	})
}

func tableFinding(name string, category plugin.AssessCategory, safeModeBehavior, forcedModeBehavior string) plugin.TableFinding {
	return plugin.TableFinding{
		TableName:          name,
		Category:           category,
		SafeModeBehavior:   safeModeBehavior,
		ForcedModeBehavior: forcedModeBehavior,
	}
}

func (c *Client) columnFinding(change schema.TableColumnChange, oldTable, newTable *schema.Table, tableAutoMigratable bool) plugin.ColumnFinding {
	finding := plugin.ColumnFinding{
		ColumnName: change.ColumnName,
		OldType:    c.sqliteColumnType(oldTable, change.ColumnName),
		NewType:    c.sqliteColumnType(newTable, change.ColumnName),
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

func (c *Client) sameTypeColumnFinding(columnName string, oldTable, newTable *schema.Table, tableAutoMigratable bool) plugin.ColumnFinding {
	finding := plugin.ColumnFinding{
		ColumnName:         columnName,
		Category:           plugin.AssessCategoryNoChange,
		OldType:            c.sqliteColumnType(oldTable, columnName),
		NewType:            c.sqliteColumnType(newTable, columnName),
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
	}
	if !tableAutoMigratable {
		finding.SafeModeBehavior = behaviorRejectChanges
		finding.ForcedModeBehavior = behaviorRecreateTable
	}
	return finding
}

func (c *Client) sqliteColumnType(table *schema.Table, columnName string) string {
	column := table.Columns.Get(columnName)
	if column == nil {
		return ""
	}
	return c.arrowTypeToSqliteStr(column.Type)
}

func autoMigrationBehavior(change schema.TableColumnChange) string {
	switch change.Type {
	case schema.TableColumnChangeTypeAdd:
		return "adds the column"
	case schema.TableColumnChangeTypeRemove:
		return "keeps the column, new rows leave it empty"
	default:
		return behaviorNoChange
	}
}
