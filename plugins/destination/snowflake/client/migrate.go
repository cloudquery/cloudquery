package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"golang.org/x/sync/errgroup"
)

const (
	sqlTableList = "select table_name from information_schema.tables where table_schema=current_schema();"
	sqlTableInfo = "select column_name, data_type, is_nullable from information_schema.columns where table_schema=current_schema() and table_name ilike ?;"
)

type columnInfo struct {
	name    string
	typ     string
	notNull bool
}

type tableInfo struct {
	columns []columnInfo
}

func (i *tableInfo) getColumn(name string) []columnInfo {
	var cols []columnInfo
	for idx, col := range i.columns {
		if strings.ToUpper(col.name) == name {
			cols = append(cols, i.columns[idx])
		}
	}
	return cols
}

func (c *Client) MigrateTables(ctx context.Context, msgs message.WriteMigrateTables) error {
	tableList, err := c.listTables(ctx)
	if err != nil {
		return fmt.Errorf("failed to get list of tables: %w", err)
	}

	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(c.spec.MigrateConcurrency)
	for _, msg := range msgs {
		table := msg.Table
		migrateForce := msg.MigrateForce
		tableName := table.Name
		g.Go(func() error {
			c.logger.Debug().Str("table", tableName).Msg("Migrating table")
			if tableExists(tableList, tableName) {
				c.logger.Debug().Str("table", tableName).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(gctx, table, migrateForce); err != nil {
					return err
				}
			} else {
				c.logger.Debug().Str("table", tableName).Msg("Table doesn't exist, creating")
				if err := c.createTableIfNotExist(gctx, table); err != nil {
					return err
				}
			}
			return nil
		})
	}
	return g.Wait()
}

func (c *Client) listTables(ctx context.Context) ([]string, error) {
	var tables []string
	rows, err := c.db.QueryContext(ctx, sqlTableList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t string
		if err := rows.Scan(&t); err != nil {
			return nil, err
		}
		tables = append(tables, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tables, nil
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table, force bool) error {
	var err error
	var info *tableInfo
	tableName := table.Name
	if info, err = c.getTableInfo(ctx, tableName); err != nil {
		return fmt.Errorf("failed to get table %s columns types: %w", tableName, err)
	}

	for _, col := range table.Columns {
		columnName := strings.ToUpper(col.Name)
		columnType := c.SchemaTypeToSnowflake(col.Type)
		snowflakeColumn := info.getColumn(columnName)

		if len(snowflakeColumn) == 1 && !strings.EqualFold(snowflakeColumn[0].typ, columnType) {
			if !force {
				return fmt.Errorf("column %s on table %s has different type than schema, expected %s got %s. migrate manually or consider using 'migrate_mode: forced'", col.Name, tableName, columnType, snowflakeColumn[0].typ)
			}
			c.logger.Debug().Str("table", tableName).Str("column", col.Name).Str("current_type", snowflakeColumn[0].typ).Str("want_type", columnType).Msg("Column type mismatch, dropping to recreate")
			sql := fmt.Sprintf("alter table %s drop column %q", tableName, columnName)
			if _, err := c.db.ExecContext(ctx, sql); err != nil {
				return fmt.Errorf("failed to drop column %s from table %s: %w", col.Name, tableName, err)
			}
			snowflakeColumn = nil
			// proceed to add column
		}

		switch {
		case len(snowflakeColumn) == 0:
			c.logger.Debug().Str("table", tableName).Str("column", col.Name).Msg("Column doesn't exist, creating")
			sql := fmt.Sprintf("alter table %s add column %q %s", tableName, columnName, columnType)
			if _, err := c.db.ExecContext(ctx, sql); err != nil {
				return fmt.Errorf("failed to add column %s on table %s: %w", col.Name, tableName, err)
			}

		// have multiple columns, drop all but one
		case len(snowflakeColumn) > 1 && !force:
			return fmt.Errorf("table %s has multiple columns for %s. migrate manually or consider using 'migrate_mode: forced'", tableName, col.Name)

		case len(snowflakeColumn) > 1:
			for _, sc := range snowflakeColumn {
				if sc.name != columnName {
					c.logger.Debug().Str("table", tableName).Str("column", columnName).Msg("Column exists with different name, dropping")
					sql := fmt.Sprintf("alter table %s drop column %q", tableName, sc.name)
					if _, err := c.db.ExecContext(ctx, sql); err != nil {
						return fmt.Errorf("failed to drop column %s on table %s: %w", sc.name, tableName, err)
					}
				}
			}

		case snowflakeColumn[0].name != columnName: // case sensitivity
			c.logger.Debug().Str("table", tableName).Str("column", columnName).Str("current_name", snowflakeColumn[0].name).Msg("Column name doesn't match, migrating")
			sql := fmt.Sprintf("alter table %s rename column %q TO %q", tableName, snowflakeColumn[0].name, columnName)
			if _, err := c.db.ExecContext(ctx, sql); err != nil {
				return fmt.Errorf("failed to rename column %s on table %s: %w", snowflakeColumn[0].name, tableName, err)
			}
		}
	}
	return nil
}

func (c *Client) createTableIfNotExist(ctx context.Context, table *schema.Table) error {
	var sb strings.Builder
	// TODO sanitize tablename
	tableName := table.Name
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(tableName)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	for i, col := range table.Columns {
		sqlType := c.SchemaTypeToSnowflake(col.Type)
		// TODO: sanitize column name
		fieldDef := `"` + strings.ToUpper(col.Name) + `" ` + sqlType
		if col.Name == schema.CqIDColumn.Name {
			// _cq_id column should always have a "unique not null" constraint
			fieldDef += " UNIQUE NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
	}

	sb.WriteString(")")
	_, err := c.db.ExecContext(ctx, sb.String())
	if err != nil {
		return fmt.Errorf("failed to create table with '%s': %w", sb.String(), err)
	}
	return nil
}

func (c *Client) getTableInfo(ctx context.Context, tableName string) (*tableInfo, error) {
	info := tableInfo{}
	rows, err := c.db.QueryContext(ctx, sqlTableInfo, tableName)
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

func tableExists(list []string, table string) bool {
	tbl := strings.ToUpper(table)
	for _, t := range list {
		if strings.ToUpper(t) == tbl {
			return true
		}
	}
	return false
}
