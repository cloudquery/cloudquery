package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

const (
	// Use ILIKE for case insensitivity

	isTableExistSQL = "SELECT count(*) FROM information_schema.tables WHERE table_name ILIKE ?;"
	sqlTableInfo    = "select column_name, data_type, is_nullable from information_schema.columns where table_name ILIKE ?;"
)

type columnInfo struct {
	name    string
	typ     string
	notNull bool
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

// This is the responsibility of the CLI of the client to lock before running migration
func (c *Client) Migrate(ctx context.Context, tables schema.Schemas) error {
	for _, table := range tables {
		tableName := schema.TableName(table)
		c.logger.Debug().Str("table", tableName).Msg("Migrating table")
		tableExist, err := c.isTableExistSQL(ctx, tableName)
		if err != nil {
			return fmt.Errorf("failed to check if table %s exists: %w", tableName, err)
		}
		if tableExist {
			c.logger.Debug().Str("table", tableName).Msg("Table exists, auto-migrating")
			if err := c.autoMigrateTable(ctx, table); err != nil {
				return err
			}
		} else {
			c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(ctx, table); err != nil {
				return err
			}
		}
		if err := c.Migrate(ctx, table.Relations); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) isTableExistSQL(_ context.Context, table string) (bool, error) {
	var tableExist int
	if err := c.db.QueryRow(isTableExistSQL, table).Scan(&tableExist); err != nil {
		return false, fmt.Errorf("failed to check if table %s exists: %w", table, err)
	}
	return tableExist == 1, nil
}

func (c *Client) autoMigrateTable(_ context.Context, table *arrow.Schema) error {
	var err error
	var info *tableInfo
	tableName := schema.TableName(table)
	if info, err = c.getTableInfo(tableName); err != nil {
		return fmt.Errorf("failed to get table %s columns types: %w", tableName, err)
	}

	for _, col := range table.Fields() {
		columnName := col.Name
		columnType := c.SchemaTypeToSnowflake(col.Type)
		snowflakeColumn := info.getColumn(columnName)

		switch {
		case snowflakeColumn == nil:
			c.logger.Debug().Str("table", tableName).Str("column", col.Name).Msg("Column doesn't exist, creating")
			sql := "alter table " + tableName + " add column \"" + columnName + "\"" + columnType
			fmt.Println(sql)
			if _, err := c.db.Exec(sql); err != nil {
				return fmt.Errorf("failed to add column %s on table %s: %w", col.Name, tableName, err)
			}
		case !strings.EqualFold(snowflakeColumn.typ, columnType):
			return fmt.Errorf("column %s on table %s has different type than schema, expected %s got %s. Try dropping the column and re-running", col.Name, table.Name, columnType, snowflakeColumn.typ)
		}
	}
	return nil
}

func (c *Client) createTableIfNotExist(_ context.Context, table *schema.Table) error {
	var sb strings.Builder
	// TODO sanitize tablename
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(table.Name)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	for i, col := range table.Columns {
		sqlType := c.SchemaTypeToSnowflake(col.Type)
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
	rows, err := c.db.Query(sqlTableInfo, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		colInfo := columnInfo{}

		// 'information_schema.is_nullable' is a string containing 'YES' or 'NO'.
		// We save it here as a string, and parse it later.
		var isNullableTemp string

		if err := rows.Scan(
			&colInfo.name,
			&colInfo.typ,
			&isNullableTemp); err != nil {
			return nil, err
		}

		colInfo.notNull, err = parseYesNoString(isNullableTemp)
		if err != nil {
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

func parseYesNoString(str string) (bool, error) {
	switch str {
	case "YES":
		return true, nil
	case "NO":
		return false, nil
	default:
		return false, fmt.Errorf("failed to parse yes/no string: %s", str)
	}
}
