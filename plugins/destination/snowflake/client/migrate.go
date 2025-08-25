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
	sqlTableInfo = "SELECT table_name, column_name, data_type, is_nullable FROM information_schema.columns WHERE table_schema=CURRENT_SCHEMA() AND UPPER(table_name) IN "
)

type columnInfo struct {
	name     string
	typ      string
	nullable snowflakeYesNo
}

type tableInfo struct {
	name    string
	columns []columnInfo
}

type snowflakeYesNo bool

func (s *snowflakeYesNo) Scan(value any) error {
	if value == nil {
		*s = false
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot scan %T into snowflakeYesNo", value)
	}

	switch str {
	case "YES":
		*s = true
	case "NO":
		*s = false
	default:
		return fmt.Errorf("failed to scan yes/no string: %s", str)
	}
	return nil
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
	var wantTables []string
	for _, msg := range msgs {
		wantTables = append(wantTables, msg.Table.Name)
	}

	tableInfos, err := c.getTableInfo(ctx, wantTables)
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

			if currentTable, tableExists := tableInfos[strings.ToUpper(tableName)]; tableExists {
				c.logger.Debug().Str("table", tableName).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(gctx, table, currentTable, migrateForce); err != nil {
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

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table, currentTable tableInfo, force bool) error {
	tableName := table.Name

	for _, col := range table.Columns {
		columnName := strings.ToUpper(col.Name)
		columnType := c.SchemaTypeToSnowflake(col.Type)
		snowflakeColumn := currentTable.getColumn(columnName)

		if len(snowflakeColumn) == 1 && !strings.EqualFold(snowflakeColumn[0].typ, columnType) {
			if !force {
				return fmt.Errorf("column %s on table %s has different type than schema, expected %s got %s. migrate manually or consider using 'migrate_mode: forced'", col.Name, tableName, columnType, snowflakeColumn[0].typ)
			}
			sfCol := snowflakeColumn[0]
			c.logger.Debug().
				Str("table", tableName).
				Str("column", sfCol.name).
				Str("current_type", sfCol.typ).
				Str("want_type", columnType).
				Msg("Column type mismatch, dropping to recreate")
			sql := fmt.Sprintf("alter table %s drop column %q", tableName, sfCol.name)
			if _, err := c.db.ExecContext(ctx, sql); err != nil {
				return fmt.Errorf("failed to drop column %s from table %s: %w", sfCol.name, tableName, err)
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
		} else if col.NotNull {
			fieldDef += " NOT NULL"
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

func (c *Client) getTableInfo(ctx context.Context, tableNames []string) (map[string]tableInfo, error) {
	infos := make(map[string]tableInfo, len(tableNames))

	tnAny := make([]any, len(tableNames))
	for i := range tableNames {
		tnAny[i] = strings.ToUpper(tableNames[i])
	}
	completeSQL := sqlTableInfo + "(" + strings.Repeat("?,", len(tableNames)-1) + "?)"

	rows, err := c.db.QueryContext(ctx, completeSQL, tnAny...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tableName string
		colInfo := columnInfo{}

		if err := rows.Scan(
			&tableName,
			&colInfo.name,
			&colInfo.typ,
			&colInfo.nullable); err != nil {
			return nil, err
		}

		colInfo.typ = strings.ToLower(colInfo.typ)
		info := infos[strings.ToUpper(tableName)]
		info.name = tableName
		info.columns = append(info.columns, colInfo)
		infos[strings.ToUpper(tableName)] = info
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return infos, nil
}
