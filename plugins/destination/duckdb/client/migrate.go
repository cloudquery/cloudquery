package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
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
	pk           bool
}

type tableInfo struct {
	columns []columnInfo
}

func (c *Client) duckdbTables(tables schema.Tables) (schema.Tables, error) {
	allTables := tables.FlattenTables()
	var schemaTables schema.Tables
	for _, table := range allTables {
		info, err := c.getTableInfo(table.Name)
		if err != nil {
			return nil, err
		}
		if info == nil {
			continue
		}
		schemaTable := schema.Table{
			Name: table.Name,
		}
		for _, col := range info.columns {
			schemaTable.Columns = append(schemaTable.Columns, schema.Column{
				Name: col.name,
				Type: c.duckdbTypeToSchema(col.typ),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: col.pk,
					NotNull:    col.notNull,
				},
			})
		}
		schemaTables = append(schemaTables, &schemaTable)
	}

	return schemaTables, nil
}

func (c *Client) normalizeColumns(tables schema.Tables) schema.Tables {
	allTables := tables.FlattenTables()
	var normalized schema.Tables
	for _, table := range allTables {
		tableCopy := table.Copy(table.Parent)
		for i, col := range tableCopy.Columns {
			// In DuckDB, a PK column must be NOT NULL, so we need to make sure that the schema we're comparing to has the same
			// constraint.
			if !c.enabledPks() {
				tableCopy.Columns[i].CreationOptions.PrimaryKey = false
			} else if col.CreationOptions.PrimaryKey {
				tableCopy.Columns[i].CreationOptions.NotNull = true
			}
			// Since multiple schema types can map to the same duckdb type we need to normalize them to avoid false positives when detecting schema changes
			tableCopy.Columns[i].Type = c.duckdbTypeToSchema(c.SchemaTypeToDuckDB(table.Columns[i].Type))
		}

		normalized = append(normalized, tableCopy)
	}

	return normalized
}

func (c *Client) nonAutoMigrableTables(tables schema.Tables, duckdbTables schema.Tables) ([]string, [][]schema.TableColumnChange) {
	var result []string
	var tableChanges [][]schema.TableColumnChange
	for _, t := range tables {
		duckdbTable := duckdbTables.Get(t.Name)
		if duckdbTable == nil {
			continue
		}
		changes := t.GetChanges(duckdbTable)
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
			if err := c.addColumn(table.Name, change.Current.Name, c.SchemaTypeToDuckDB(change.Current.Type)); err != nil {
				return err
			}
		}
	}
	return nil
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

// Migrate migrates to the latest schema
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	duckdbTables, err := c.duckdbTables(tables)
	if err != nil {
		return err
	}

	normalizedTables := c.normalizeColumns(tables)
	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonAutoMigrableTables, changes := c.nonAutoMigrableTables(normalizedTables, duckdbTables)
		if len(nonAutoMigrableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonAutoMigrableTables, ","), changes)
		}
	}

	for _, table := range normalizedTables {
		c.logger.Info().Str("table", table.Name).Msg("Migrating table")
		if len(table.Columns) == 0 {
			c.logger.Info().Str("table", table.Name).Msg("Table with no columns, skipping")
			continue
		}
		duckdb := duckdbTables.Get(table.Name)
		if duckdb == nil {
			c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(table); err != nil {
				return err
			}
			continue
		}
		changes := table.GetChanges(duckdb)
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
	// TODO sanitize table name
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(`"` + table.Name + `"`)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	for i, col := range table.Columns {
		sqlType := c.SchemaTypeToDuckDB(col.Type)
		if sqlType == "" {
			c.logger.Warn().Str("table", table.Name).Str("column", col.Name).Msg("Column type is not supported, skipping")
			continue
		}
		// TODO: sanitize column name
		fieldDef := `"` + col.Name + `" ` + sqlType
		if c.enabledPks() {
			if col.CreationOptions.PrimaryKey {
				fieldDef += " PRIMARY KEY"
			}
		}
		if col.CreationOptions.NotNull {
			fieldDef += " NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
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
		if strings.Contains(err.Error(), fmt.Sprintf("Table with name %s does not exist!", tableName)) {
			// Table doesn't exist
			return nil, nil
		}
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
