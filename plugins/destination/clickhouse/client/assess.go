package client

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v8/typeconv"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v8/typeconv/ch/types"
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
		return tableFinding(pair.New.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorCreateTable, behaviorCreateTable), nil
	}
	if pair.New == nil {
		return tableFinding(pair.Old.Name, plugin.AssessCategoryTableRemoved, behaviorKeepTable, behaviorKeepTable), nil
	}

	oldTable, err := typeconv.CanonizedTable(pair.Old)
	if err != nil {
		return plugin.TableFinding{}, err
	}
	newTable, err := typeconv.CanonizedTable(pair.New)
	if err != nil {
		return plugin.TableFinding{}, err
	}
	changes := newTable.GetChanges(oldTable)
	keyChanges, err := c.partitionOrSortingKeyChangesBetween(oldTable, newTable)
	if err != nil {
		return plugin.TableFinding{}, err
	}

	if len(changes) == 0 && len(keyChanges) == 0 {
		return tableFinding(newTable.Name, plugin.AssessCategoryNoChange, behaviorNoChange, behaviorNoChange), nil
	}

	tableAutoMigratable := !forceMigrationNeeded(changes, keyChanges)
	finding := tableFinding(newTable.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorMigrateTable, behaviorMigrateTable)
	if !tableAutoMigratable {
		finding = tableFinding(newTable.Name, plugin.AssessCategoryManualMigrationRequired, behaviorRejectChanges, behaviorRecreateTable)
	}
	for _, change := range changes {
		columnFinding, err := columnFinding(change, oldTable, newTable, tableAutoMigratable)
		if err != nil {
			return plugin.TableFinding{}, err
		}
		finding.Columns = append(finding.Columns, columnFinding)
	}
	return finding, nil
}

func (c *Client) partitionOrSortingKeyChangesBetween(oldTable, newTable *schema.Table) ([]tableSchemaChange, error) {
	havePartitionKey, haveSortingKey, err := resolvePartitionKeyAndSortingKey(oldTable, c.spec.Partition, c.spec.OrderBy)
	if err != nil {
		return nil, err
	}
	wantPartitionKey, wantSortingKey, err := resolvePartitionKeyAndSortingKey(newTable, c.spec.Partition, c.spec.OrderBy)
	if err != nil {
		return nil, err
	}
	return partitionOrSortingKeyChanges(havePartitionKey, haveSortingKey, wantPartitionKey, wantSortingKey), nil
}

func tableFinding(name string, category plugin.AssessCategory, safeModeBehavior, forcedModeBehavior string) plugin.TableFinding {
	return plugin.TableFinding{
		TableName:          name,
		Category:           category,
		SafeModeBehavior:   safeModeBehavior,
		ForcedModeBehavior: forcedModeBehavior,
	}
}

func columnFinding(change schema.TableColumnChange, oldTable, newTable *schema.Table, tableAutoMigratable bool) (plugin.ColumnFinding, error) {
	oldType, err := clickHouseColumnType(oldTable, change.ColumnName)
	if err != nil {
		return plugin.ColumnFinding{}, err
	}
	newType, err := clickHouseColumnType(newTable, change.ColumnName)
	if err != nil {
		return plugin.ColumnFinding{}, err
	}
	finding := plugin.ColumnFinding{
		ColumnName: change.ColumnName,
		OldType:    oldType,
		NewType:    newType,
	}
	if tableAutoMigratable || !needsTableDrop(change) {
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
	return finding, nil
}

func clickHouseColumnType(table *schema.Table, columnName string) (string, error) {
	column := table.Columns.Get(columnName)
	if column == nil {
		return "", nil
	}
	return types.FieldType(column.ToArrowField())
}

func autoMigrationBehavior(change schema.TableColumnChange) string {
	if change.Type == schema.TableColumnChangeTypeAdd {
		return "adds the column"
	}
	return "keeps the column, new rows leave it empty"
}
