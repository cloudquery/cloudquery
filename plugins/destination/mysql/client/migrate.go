package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

func normalizeSchemas(tables schema.Schemas) (schema.Schemas, error) {
	var normalized schema.Schemas
	for _, sc := range tables {
		tableName := schema.TableName(sc)
		fields := make([]arrow.Field, 0)
		for _, f := range sc.Fields() {
			keys := make([]string, 0)
			values := make([]string, 0)
			origKeys := f.Metadata.Keys()
			origValues := f.Metadata.Values()
			for k, v := range origKeys {
				if v != schema.MetadataUnique {
					keys = append(keys, v)
					values = append(values, origValues[k])
				}
			}
			normalizedType, err := mySQLTypeToArrowType(tableName, f.Name, arrowTypeToMySqlStr(f.Type))
			if err != nil {
				return nil, err
			}
			fields = append(fields, arrow.Field{
				Name:     f.Name,
				Type:     normalizedType,
				Nullable: f.Nullable && !schema.IsPk(f),
				Metadata: arrow.NewMetadata(keys, values),
			})
		}

		md := sc.Metadata()
		normalized = append(normalized, arrow.NewSchema(fields, &md))
	}

	return normalized, nil
}

func (c *Client) nonAutoMigrtableTables(tables schema.Schemas, schemaTables schema.Schemas) (names []string, changes [][]schema.FieldChange) {
	var tableChanges [][]schema.FieldChange
	for _, t := range tables {
		tableName := schema.TableName(t)
		schemaTable := schemaTables.SchemaByName(tableName)
		if schemaTable == nil {
			continue
		}
		changes := schema.GetSchemaChanges(t, schemaTable)
		if !c.canAutoMigrate(changes) {
			names = append(names, tableName)
			tableChanges = append(tableChanges, changes)
		}
	}
	return names, tableChanges
}

func (*Client) canAutoMigrate(changes []schema.FieldChange) bool {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd && (schema.IsPk(change.Current) || !change.Current.Nullable) {
			return false
		}

		if change.Type == schema.TableColumnChangeTypeRemove && (schema.IsPk(change.Previous) || !change.Previous.Nullable) {
			return false
		}

		if change.Type == schema.TableColumnChangeTypeUpdate {
			return false
		}
	}
	return true
}

func (c *Client) autoMigrateTable(ctx context.Context, table *arrow.Schema, changes []schema.FieldChange) error {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd {
			err := c.addColumn(ctx, table, change.Current)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, tables schema.Schemas) error {
	schemaTables, err := c.schemaTables(ctx, tables)
	if err != nil {
		return err
	}

	normalizedTables, err := normalizeSchemas(tables)
	if err != nil {
		return err
	}

	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonAutoMigrtableTables, changes := c.nonAutoMigrtableTables(normalizedTables, schemaTables)
		if len(nonAutoMigrtableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonAutoMigrtableTables, ","), changes)
		}
	}

	for _, table := range normalizedTables {
		tableName := schema.TableName(table)
		if tableName == "" {
			return fmt.Errorf("schema %s has no table name", table.String())
		}
		c.logger.Info().Str("table", tableName).Msg("Migrating table")
		if len(table.Fields()) == 0 {
			c.logger.Info().Str("table", tableName).Msg("Table with no columns, skipping")
			continue
		}
		schemaTable := schemaTables.SchemaByName(tableName)
		if schemaTable == nil {
			c.logger.Info().Str("table", tableName).Msg("Table doesn't exist, creating")
			if err := c.createTable(ctx, table); err != nil {
				return err
			}
			continue
		}

		changes := schema.GetSchemaChanges(table, schemaTable)
		if c.canAutoMigrate(changes) {
			c.logger.Info().Str("table", tableName).Msg("Table exists, auto-migrating")
			if err := c.autoMigrateTable(ctx, table, changes); err != nil {
				return err
			}
			continue
		}

		c.logger.Info().Str("table", tableName).Msg("Table exists, force migration required")
		if err := c.recreateTable(ctx, table); err != nil {
			return err
		}
	}

	return nil
}
