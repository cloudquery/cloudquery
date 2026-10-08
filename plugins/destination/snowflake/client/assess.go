package client

import (
	"context"
	"strings"

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
	behaviorMigrateFails  = "fails, because it adds the primary key before the new column exists"
)

var _ plugin.Assessor = (*Client)(nil)

func (c *Client) AssessTables(_ context.Context, tables []plugin.TablePair, _ plugin.AssessOptions) ([]plugin.TableFinding, error) {
	findings := make([]plugin.TableFinding, len(tables))
	for i, pair := range tables {
		findings[i] = c.assessTable(pair)
	}
	return findings, nil
}

func (*Client) assessTable(pair plugin.TablePair) plugin.TableFinding {
	if pair.Old == nil {
		return tableFinding(pair.New.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorCreateTable, behaviorCreateTable)
	}
	if pair.New == nil {
		return tableFinding(pair.Old.Name, plugin.AssessCategoryTableRemoved, behaviorKeepTable, behaviorKeepTable)
	}

	oldTable := normalizeTable(pair.Old)
	newTable := normalizeTable(pair.New)
	changes := getTableChangesCaseInsensitive(newTable, oldTable)
	tableAutoMigratable := canAutoMigrate(changes)

	var finding plugin.TableFinding
	switch {
	case len(changes) == 0:
		finding = tableFinding(newTable.Name, plugin.AssessCategoryNoChange, behaviorNoChange, behaviorNoChange)
	case tableAutoMigratable && addsPrimaryKeyColumn(changes):
		return migrationFailsFinding(newTable.Name, changes, oldTable, newTable)
	case tableAutoMigratable:
		finding = tableFinding(newTable.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorMigrateTable, behaviorMigrateTable)
	default:
		finding = tableFinding(newTable.Name, plugin.AssessCategoryManualMigrationRequired, behaviorRejectChanges, behaviorRecreateTable)
	}
	for _, change := range changes {
		finding.Columns = append(finding.Columns, columnFinding(change, oldTable, newTable, tableAutoMigratable))
	}
	for _, columnName := range columnsStoredUnchanged(pair, changes) {
		finding.Columns = append(finding.Columns, unchangedColumnFinding(columnName, oldTable, newTable, tableAutoMigratable))
	}
	return finding
}

func addsPrimaryKeyColumn(changes []schema.TableColumnChange) bool {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd && change.Current.PrimaryKey {
			return true
		}
	}
	return false
}

func migrationFailsFinding(name string, changes []schema.TableColumnChange, oldTable, newTable *schema.Table) plugin.TableFinding {
	finding := tableFinding(name, plugin.AssessCategoryManualMigrationRequired, behaviorMigrateFails, behaviorMigrateFails)
	for _, change := range changes {
		finding.Columns = append(finding.Columns, plugin.ColumnFinding{
			ColumnName:         change.ColumnName,
			Category:           plugin.AssessCategoryManualMigrationRequired,
			OldType:            snowflakeColumnType(oldTable, change.ColumnName),
			NewType:            snowflakeColumnType(newTable, change.ColumnName),
			SafeModeBehavior:   behaviorMigrateFails,
			ForcedModeBehavior: behaviorMigrateFails,
		})
	}
	return finding
}

func normalizeTable(table *schema.Table) *schema.Table {
	normalized := table.Copy(nil)
	for i, column := range normalized.Columns {
		normalized.Columns[i].Type = effectiveType(column.Type)
	}
	return normalized
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
		OldType:    snowflakeColumnType(oldTable, change.ColumnName),
		NewType:    snowflakeColumnType(newTable, change.ColumnName),
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

// columnsStoredUnchanged lists source column changes that Snowflake stores identically, e.g. list<string> and json both map to variant.
func columnsStoredUnchanged(pair plugin.TablePair, snowflakeChanges []schema.TableColumnChange) []string {
	seen := make(map[string]bool, len(snowflakeChanges))
	for _, change := range snowflakeChanges {
		seen[change.ColumnName] = true
	}
	var columnNames []string
	for _, change := range getTableChangesCaseInsensitive(pair.New, pair.Old) {
		if seen[change.ColumnName] {
			continue
		}
		seen[change.ColumnName] = true
		columnNames = append(columnNames, change.ColumnName)
	}
	return columnNames
}

func unchangedColumnFinding(columnName string, oldTable, newTable *schema.Table, tableAutoMigratable bool) plugin.ColumnFinding {
	finding := plugin.ColumnFinding{
		ColumnName:         columnName,
		Category:           plugin.AssessCategoryNoChange,
		OldType:            snowflakeColumnType(oldTable, columnName),
		NewType:            snowflakeColumnType(newTable, columnName),
		SafeModeBehavior:   behaviorNoChange,
		ForcedModeBehavior: behaviorNoChange,
	}
	if !tableAutoMigratable {
		finding.SafeModeBehavior = behaviorRejectChanges
		finding.ForcedModeBehavior = behaviorRecreateTable
	}
	return finding
}

func snowflakeColumnType(table *schema.Table, columnName string) string {
	column := findColumn(table.Columns, columnName)
	if column == nil {
		return ""
	}
	return SchemaTypeToSnowflake(column.Type)
}

func autoMigrationBehavior(change schema.TableColumnChange) string {
	var actions []string
	switch change.Type {
	case schema.TableColumnChangeTypeAdd:
		actions = append(actions, "adds the column")
	case schema.TableColumnChangeTypeRemove:
		actions = append(actions, "keeps the column, new rows leave it empty")
		if change.Previous.PrimaryKey {
			actions = append(actions, "drops the primary key")
		}
	case schema.TableColumnChangeTypeRemoveUniqueConstraint:
		actions = append(actions, "drops the unique constraint")
	case schema.TableColumnChangeTypeUpdate:
		if change.Previous.PrimaryKey != change.Current.PrimaryKey {
			actions = append(actions, "updates the primary key")
		}
		if change.Previous.NotNull && !change.Current.NotNull {
			actions = append(actions, "drops the not null constraint")
		}
	}
	return strings.Join(actions, ", ")
}
