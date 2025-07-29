package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	sqlTableInfo      = "PRAGMA table_info('%s');"
	isColumnUniqueSQL = "select count(*) from duckdb_constraints where table_name = $1 and constraint_type = 'UNIQUE' and constraint_column_names=[$2]"
)

type columnInfo struct {
	index        int
	name         string
	typ          string
	notNull      bool
	defaultValue any
	pk           bool
	unique       bool
}

type tableInfo struct {
	columns []columnInfo
}

func (*Client) normalizeColumns(tables schema.Tables) schema.Tables {
	var normalized schema.Tables
	for _, table := range tables {
		normalizedTable := *table
		normalizedTable.Columns = make(schema.ColumnList, len(table.Columns))
		for i := range table.Columns {
			// In DuckDB, a PK column must be NOT NULL, so we need to make sure that the schema we're comparing to has the same
			// constraint.
			normalizedColumn := table.Columns[i]
			if normalizedColumn.PrimaryKey {
				normalizedColumn.NotNull = true
			}
			// Since multiple schema types can map to the same duckdb type we need to normalize them to avoid false positives when detecting schema changes
			normalizedColumn.Type = duckDBToArrow(arrowToDuckDB(normalizedColumn.Type))
			normalizedTable.Columns[i] = normalizedColumn
		}
		normalized = append(normalized, &normalizedTable)
	}

	return normalized
}

func (c *Client) nonAutoMigratableTables(tables schema.Tables, duckdbTables schema.Tables) map[string][]schema.TableColumnChange {
	result := make(map[string][]schema.TableColumnChange)
	for _, t := range tables {
		duckdbTable := duckdbTables.Get(t.Name)
		if duckdbTable == nil {
			continue
		}
		changes := t.GetChanges(duckdbTable)
		if !c.canAutoMigrate(changes) {
			result[t.Name] = changes
		}
	}
	return result
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table, changes []schema.TableColumnChange) error {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd {
			if err := c.addColumn(ctx, table.Name, change.Current.Name, arrowToDuckDB(change.Current.Type)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (*Client) canAutoMigrate(changes []schema.TableColumnChange) bool {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd && (change.Current.PrimaryKey || change.Current.NotNull) {
			return false
		}

		if change.Type == schema.TableColumnChangeTypeRemove && (change.Previous.PrimaryKey || change.Previous.NotNull) {
			return false
		}

		if change.Type == schema.TableColumnChangeTypeUpdate {
			return false
		}
		if change.Type == schema.TableColumnChangeTypeRemoveUniqueConstraint {
			return false
		}
	}
	return true
}

// Migrate migrates to the latest schema
func (c *Client) MigrateTables(ctx context.Context, msgs message.WriteMigrateTables) error {
	tables := make(schema.Tables, len(msgs))
	for i, msg := range msgs {
		tables[i] = msg.Table
	}

	duckdbTables := make(schema.Tables, 0, len(tables))
	for _, table := range tables {
		t, err := c.getTableInfo(ctx, table.Name)
		if err != nil {
			return err
		}
		if t != nil {
			duckdbTables = append(duckdbTables, t)
		}
	}

	normalizedTables := c.normalizeColumns(tables)
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

	nonAutoMigratableTables := c.nonAutoMigratableTables(normalizedTablesSafeMode, duckdbTables)
	if len(nonAutoMigratableTables) > 0 {
		return fmt.Errorf("\nCan't migrate tables automatically, migrate manually or consider using 'migrate_mode: forced'. Non auto migratable tables changes:\n\n%s", schema.GetChangesSummary(nonAutoMigratableTables))
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
			if err := c.createTableIfNotExist(ctx, table.Name, table); err != nil {
				return err
			}
			continue
		}

		changes := table.GetChanges(duckdb)
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

	return nil
}

func (c *Client) recreateTable(ctx context.Context, table *schema.Table) error {
	sql := "drop table if exists " + sanitizeID(table.Name)
	if err := c.exec(ctx, sql); err != nil {
		return err
	}
	return c.createTableIfNotExist(ctx, table.Name, table)
}

func (c *Client) addColumn(ctx context.Context, tableName string, columnName string, columnType string) error {
	sql := "alter table " + sanitizeID(tableName) + " add column " + sanitizeID(columnName) + " " + columnType
	return c.exec(ctx, sql)
}

func (c *Client) createTableIfNotExist(ctx context.Context, tableName string, table *schema.Table) error {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(sanitizeID(tableName))
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	var pks []string
	for i, col := range table.Columns {
		sqlType := arrowToDuckDB(col.Type)
		fieldDef := sanitizeID(col.Name) + ` ` + sqlType
		if col.PrimaryKey {
			pks = append(pks, col.Name)
		}
		if col.Unique {
			fieldDef += " UNIQUE"
		}
		if col.NotNull {
			fieldDef += " NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
	}
	if len(pks) > 0 {
		sb.WriteString(", PRIMARY KEY (")
		for i, pk := range pks {
			sb.WriteString(sanitizeID(pk))
			if i != len(pks)-1 {
				sb.WriteString(",")
			}
		}
		sb.WriteString(")")
	}
	sb.WriteString(")")
	return c.exec(ctx, sb.String())
}

func (c *Client) isColumnUnique(ctx context.Context, tableName string, columName string) (bool, error) {
	rows, err := c.db.QueryContext(ctx, isColumnUniqueSQL, tableName, columName)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		var n int
		if err := rows.Scan(&n); err != nil {
			return false, err
		}
		if n == 1 {
			return true, nil
		}
	}
	if err := rows.Err(); err != nil {
		return false, err
	}

	return false, nil
}

func (c *Client) getTableInfo(ctx context.Context, tableName string) (*schema.Table, error) {
	info := tableInfo{}
	rows, err := c.db.QueryContext(ctx, fmt.Sprintf(sqlTableInfo, tableName))
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
		colInfo.unique, err = c.isColumnUnique(ctx, tableName, colInfo.name)
		if err != nil {
			return nil, err
		}
		info.columns = append(info.columns, colInfo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(info.columns) == 0 {
		// Table doesn't exist
		return nil, nil
	}

	columns := make(schema.ColumnList, len(info.columns))
	for i, col := range info.columns {
		columns[i] = schema.Column{
			Name:       col.name,
			Type:       duckDBToArrow(col.typ),
			NotNull:    col.notNull,
			PrimaryKey: col.pk,
			Unique:     col.unique,
		}
	}

	return &schema.Table{
		Name:    tableName,
		Columns: columns,
	}, nil
}
