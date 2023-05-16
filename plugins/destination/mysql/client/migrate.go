package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) normalizeTables(tables schema.Tables) (schema.Tables, error) {
	flattened := tables.FlattenTables()
	normalized := make(schema.Tables, len(flattened))
	var err error
	for i, table := range flattened {
		normalized[i], err = c.normalizeTable(table)
		if err != nil {
			return nil, err
		}
	}
	return normalized, nil
}

func (c *Client) normalizeTable(table *schema.Table) (*schema.Table, error) {
	columns := make([]schema.Column, len(table.Columns))
	for i, col := range table.Columns {
		normalized, err := c.normalizeField(col.ToArrowField())
		if err != nil {
			return nil, err
		}
		columns[i] = schema.NewColumnFromArrowField(*normalized)
	}
	return &schema.Table{Name: table.Name, Columns: columns}, nil
}

func (c *Client) normalizeField(field arrow.Field) (*arrow.Field, error) {
	normalizedType, err := mySQLTypeToArrowType("", "", arrowTypeToMySqlStr(field.Type))
	if err != nil {
		return nil, err
	}
	return &arrow.Field{
		Name:     field.Name,
		Type:     normalizedType,
		Nullable: field.Nullable,
		Metadata: field.Metadata,
	}, nil
}

func (c *Client) nonAutoMigratableTables(tables schema.Tables, mysqlTables schema.Tables) ([]string, [][]schema.TableColumnChange) {
	var result []string
	var tableChanges [][]schema.TableColumnChange
	for _, t := range tables {
		sqliteTable := mysqlTables.Get(t.Name)
		if sqliteTable == nil {
			continue
		}
		changes := sqliteTable.GetChanges(t)
		if !c.canAutoMigrate(changes) {
			result = append(result, t.Name)
			tableChanges = append(tableChanges, changes)
		}
	}
	return result, tableChanges
}

func (*Client) canAutoMigrate(changes []schema.TableColumnChange) bool {
	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeAdd:
			if change.Current.PrimaryKey || change.Current.NotNull {
				return false
			}
		case schema.TableColumnChangeTypeRemove:
			if change.Previous.PrimaryKey || change.Previous.NotNull {
				return false
			}
		case schema.TableColumnChangeTypeUpdate:
			return false
		default:
			panic("unknown change type")
		}
	}
	return true
}
func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table, changes []schema.TableColumnChange) error {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd {
			if err := c.addColumn(ctx, table, table.Columns.Get(change.ColumnName)); err != nil {
				return err
			}
		}
	}
	return nil
}

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	mysqlTables, err := c.schemaTables(ctx, tables)
	if err != nil {
		return err
	}

	normalizedTables, err := c.normalizeTables(tables)
	if err != nil {
		return err
	}

	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonAutoMigrtableTables, changes := c.nonAutoMigratableTables(normalizedTables, mysqlTables)
		if len(nonAutoMigrtableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonAutoMigrtableTables, ","), changes)
		}
	}

	for _, table := range normalizedTables {
		c.logger.Info().Str("table", table.Name).Msg("Migrating table")
		if len(table.Columns) == 0 {
			c.logger.Info().Str("table", table.Name).Msg("Table with no columns, skipping")
			continue
		}

		mysql := mysqlTables.Get(table.Name)
		if mysql == nil {
			c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
			if err := c.createTable(ctx, table); err != nil {
				return err
			}
		} else {
			changes := table.GetChanges(mysql)
			if c.canAutoMigrate(changes) {
				c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(ctx, table, changes); err != nil {
					return err
				}
			} else {
				c.logger.Info().Str("table", table.Name).Msg("Table exists, force migration required")
				if err := c.recreateTable(ctx, table); err != nil {
					return err
				}
			}
		}

	}

	return nil
}
