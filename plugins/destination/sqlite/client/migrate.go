package client

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

const (
	isTableExistSQL = "SELECT count(name) FROM sqlite_master WHERE type='table' AND name=?;"
	sqlTableInfo    = "PRAGMA table_info('%s');"
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

func (i *tableInfo) getColumn(name string) *columnInfo {
	for _, col := range i.columns {
		if col.name == name {
			return &col
		}
	}
	return nil
}

type columnChange struct {
	name    string
	oldType string
	newType string
	new     bool
	oldPk   bool
	newPk   bool
}

func (c *columnChange) isNonPKTypeChange() bool {
	return !c.new && c.oldType != c.newType && !c.newPk
}

func (c *columnChange) isPKTypeChange() bool {
	return !c.new && c.oldType != c.newType && c.newPk
}

func (c *columnChange) isNewPKColumn() bool {
	return c.new && c.newPk
}

func (c *columnChange) isPKAddToExistingColumn() bool {
	return !c.new && !c.oldPk && c.newPk
}

func (c *columnChange) isPKRemoveFromExistingColumn() bool {
	return !c.new && c.oldPk && !c.newPk
}

func (c *columnChange) needsTableDrop() bool {
	return c.isNewPKColumn() || c.isPKAddToExistingColumn() || c.isPKRemoveFromExistingColumn() || c.isPKTypeChange()
}

func (c *columnChange) isInternal() bool {
	return strings.HasPrefix(c.name, "_cq_")
}

type tableChange struct {
	table         *schema.Table
	new           bool
	columnChanges []*columnChange
}

type migrationMessage struct {
	err  string
	info string
}

type migrationsMessages []migrationMessage

func (m migrationsMessages) Errors() []string {
	errs := make([]string, 0, len(m))
	for _, msg := range m {
		errs = append(errs, msg.err)
	}
	return errs
}

func (m migrationsMessages) Infos() []string {
	infos := make([]string, 0, len(m))
	for _, msg := range m {
		infos = append(infos, msg.info)
	}
	return infos
}

func (c *Client) getColumnChange(col schema.Column, sqliteColumn *columnInfo) *columnChange {
	columnName := col.Name
	columnType := c.SchemaTypeToSqlite(col.Type)

	if sqliteColumn == nil {
		return &columnChange{name: columnName, oldType: columnType, newType: columnType, new: true, oldPk: c.enabledPks() && col.CreationOptions.PrimaryKey, newPk: c.enabledPks() && col.CreationOptions.PrimaryKey}
	}

	return &columnChange{name: columnName, oldType: sqliteColumn.typ, newType: columnType, oldPk: c.enabledPks() && sqliteColumn.pk != 0, newPk: c.enabledPks() && col.CreationOptions.PrimaryKey}
}

func (c *Client) getColumnChanges(table *schema.Table) ([]*columnChange, error) {
	var err error
	var info *tableInfo
	if info, err = c.getTableInfo(table.Name); err != nil {
		return nil, fmt.Errorf("failed to get table %s columns types: %w", table.Name, err)
	}

	columnChanges := make([]*columnChange, len(table.Columns))
	for i, col := range table.Columns {
		columnChanges[i] = c.getColumnChange(col, info.getColumn(col.Name))
	}

	sort.SliceStable(columnChanges, func(i, j int) bool {
		change1 := columnChanges[i]
		change2 := columnChanges[j]

		// Changes that require dropping the table comes first
		if change1.needsTableDrop() && !(change2.needsTableDrop()) {
			return true
		}

		if !(change1.needsTableDrop()) && change2.needsTableDrop() {
			return false
		}

		// Internal columns come last
		if change1.isInternal() && !change2.isInternal() {
			return false
		}

		if !change1.isInternal() && change2.isInternal() {
			return true
		}

		return change1.name < change2.name
	})

	return columnChanges, nil
}

func (c *Client) getTableChange(table *schema.Table) (*tableChange, error) {
	tableExist, err := c.isTableExistSQL(table.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to check if table %s exists: %w", table.Name, err)
	}
	tableChange := &tableChange{table: table, new: !tableExist}
	if tableExist {
		columnChanges, err := c.getColumnChanges(table)
		if err != nil {
			return nil, err
		}
		tableChange.columnChanges = columnChanges
	}
	return tableChange, nil
}

func (c *Client) getSchemaChanges(tables schema.Tables) ([]*tableChange, error) {
	changes := make([]*tableChange, 0, len(tables))
	for _, table := range tables {
		tableChange, err := c.getTableChange(table)
		if err != nil {
			return nil, err
		}
		changes = append(changes, tableChange)
		relationChanges, err := c.getSchemaChanges(table.Relations)
		if err != nil {
			return nil, err
		}
		changes = append(changes, relationChanges...)
	}
	return changes, nil
}

func getMigrationMessages(changes []*tableChange) migrationsMessages {
	var messages migrationsMessages
	for _, tableChange := range changes {
		if tableChange.new {
			continue
		}
		for _, colChange := range tableChange.columnChanges {
			if colChange.isNewPKColumn() {
				messages = append(messages, migrationMessage{
					err:  fmt.Sprintf("can't migrate table %q since adding the new PK column %q is not supported. Try dropping this table", tableChange.table.Name, colChange.name),
					info: fmt.Sprintf("table %q will be dropped and recreated due to adding %q as a PK", tableChange.table.Name, colChange.name),
				})
				// no need to report other errors as the user needs to drop the table altogether
				break
			}
			if colChange.isPKAddToExistingColumn() {
				messages = append(messages, migrationMessage{
					err:  fmt.Sprintf("can't migrate table %q since making the existing column %q a PK is not supported. Try dropping this table", tableChange.table.Name, colChange.name),
					info: fmt.Sprintf("table %q will be dropped and recreated due to adding %q as a PK", tableChange.table.Name, colChange.name),
				})
				// no need to report other errors as the user needs to drop the table altogether
				break
			}
			if colChange.isPKRemoveFromExistingColumn() {
				messages = append(messages, migrationMessage{
					err:  fmt.Sprintf("can't migrate table %q since removing an existing column %q as a PK is not supported. Try dropping this table", tableChange.table.Name, colChange.name),
					info: fmt.Sprintf("table %q will be dropped and recreated due to removing %q as a PK", tableChange.table.Name, colChange.name),
				})
				// no need to report other errors as the user needs to drop the table altogether
				break
			}
			if colChange.isPKTypeChange() {
				messages = append(messages, migrationMessage{
					err:  fmt.Sprintf("can't migrate table %q since changing the type of the PK column %q from %q to %q is not supported. Try dropping this table", tableChange.table.Name, colChange.name, colChange.oldType, colChange.newType),
					info: fmt.Sprintf("table %q will be dropped and recreated due to the type change of the PK column %q", tableChange.table.Name, colChange.name),
				})
				// no need to report other errors as the user needs to drop the table altogether
				break
			}
			if colChange.isNonPKTypeChange() {
				messages = append(messages, migrationMessage{
					err:  fmt.Sprintf("can't migrate table %q since changing the type of column %q from %q to %q is not supported. Try dropping this column for this table", tableChange.table.Name, colChange.name, colChange.oldType, colChange.newType),
					info: fmt.Sprintf("column %q of table %q will be dropped and recreated", colChange.name, tableChange.table.Name),
				})
			}
		}
	}
	return messages
}

// This is the responsibility of the CLI of the client to lock before running migration
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	schemaChanges, err := c.getSchemaChanges(tables)
	if err != nil {
		return err
	}

	migrationMessages := getMigrationMessages(schemaChanges)
	if len(migrationMessages) > 0 {
		if c.spec.MigrateMode == specs.MigrateModeSafe {
			return fmt.Errorf("failed to migrate schema:\n%s\n\nTo force a migration add \"migrate_mode: %s\" to your destination spec", strings.Join(migrationMessages.Errors(), "\n"), specs.MigrateModeForced.String())
		}
		for _, msg := range migrationMessages.Infos() {
			c.logger.Info().Msg(msg)
		}
	}

	for _, tableChange := range schemaChanges {
		table := tableChange.table
		c.logger.Debug().Str("table", table.Name).Msg("Migrating table")
		if tableChange.new {
			c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
			err := c.createTableIfNotExist(tableChange.table)
			if err != nil {
				return err
			}
		} else {
			c.logger.Debug().Str("table", table.Name).Msg("Table exists, auto-migrating")
			for _, colChange := range tableChange.columnChanges {
				tableName := tableChange.table.Name
				columnName := colChange.name
				columnType := colChange.newType
				// If this is a new PK column we need to drop the table
				if colChange.isNewPKColumn() {
					c.logger.Debug().Str("table", tableName).Str("column", colChange.name).Msg("New column is a primary key, dropping and adding table since in forced migrate mode")
					err := c.recreateTable(table)
					if err != nil {
						return err
					}
					break
				}
				// SQLite doesn't support PK additions on tables so we need to drop and add the table
				if colChange.isPKAddToExistingColumn() {
					c.logger.Debug().Str("table", table.Name).Str("column", colChange.name).Msg("Primary key added for existing column, dropping and adding table since in forced migrate mode")
					err := c.recreateTable(table)
					if err != nil {
						return err
					}
					break
				}
				// SQLite doesn't support PK removals on tables so we need to drop and add the table
				if colChange.isPKRemoveFromExistingColumn() {
					c.logger.Debug().Str("table", table.Name).Str("column", colChange.name).Msg("Primary key removed for existing column, dropping and adding table since in forced migrate mode")
					err := c.recreateTable(table)
					if err != nil {
						return err
					}
					break
				}
				// Since we can't recreate a PK column in SQLite we need to drop and add the table if a type changed for an existing PK column
				if colChange.isPKTypeChange() {
					c.logger.Debug().Str("table", table.Name).Str("column", colChange.name).Msg("Type changed for existing primary key column, dropping and adding table since in forced migrate mode")
					err := c.recreateTable(table)
					if err != nil {
						return err
					}
					break
				}
				// if this is an existing column with type change we need to drop and add it
				if colChange.isNonPKTypeChange() {
					c.logger.Debug().Str("table", table.Name).Str("column", colChange.name).Msg("Existing column type changed, dropping and adding column since in forced migrate mode")
					err := c.dropColumn(tableName, columnName)
					if err != nil {
						return err
					}
					err = c.addColumn(tableName, columnName, columnType)
					if err != nil {
						return err
					}
					continue
				}
				if colChange.new {
					c.logger.Debug().Str("table", tableName).Str("column", colChange.name).Msg("Column doesn't exist, creating")
					err := c.addColumn(tableName, columnName, columnType)
					if err != nil {
						return err
					}
					continue
				}
			}
		}
	}

	return nil
}

func (c *Client) isTableExistSQL(table string) (bool, error) {
	var tableExist int
	if err := c.db.QueryRow(isTableExistSQL, table).Scan(&tableExist); err != nil {
		return false, fmt.Errorf("failed to check if table %s exists: %w", table, err)
	}
	return tableExist == 1, nil
}

func (c *Client) recreateTable(table *schema.Table) error {
	sql := "drop table if exists \"" + table.Name + "\""
	if _, err := c.db.Exec(sql); err != nil {
		return fmt.Errorf("failed to drop table %s: %w", table.Name, err)
	}
	return c.createTableIfNotExist(table)
}

func (c *Client) dropColumn(tableName string, columnName string) error {
	sql := "alter table " + tableName + " drop column " + columnName
	if _, err := c.db.Exec(sql); err != nil {
		return fmt.Errorf("failed to drop column %s on table %s: %w", columnName, tableName, err)
	}
	return nil
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
	// TODO sanitize tablename
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(`"` + table.Name + `"`)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	primaryKeys := []string{}
	for i, col := range table.Columns {
		sqlType := c.SchemaTypeToSqlite(col.Type)
		if sqlType == "" {
			c.logger.Warn().Str("table", table.Name).Str("column", col.Name).Msg("Column type is not supported, skipping")
			continue
		}
		// TODO: sanitize column name
		fieldDef := `"` + col.Name + `" ` + sqlType
		if col.Name == "_cq_id" {
			// _cq_id column should always have a "unique not null" constraint
			fieldDef += " UNIQUE NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
		if c.enabledPks() && col.CreationOptions.PrimaryKey {
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
	} else {
		// if no primary keys are defined, add a PK constraint for _cq_id
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(table.Name)
		sb.WriteString("_cqpk PRIMARY KEY (_cq_id)")
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
	return &info, nil
}

func (c *Client) enabledPks() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}
