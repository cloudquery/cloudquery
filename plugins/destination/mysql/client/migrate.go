package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) normalizeTables(tables schema.Tables) schema.Tables {
	flattened := tables.FlattenTables()
	normalized := make(schema.Tables, len(flattened))
	for i, table := range flattened {
		normalized[i] = c.normalizeTable(table)
	}
	return normalized
}

func (c *Client) normalizeTable(table *schema.Table) *schema.Table {
	columns := make([]schema.Column, len(table.Columns))
	for i, col := range table.Columns {
		columns[i] = c.normalizeColumn(col)
	}
	return &schema.Table{Name: table.Name, Columns: columns}
}

func (*Client) normalizeColumn(col schema.Column) schema.Column {
	field := col.ToArrowField()
	normalizedType := mySQLTypeToArrowType(arrowTypeToMySqlStr(field.Type))
	// In MySQL primary keys are implicitly not null
	notNull := col.NotNull || col.PrimaryKey
	return schema.NewColumnFromArrowField(arrow.Field{
		Name:     field.Name,
		Type:     normalizedType,
		Nullable: !notNull,
		Metadata: field.Metadata,
	})
}

func (c *Client) nonAutoMigratableTables(tables schema.Tables, mysqlTables schema.Tables) ([]string, [][]schema.TableColumnChange) {
	var result []string
	var tableChanges [][]schema.TableColumnChange
	for _, t := range tables {
		mysqlTable := mysqlTables.Get(t.Name)
		if mysqlTable == nil {
			continue
		}
		changes := mysqlTable.GetChanges(t)
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

func getTables(msgs message.WriteMigrateTables) schema.Tables {
	tables := make(schema.Tables, len(msgs))
	for i, msg := range msgs {
		tables[i] = msg.Table
	}
	return tables
}

// Migrate relies on the CLI/client to lock before running migration.
func (c *Client) MigrateTables(ctx context.Context, msgs message.WriteMigrateTables) error {
	tables := getTables(msgs)
	mysqlTables, err := c.schemaTables(ctx, tables)
	if err != nil {
		return err
	}

	normalizedTables := c.normalizeTables(tables)
	normalizedTablesSafeMode := make(schema.Tables, 0, len(normalizedTables))
	for _, table := range normalizedTables {
		msg := msgs.GetMessageByTable(table.Name)
		if msg == nil {
			continue
		}
		if !msg.MigrateForce {
			normalizedTablesSafeMode = append(normalizedTablesSafeMode, table)
		}
	}

	nonAutoMigrtableTables, changes := c.nonAutoMigratableTables(normalizedTablesSafeMode, mysqlTables)
	if len(nonAutoMigrtableTables) > 0 {
		return fmt.Errorf("tables %s with changes %v require migration. Migrate manually or consider using 'migrate_mode: forced'", strings.Join(nonAutoMigrtableTables, ","), changes)
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
