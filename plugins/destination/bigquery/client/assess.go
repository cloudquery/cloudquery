package client

import (
	"context"
	"slices"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

const (
	behaviorNoChange             = "makes no changes"
	behaviorCreateTable          = "creates the table"
	behaviorKeepTable            = "keeps the existing table"
	behaviorMigrateTable         = "applies the changes to the existing table"
	behaviorRejectChanges        = "rejects the changes"
	behaviorAddColumn            = "adds the column"
	behaviorKeepColumn           = "keeps the column, new rows leave it empty"
	behaviorKeepColumnDefinition = "keeps the existing column definition"
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
		return tableFinding(pair.New.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorCreateTable)
	}
	if pair.New == nil {
		return tableFinding(pair.Old.Name, plugin.AssessCategoryTableRemoved, behaviorKeepTable)
	}

	oldSchema := c.bigQuerySchemaForTable(pair.Old)
	newSchema := c.bigQuerySchemaForTable(pair.New)
	_, err := mergeSchemas(oldSchema, newSchema)
	rejected := err != nil
	columns := assessColumns(oldSchema, newSchema, rejected)

	var finding plugin.TableFinding
	switch {
	case len(columns) == 0:
		finding = tableFinding(pair.New.Name, plugin.AssessCategoryNoChange, behaviorNoChange)
	case rejected:
		finding = tableFinding(pair.New.Name, plugin.AssessCategoryManualMigrationRequired, behaviorRejectChanges)
	case slices.ContainsFunc(columns, isManualMigration):
		finding = tableFinding(pair.New.Name, plugin.AssessCategoryManualMigrationRequired, behaviorMigrateTable)
	default:
		finding = tableFinding(pair.New.Name, plugin.AssessCategoryAutomaticallyMigratable, behaviorMigrateTable)
	}
	finding.Columns = columns
	return finding
}

func tableFinding(name string, category plugin.AssessCategory, behavior string) plugin.TableFinding {
	return plugin.TableFinding{
		TableName:          name,
		Category:           category,
		SafeModeBehavior:   behavior,
		ForcedModeBehavior: behavior,
	}
}

func isManualMigration(finding plugin.ColumnFinding) bool {
	return finding.Category == plugin.AssessCategoryManualMigrationRequired
}

func assessColumns(oldSchema, newSchema bigquery.Schema, tableRejected bool) []plugin.ColumnFinding {
	oldFields := fieldsByName(oldSchema)
	newFields := fieldsByName(newSchema)
	var findings []plugin.ColumnFinding
	for _, have := range oldSchema {
		want := newFields[have.Name]
		if want != nil && bigQueryTypeName(have) == bigQueryTypeName(want) {
			continue
		}
		findings = append(findings, columnFinding(have.Name, have, want, tableRejected))
	}
	for _, want := range newSchema {
		if oldFields[want.Name] == nil {
			findings = append(findings, columnFinding(want.Name, nil, want, tableRejected))
		}
	}
	return findings
}

func columnFinding(name string, have, want *bigquery.FieldSchema, tableRejected bool) plugin.ColumnFinding {
	category, behavior := columnChange(have, want)
	if tableRejected {
		behavior = behaviorRejectChanges
	}
	return plugin.ColumnFinding{
		ColumnName:         name,
		Category:           category,
		OldType:            optionalTypeName(have),
		NewType:            optionalTypeName(want),
		SafeModeBehavior:   behavior,
		ForcedModeBehavior: behavior,
	}
}

func columnChange(have, want *bigquery.FieldSchema) (plugin.AssessCategory, string) {
	switch {
	case have == nil:
		return plugin.AssessCategoryAutomaticallyMigratable, behaviorAddColumn
	case checkFieldMigration(have, want) != nil:
		return plugin.AssessCategoryManualMigrationRequired, behaviorRejectChanges
	case want == nil:
		return plugin.AssessCategoryAutomaticallyMigratable, behaviorKeepColumn
	default:
		// mergeSchemas only compares top-level types, so mode and nested-field changes keep the old definition
		return plugin.AssessCategoryManualMigrationRequired, behaviorKeepColumnDefinition
	}
}

func optionalTypeName(field *bigquery.FieldSchema) string {
	if field == nil {
		return ""
	}
	return bigQueryTypeName(field)
}
