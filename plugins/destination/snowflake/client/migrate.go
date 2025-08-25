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
				if err := c.autoMigrateTable(ctx, table, existingTable, changes); err != nil {
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

func (c *Client) autoMigrateTable(ctx context.Context, table, oldTable *schema.Table, changes []schema.TableColumnChange) error {
	tableName := table.Name
	for _, change := range changes {
		var err error

		switch {
		case change.Type == schema.TableColumnChangeTypeMoveToCQOnly:
			// Drop PKs automatically (and add PK constraint to _cq_id in next step)
			if len(oldTable.PrimaryKeys()) > 0 {
				err = c.alterTableDropPK(ctx, tableName)
			}
		case change.Type == schema.TableColumnChangeTypeRemove:
			// Not dropping columns automatically
		case change.Type == schema.TableColumnChangeTypeAdd:
			err = c.addColumn(ctx, tableName, change.Current)
		case change.Type == schema.TableColumnChangeTypeUpdate && change.Current.NotNull != change.Previous.NotNull && !change.Current.NotNull:
			err = c.alterColumnDropNotNull(ctx, tableName, change.Current)
		case change.Type == schema.TableColumnChangeTypeUpdate && change.Current.PrimaryKey != change.Previous.PrimaryKey && change.Current.PrimaryKey:
			err = c.alterColumnAddPK(ctx, tableName, change.Current)
		}

		if err != nil {
			return err
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

func (c *Client) alterColumnAddPK(ctx context.Context, tableName string, column schema.Column) error {
	c.logger.Info().Str("table", tableName).Str("column", column.Name).Msg("Adding PRIMARY KEY to column")
	pkeyName := column.Name + "_pk"

	sql := "alter table identifier(?) add constraint " + sanitizeColumn(pkeyName) + " primary key(" + sanitizeColumn(column.Name) + ")"
	if _, err := c.db.ExecContext(ctx, sql, tableName); err != nil {
		return fmt.Errorf("failed to add PRIMARY KEY for column %s on table %s: %w", column.Name, tableName, err)
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
	// The SDK can detect more granular changes than we can handle
	// We know that when the `TableColumnChangeTypeMoveToCQOnly` is present there will be other changes that were found as well
	// As long as the only change is to remove PK from columns and add it to _cq_id, we can skip handling the changes
	// But we need to make sure there are no other changes
	columnsAddingPK := []string{}
	columnsRemovingPK := []string{}
	cqMigration := false
	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeUpdate:
			if change.Current.Type != change.Previous.Type {
				continue
			}
			if change.Current.PrimaryKey && !change.Previous.PrimaryKey {
				columnsAddingPK = append(columnsAddingPK, change.ColumnName)
			}
			if !change.Current.PrimaryKey && change.Previous.PrimaryKey {
				columnsRemovingPK = append(columnsRemovingPK, change.ColumnName)
			}

		case schema.TableColumnChangeTypeMoveToCQOnly:
			cqMigration = true
		}
	}

	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeRemoveUniqueConstraint:
			continue
		case schema.TableColumnChangeTypeAdd:
			if change.Current.PrimaryKey || change.Current.NotNull {
				return false
			}
		case schema.TableColumnChangeTypeRemove:
			if change.Previous.PrimaryKey || change.Previous.NotNull {
				return false
			}
		case schema.TableColumnChangeTypeMoveToCQOnly:
			continue
		case schema.TableColumnChangeTypeUpdate:
			if cqMigration && ((len(columnsAddingPK) == 1 && columnsAddingPK[0] == schema.CqIDColumn.Name) || containsFold(columnsRemovingPK, change.ColumnName)) {
				// We don't need to handle these changes as they are a part of the CQID migration
				continue
			}
			if change.Previous.NotNull != change.Current.NotNull && !change.Current.NotNull {
				// Support removing NOT NULL constraint
				continue
			}
			if change.Previous.PrimaryKey != change.Current.PrimaryKey && change.Current.PrimaryKey {
				// Support adding PRIMARY KEY constraint. Removal is automatically not supported if it's in the change list (for another column)
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

func containsFold(list []string, s string) bool {
	for _, item := range list {
		if strings.EqualFold(item, s) {
			return true
		}
	}
	return false
}

// getTableChangesCaseInsensitive returns changes between two tables when `t` is the new one and `old` is the old one.
// This is a case-insensitive version of schema.Table.GetChanges to handle Snowflake's case insensitivity/ambiguity.
func getTableChangesCaseInsensitive(t, old *schema.Table) []schema.TableColumnChange {
	var changes []schema.TableColumnChange

	//  Special case: Moving from individual pks to singular PK on _cq_id
	newPks := t.PrimaryKeys()

	if len(newPks) == 1 && strings.EqualFold(newPks[0], schema.CqIDColumn.Name) && !containsFold(old.PrimaryKeys(), schema.CqIDColumn.Name) && len(old.PrimaryKeys()) > 0 {
		changes = append(changes, schema.TableColumnChange{
			Type: schema.TableColumnChangeTypeMoveToCQOnly,
		})
	}
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
