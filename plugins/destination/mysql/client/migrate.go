package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func (c *Client) normalizedTables(tables schema.Tables) schema.Tables {
	var normalized schema.Tables
	for _, table := range tables.FlattenTables() {
		for i := range table.Columns {
			// Since multiple schema types can map to the same MySQL type we need to normalize them to avoid false positives when detecting schema changes
			// This should never fail we convert an internal schema type to an MySQL type and back
			schemaType, _ := SchemaType(table.Name, table.Columns[i].Name, SQLType(table.Columns[i].Type))
			table.Columns[i].Type = schemaType
		}
		// If there are no PKs, we use CqID as PK
		pks := table.PrimaryKeys()
		if !c.pkEnabled() || len(pks) == 0 {
			table.Columns.Get(schema.CqIDColumn.Name).CreationOptions.PrimaryKey = true
		}

		for _, col := range table.Columns {
			col.CreationOptions.NotNull = col.CreationOptions.NotNull || col.CreationOptions.PrimaryKey
		}

		normalized = append(normalized, table)
	}

	return normalized
}

func (c *Client) nonAutoMigrableTables(tables schema.Tables, schemaTables schema.Tables) (names []string, changes [][]schema.TableColumnChange) {
	var tableChanges [][]schema.TableColumnChange
	for _, t := range tables {
		schemaTable := schemaTables.Get(t.Name)
		if schemaTable == nil {
			continue
		}
		changes := t.GetChanges(schemaTable)
		if !c.canAutoMigrate(changes) {
			names = append(names, t.Name)
			tableChanges = append(tableChanges, changes)
		}
	}
	return names, tableChanges
}

func (*Client) canAutoMigrate(changes []schema.TableColumnChange) bool {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd && (change.Current.CreationOptions.PrimaryKey || change.Current.CreationOptions.NotNull) {
			return false
		}

		if change.Type == schema.TableColumnChangeTypeRemove && (change.Previous.CreationOptions.PrimaryKey || change.Previous.CreationOptions.NotNull) {
			return false
		}

		if change.Type == schema.TableColumnChangeTypeUpdate {
			return false
		}
	}
	return true
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table, changes []schema.TableColumnChange) error {
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
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	schemaTables, err := c.schemaTables(ctx, tables)
	if err != nil {
		return err
	}

	normalizedTables := c.normalizedTables(tables)

	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonAutoMigrableTables, changes := c.nonAutoMigrableTables(normalizedTables, schemaTables)
		if len(nonAutoMigrableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonAutoMigrableTables, ","), changes)
		}
	}

	for _, table := range normalizedTables {
		c.logger.Info().Str("table", table.Name).Msg("Migrating table")
		schemaTable := schemaTables.Get(table.Name)
		if schemaTable == nil {
			c.logger.Info().Str("table", table.Name).Msg("Table doesn't exist, creating")
			if err := c.createTable(ctx, table); err != nil {
				return err
			}
			continue
		}

		changes := table.GetChanges(schemaTable)
		if c.canAutoMigrate(changes) {
			c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")
			if err := c.autoMigrateTable(ctx, table, changes); err != nil {
				return err
			}
			continue
		}

		c.logger.Info().Str("table", table.Name).Msg("Table exists, force migration required")
		if err := c.recreateTable(ctx, table); err != nil {
			return err
		}
	}

	return nil
}
