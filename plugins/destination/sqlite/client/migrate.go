package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/schema"
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

func (c *Client) normalizeTables(tables schema.Tables) (schema.Tables, error) {
	flattened := tables.FlattenTables()
	canonized := make(schema.Tables, len(flattened))
	var err error
	for i, table := range flattened {
		canonized[i], err = c.normalizeTable(table)
		if err != nil {
			return nil, err
		}
	}
	return canonized, nil
}

func (c *Client) normalizeTable(table *schema.Table) (*schema.Table, error) {
	columns := make([]schema.Column, len(table.Columns))
	for i, col := range table.Columns {
		canonized := c.normalizeField(col.ToArrowField())
		columns[i] = schema.NewColumnFromArrowField(*canonized)
	}
	return &schema.Table{Name: table.Name, Columns: columns}, nil
}

func (c *Client) normalizeField(field arrow.Field) *arrow.Field {
	// 1 - convert to the ClickHouse
	fieldType := c.arrowTypeToSqlite(field.Type)

	// 2 - convert back to Apache Arrow
	return &arrow.Field{
		Name:     field.Name,
		Type:     fieldType,
		Nullable: field.Nullable,
		Metadata: field.Metadata,
	}
}

func (c *Client) nonAutoMigratableTables(tables schema.Tables, sqliteTables schema.Tables) ([]string, [][]schema.TableColumnChange) {
	var result []string
	var tableChanges [][]schema.TableColumnChange
	for _, t := range tables {

		sqliteTable := sqliteTables.Get(t.Name)
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

func (c *Client) autoMigrateTable(table *schema.Table, changes []schema.TableColumnChange) error {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd {
			if err := c.addColumn(table.Name, change.Current.Name, c.arrowTypeToSqliteStr(change.Current.Type)); err != nil {
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
		case schema.TableColumnChangeTypeUpdate:
			return false
		default:
			panic("unknown change type")
		}
	}
	return true
}

// This is the responsibility of the CLI of the client to lock before running migration
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	normalizedTables, err := c.normalizeTables(tables)
	if err != nil {
		return err
	}

	sqliteTables, err := c.sqliteTables(normalizedTables)
	if err != nil {
		return err
	}

	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonAutoMigratableTables, changes := c.nonAutoMigratableTables(normalizedTables, sqliteTables)
		if len(nonAutoMigratableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonAutoMigratableTables, ","), changes)
		}
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
				if err := c.autoMigrateTable(table, changes); err != nil {
					return err
				}
			} else {
				c.logger.Info().Str("table", table.Name).Msg("Table exists, force migration required")
				if err := c.recreateTable(table); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (c *Client) recreateTable(table *schema.Table) error {

	sql := "drop table if exists \"" + table.Name + "\""
	if _, err := c.db.Exec(sql); err != nil {
		return fmt.Errorf("failed to drop table %s: %w", table.Name, err)
	}
	return c.createTableIfNotExist(table)
}

func (c *Client) addColumn(tableName string, columnName string, columnType string) error {
	sql := "alter table \"" + tableName + "\" add column \"" + columnName + "\" \"" + columnType + `"`
	if _, err := c.db.Exec(sql); err != nil {
		return fmt.Errorf("failed to add column %s on table %s: %w", columnName, tableName, err)
	}
	return nil
}

func (c *Client) createTableIfNotExist(table *schema.Table) error {
	var sb strings.Builder

	// TODO sanitize table.Name
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(`"` + table.Name + `"`)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	primaryKeys := []string{}
	for i, col := range table.Columns {
		sqlType := c.arrowTypeToSqliteStr(col.Type)
		if sqlType == "" {
			c.logger.Warn().Str("table", table.Name).Str("column", col.Name).Msg("Column type is not supported, skipping")
			continue
		}
		// TODO: sanitize column name
		fieldDef := `"` + col.Name + `" ` + sqlType
		if col.NotNull {
			fieldDef += " NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}

		if c.enabledPks() && col.PrimaryKey {
			primaryKeys = append(primaryKeys, `"`+col.Name+`"`)
		}
	}

	if len(primaryKeys) > 0 {
		// add composite PK constraint on primary key columns
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(table.Name)
		sb.WriteString("_cqpk PRIMARY KEY (")
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

func (c *Client) enabledPks() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}
