package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

const (
	isTableExistSQL = "SELECT count(name) FROM sqlite_master WHERE type='table' AND name=?;"

	// https://wiki.postgresql.org/wiki/Retrieve_primary_key_columns
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

func (i *tableInfo) getColumn(name string) *columnInfo {
	for _, col := range i.columns {
		if col.name == name {
			return &col
		}
	}
	return nil
}

// This is the responsibility of the CLI of the client to lock before running migration
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	for _, table := range tables {
		c.logger.Debug().Str("table", table.Name).Msg("Migrating table")
		tableExist, err := c.isTableExistSQL(ctx, table.Name)
		if err != nil {
			return fmt.Errorf("failed to check if table %s exists: %w", table.Name, err)
		}
		if tableExist {
			c.logger.Debug().Str("table", table.Name).Msg("Table exists, auto-migrating")
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

func (c *Client) autoMigrateTable(_ context.Context, table *schema.Table) error {
	var err error
	var info *tableInfo

	if info, err = c.getTableInfo(table.Name); err != nil {
		return fmt.Errorf("failed to get table %s columns types: %w", table.Name, err)
	}

	for _, col := range table.Columns {
		columnName := col.Name
		columnType := c.SchemaTypeToSqlite(col.Type)
		sqliteColumn := info.getColumn(columnName)

		switch {
		case sqliteColumn == nil:
			c.logger.Debug().Str("table", table.Name).Str("column", col.Name).Msg("Column doesn't exist, creating")
			sql := "alter table \"" + table.Name + "\" add column \"" + columnName + "\" \"" + columnType + `"`
			if col.CreationOptions.PrimaryKey {
				return fmt.Errorf("failed to auto-migrate table %s: primary key changes are not supported. please drop and re-run", table.Name)
			}
			if _, err := c.db.Exec(sql); err != nil {
				return fmt.Errorf("failed to add column %s on table %s: %w", col.Name, table.Name, err)
			}
		case sqliteColumn.typ != columnType:
			return fmt.Errorf("column %s on table %s has different type than schema, expected %s got %s. Try dropping the column and re-running", col.Name, table.Name, columnType, sqliteColumn.typ)
		}
	}
	return nil
}

func (c *Client) createTableIfNotExist(_ context.Context, table *schema.Table) error {
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
