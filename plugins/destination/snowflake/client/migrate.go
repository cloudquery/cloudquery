package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"golang.org/x/sync/errgroup"
)

func (c *Client) MigrateTables(ctx context.Context, msgs message.WriteMigrateTables) error {
	safeTables := make(map[string]bool, len(msgs))
	tables := make(schema.Tables, 0, len(msgs))
	for _, msg := range msgs {
		// Switch to effective types
		for i, col := range msg.Table.Columns {
			msg.Table.Columns[i].Type = SnowflakeToSchemaType(SchemaTypeToSnowflake(col.Type))
		}

		tables = append(tables, msg.Table)
		safeTables[msg.Table.Name] = !msg.MigrateForce
	}

	existingTables, err := c.getTableInfo(ctx, tables.TableNames())
	if err != nil {
		return fmt.Errorf("failed to get list of tables: %w", err)
	}

	nonAutoMigratableTables := c.nonAutoMigratableTables(tables, existingTables, safeTables)
	if len(nonAutoMigratableTables) > 0 {
		return fmt.Errorf("\nCan't migrate tables automatically, migrate manually or consider using 'migrate_mode: forced'. Non auto migratable tables changes:\n\n%s", schema.GetChangesSummary(nonAutoMigratableTables))
	}

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(c.spec.MigrateConcurrency)

	for _, table := range tables {
		table := table

		g.Go(func() error {
			c.logger.Info().Str("table", table.Name).Msg("Migrating table")
			if len(table.Columns) == 0 {
				c.logger.Info().Str("table", table.Name).Msg("Table with no columns, skipping")
				return nil
			}
			existingTable := existingTables.Get(strings.ToUpper(table.Name))
			if existingTable == nil {
				c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
				return c.createTableIfNotExist(ctx, table)
			}

			changes := getTableChangesCaseInsensitive(table, existingTable)
			if c.canAutoMigrate(changes) {
				c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(ctx, table, changes); err != nil {
					return err
				}
			} else {
				c.logger.Info().Str("table", table.Name).Msg("Table exists, force migration required")
				if err := c.dropTable(ctx, table.Name); err != nil {
					return err
				}
				if err := c.createTableIfNotExist(ctx, table); err != nil {
					return err
				}
			}
			return nil
		})
	}
	return g.Wait()
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table, changes []schema.TableColumnChange) error {
	// Special PK handling: Drop PK first, add all PKs together after
	{
		var dropPKs, addPKs bool
		for _, change := range changes {
			switch {
			case change.Type == schema.TableColumnChangeTypeUpdate && change.Previous.PrimaryKey && !change.Current.PrimaryKey:
				dropPKs = true
			case change.Type == schema.TableColumnChangeTypeRemove && change.Previous.PrimaryKey:
				dropPKs = true
			case change.Type == schema.TableColumnChangeTypeUpdate && !change.Previous.PrimaryKey && change.Current.PrimaryKey:
				addPKs = true
			case change.Type == schema.TableColumnChangeTypeAdd && change.Current.PrimaryKey:
				addPKs = true
			}
		}

		if dropPKs {
			if err := c.alterTableDropPK(ctx, table.Name); err != nil {
				return err
			}
		}
		if addPKs { // PKs are being added, so let's add all of them
			pkCols := []string{}
			for _, col := range table.Columns {
				if col.PrimaryKey {
					pkCols = append(pkCols, col.Name)
				}
			}
			if len(pkCols) > 0 {
				if err := c.alterTableAddPK(ctx, table.Name, pkCols); err != nil {
					return err
				}
			}
		}
	}

	for _, change := range changes {
		//exhaustive:enforce
		switch change.Type {
		case schema.TableColumnChangeTypeUnknown:
			continue
		case schema.TableColumnChangeTypeMoveToCQOnly:
			// No need to handle specifically, will be handled by drop PK / add PK changes
			continue
		case schema.TableColumnChangeTypeRemoveUniqueConstraint:
			// Can't happen as canAutoMigrate would return false
			continue
		case schema.TableColumnChangeTypeRemove:
			// Not dropping columns automatically: Keep them
			continue
		case schema.TableColumnChangeTypeAdd:
			if err := c.addColumn(ctx, table.Name, change.Current); err != nil {
				return err
			}
		case schema.TableColumnChangeTypeUpdate:
			if change.Current.NotNull != change.Previous.NotNull && !change.Current.NotNull {
				if err := c.alterColumnDropNotNull(ctx, table.Name, change.Current); err != nil {
					return err
				}
			}
			// PK changes are already handled above
		}
	}
	return nil
}

func (c *Client) addColumn(ctx context.Context, tableName string, column schema.Column) error {
	c.logger.Info().Str("table", tableName).Str("column", column.Name).Msg("Column doesn't exist, creating")
	columnType := SchemaTypeToSnowflake(column.Type)
	sql := "alter table identifier(?) add column " + sanitizeColumn(column.Name) + " " + columnType
	if column.Unique {
		sql += " unique"
	}
	if column.NotNull {
		sql += " not null"
	}

	if _, err := c.db.ExecContext(ctx, sql, tableName); err != nil {
		return fmt.Errorf("failed to add column %s on table %s: %w", column.Name, tableName, err)
	}
	return nil
}

func (c *Client) alterColumnDropNotNull(ctx context.Context, tableName string, column schema.Column) error {
	c.logger.Info().Str("table", tableName).Str("column", column.Name).Msg("Dropping NOT NULL from column")
	sql := "alter table identifier(?) alter column " + sanitizeColumn(column.Name) + " drop not null"
	if _, err := c.db.ExecContext(ctx, sql, tableName); err != nil {
		return fmt.Errorf("failed to drop NOT NULL for column %s on table %s: %w", column.Name, tableName, err)
	}
	return nil
}

func (c *Client) alterTableAddPK(ctx context.Context, tableName string, columnNames []string) error {
	c.logger.Info().Str("table", tableName).Strs("columns", columnNames).Msg("Adding PRIMARY KEY to column")
	pkeyName := tableName + "_pk"

	sanCols := make([]string, len(columnNames))
	for i, col := range columnNames {
		sanCols[i] = sanitizeColumn(col)
	}

	sql := "alter table identifier(?) add constraint " + sanitizeColumn(pkeyName) + " primary key(" + strings.Join(sanCols, ",") + ")"
	if _, err := c.db.ExecContext(ctx, sql, tableName); err != nil {
		return fmt.Errorf("failed to add PRIMARY KEY for column %s on table %s: %w", strings.Join(columnNames, ","), tableName, err)
	}
	return nil
}

func (c *Client) alterTableDropPK(ctx context.Context, tableName string) error {
	c.logger.Info().Str("table", tableName).Msg("Dropping PRIMARY KEY from table")

	if _, err := c.db.ExecContext(ctx, "alter table identifier(?) drop primary key", tableName); err != nil {
		return fmt.Errorf("failed to drop PRIMARY KEY on table %s: %w", tableName, err)
	}
	return nil
}

func (c *Client) createTableIfNotExist(ctx context.Context, table *schema.Table) error {
	var sb strings.Builder
	sb.WriteString("create table if not exists identifier(?) (")
	totalColumns := len(table.Columns)

	primaryKeys := []string{}
	for i, col := range table.Columns {
		sfType := SchemaTypeToSnowflake(col.Type)
		fieldDef := sanitizeColumn(col.Name) + " " + sfType
		if col.Unique {
			fieldDef += " unique"
		}
		if col.NotNull {
			fieldDef += " not null"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
		if col.PrimaryKey {
			primaryKeys = append(primaryKeys, col.Name)
		}
	}

	if len(primaryKeys) > 0 {
		// add composite PK constraint on primary key columns
		sb.WriteString(", primary key (")
		sb.WriteString(strings.Join(primaryKeys, ","))
		sb.WriteString(")")
	}
	sb.WriteString(")")

	_, err := c.db.ExecContext(ctx, sb.String(), table.Name)
	if err != nil {
		c.logger.Error().Err(err).Str("table", table.Name).Str("query", sb.String()).Msg("Failed to create table")
		return fmt.Errorf("failed to create table %s: %w"+sb.String(), table.Name, err)
	}

	return nil
}

func (c *Client) dropTable(ctx context.Context, tableName string) error {
	c.logger.Info().Str("table", tableName).Msg("Dropping table")
	if _, err := c.db.ExecContext(ctx, "drop table identifier(?)", tableName); err != nil {
		return fmt.Errorf("failed to drop table %s: %w", tableName, err)
	}
	return nil
}

func (c *Client) nonAutoMigratableTables(tables schema.Tables, existingTables schema.Tables, safeTables map[string]bool) map[string][]schema.TableColumnChange {
	result := make(map[string][]schema.TableColumnChange)
	for _, t := range tables {
		existingTable := existingTables.Get(strings.ToUpper(t.Name))
		if existingTable == nil {
			continue
		}
		changes := getTableChangesCaseInsensitive(t, existingTable)
		if safeTables[t.Name] && !c.canAutoMigrate(changes) {
			result[t.Name] = changes
		}
	}
	return result
}

func (*Client) canAutoMigrate(changes []schema.TableColumnChange) bool {
	for _, change := range changes {
		//exhaustive:enforce
		switch change.Type {
		case schema.TableColumnChangeTypeRemoveUniqueConstraint:
			return false
		case schema.TableColumnChangeTypeAdd:
			if change.Current.NotNull {
				return false
			}
		case schema.TableColumnChangeTypeRemove:
			if change.Previous.NotNull {
				return false
			}
		case schema.TableColumnChangeTypeMoveToCQOnly:
			continue
		case schema.TableColumnChangeTypeUpdate:
			if change.Previous.NotNull != change.Current.NotNull && !change.Current.NotNull {
				// Support removing NOT NULL constraint
				continue
			}
			if change.Previous.PrimaryKey != change.Current.PrimaryKey {
				// Support adding or removing PRIMARY KEY constraints
				continue
			}

			return false
		default:
			return false
		}
	}
	return true
}

func sanitizeColumn(name string) string {
	// temporary, `identifier()` would be better but it doesn't work for column names
	return `"` + strings.ToUpper(name) + `"`
}

// getTableChangesCaseInsensitive returns changes between two tables when `t` is the new one and `old` is the old one.
// This is a case-insensitive version of schema.Table.GetChanges to handle Snowflake's case insensitivity/ambiguity.
func getTableChangesCaseInsensitive(t, old *schema.Table) []schema.TableColumnChange {
	var changes []schema.TableColumnChange

	for _, c := range t.Columns {
		otherColumn := findColumn(old.Columns, c.Name)
		// A column was added to the table definition
		if otherColumn == nil {
			changes = append(changes, schema.TableColumnChange{
				Type:       schema.TableColumnChangeTypeAdd,
				ColumnName: c.Name,
				Current:    c,
			})
			continue
		}

		// Column type or options (e.g. PK, Not Null) changed in the new table definition
		if !arrow.TypeEqual(c.Type, otherColumn.Type) || c.NotNull != otherColumn.NotNull || c.PrimaryKey != otherColumn.PrimaryKey {
			changes = append(changes, schema.TableColumnChange{
				Type:       schema.TableColumnChangeTypeUpdate,
				ColumnName: c.Name,
				Current:    c,
				Previous:   *otherColumn,
			})
		}

		// Unique constraint was removed
		if !c.Unique && otherColumn.Unique {
			changes = append(changes, schema.TableColumnChange{
				Type:       schema.TableColumnChangeTypeRemoveUniqueConstraint,
				ColumnName: c.Name,
				Previous:   *otherColumn,
			})
		}
	}
	// A column was removed from the table definition
	for _, c := range old.Columns {
		if findColumn(t.Columns, c.Name) == nil {
			changes = append(changes, schema.TableColumnChange{
				Type:       schema.TableColumnChangeTypeRemove,
				ColumnName: c.Name,
				Previous:   c,
			})
		}
	}
	return changes
}

// findColumn is the case-insensitive version of schema.ColumnList.Get
func findColumn(c schema.ColumnList, name string) *schema.Column {
	for i := range c {
		if strings.EqualFold(c[i].Name, name) {
			return &c[i]
		}
	}
	return nil
}
