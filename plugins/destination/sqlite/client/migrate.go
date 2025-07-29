package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	sqlTableInfo = "PRAGMA table_info('%s');"
)

type columnInfo struct {
	index        int
	name         string
	typ          string
	notNull      bool
	defaultValue any
	pk           int
}

type tableInfo struct {
	columns []columnInfo
}

func identifier(str string) string {
	return `"` + str + `"`
}

func (c *Client) sqliteTables(tables schema.Tables) (schema.Tables, error) {
	var schemaTables schema.Tables
	for _, table := range tables {
		var columns []schema.Column
		info, err := c.getTableInfo(table.Name)
		if info == nil {
			continue
		}
		if err != nil {
			return nil, err
		}
		for _, col := range info.columns {
			columns = append(columns, schema.Column{
				Name:       col.name,
				Type:       c.sqliteTypeToArrowType(col.typ),
				PrimaryKey: col.pk != 0,
				NotNull:    col.notNull,
			})
		}
		schemaTables = append(schemaTables, &schema.Table{Name: table.Name, Columns: columns})
	}
	return schemaTables, nil
}

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
		normalized := c.normalizeField(col.ToArrowField())
		columns[i] = schema.NewColumnFromArrowField(*normalized)
	}
	return &schema.Table{Name: table.Name, Columns: columns}
}

func (c *Client) normalizeField(field arrow.Field) *arrow.Field {
	return &arrow.Field{
		Name:     field.Name,
		Type:     c.arrowTypeToSqlite(field.Type),
		Nullable: field.Nullable,
		Metadata: field.Metadata,
	}
}

func (c *Client) nonAutoMigratableTables(tables schema.Tables, sqliteTables schema.Tables, safeTables map[string]bool) map[string][]schema.TableColumnChange {
	result := make(map[string][]schema.TableColumnChange)
	for _, t := range tables {
		sqliteTable := sqliteTables.Get(t.Name)
		if sqliteTable == nil {
			continue
		}
		changes := sqliteTable.GetChanges(t)
		if safeTables[t.Name] && !c.canAutoMigrate(changes) {
			result[t.Name] = changes
		}
	}
	return result
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table, changes []schema.TableColumnChange) error {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd {
			if err := c.addColumn(ctx, table.Name, change.Current.Name, c.arrowTypeToSqliteStr(change.Current.Type)); err != nil {
				return err
			}
		}
	}
	return nil
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
		case schema.TableColumnChangeTypeRemoveUniqueConstraint:
			continue
		default:
			return false
		}
	}
	return true
}

// This is the responsibility of the CLI of the client to lock before running migration
func (c *Client) MigrateTables(ctx context.Context, msgs message.WriteMigrateTables) error {
	allTables := schema.Tables{}
	safeTables := make(map[string]bool)
	for _, msg := range msgs {
		allTables = append(allTables, msg.Table)
		safeTables[msg.Table.Name] = !msg.MigrateForce
	}

	normalizedTables := c.normalizeTables(allTables)
	sqliteTables, err := c.sqliteTables(normalizedTables)
	if err != nil {
		return err
	}

	nonAutoMigratableTables := c.nonAutoMigratableTables(normalizedTables, sqliteTables, safeTables)
	if len(nonAutoMigratableTables) > 0 {
		return fmt.Errorf("\nCan't migrate tables automatically, migrate manually or consider using 'migrate_mode: forced'. Non auto migratable tables changes:\n\n%s", schema.GetChangesSummary(nonAutoMigratableTables))
	}

	for _, table := range normalizedTables {
		c.logger.Info().Str("table", table.Name).Msg("Migrating table")
		if len(table.Columns) == 0 {
			c.logger.Info().Str("table", table.Name).Msg("Table with no columns, skipping")
			continue
		}

		sqlite := sqliteTables.Get(table.Name)
		if sqlite == nil {
			c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(table); err != nil {
				return err
			}
		} else {
			changes := table.GetChanges(sqlite)
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

func (c *Client) recreateTable(ctx context.Context, table *schema.Table) error {
	sql := "drop table if exists " + identifier(table.Name)
	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to drop table %s: %w", table.Name, err)
	}
	return c.createTableIfNotExist(table)
}

func (c *Client) addColumn(ctx context.Context, tableName string, columnName string, columnType string) error {
	sql := "alter table " + identifier(tableName) + " add column " + identifier(columnName) + " " + identifier(columnType)
	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to add column %s on table %s: %w", columnName, tableName, err)
	}
	return nil
}

func (c *Client) createTableIfNotExist(table *schema.Table) error {
	var sb strings.Builder

	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(identifier(table.Name))
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	primaryKeys := []string{}
	for i, col := range table.Columns {
		sqlType := c.arrowTypeToSqliteStr(col.Type)
		if sqlType == "" {
			c.logger.Warn().Str("table", table.Name).Str("column", col.Name).Msg("Column type is not supported, skipping")
			continue
		}
		fieldDef := identifier(col.Name) + ` ` + sqlType
		if col.NotNull {
			fieldDef += " NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
		if col.PrimaryKey {
			primaryKeys = append(primaryKeys, identifier(col.Name))
		}
	}

	if len(primaryKeys) > 0 {
		// add composite PK constraint on primary key columns
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(identifier(table.Name + "_cqpk"))
		sb.WriteString(" PRIMARY KEY (")
		sb.WriteString(strings.Join(primaryKeys, ","))
		sb.WriteString(")")
	}
	sb.WriteString(")")
	_, err := c.db.Exec(sb.String())
	if err != nil {
		return fmt.Errorf("failed to create table with '%s': %w", sb.String(), err)
	}
	return nil
}

func (c *Client) getTableInfo(tableName string) (*tableInfo, error) {
	info := tableInfo{}
	rows, err := c.db.Query(fmt.Sprintf(sqlTableInfo, tableName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		colInfo := columnInfo{}
		if err := rows.Scan(
			&colInfo.index,
			&colInfo.name,
			&colInfo.typ,
			&colInfo.notNull,
			&colInfo.defaultValue,
			&colInfo.pk); err != nil {
			return nil, err
		}
		colInfo.typ = strings.ToLower(colInfo.typ)
		info.columns = append(info.columns, colInfo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(info.columns) == 0 {
		// Table doesn't exist
		return nil, nil
	}
	return &info, nil
}
