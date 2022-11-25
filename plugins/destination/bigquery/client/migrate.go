package client

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	doesTableExistSQL = "SELECT count(*) as count FROM information_schema.tables WHERE table_name=@table_name;"
	sqlTableInfo      = "select column_name, data_type, is_nullable from information_schema.columns where table_name=@table_name;"
)

type countRow struct {
	Count int64 `bigquery:"count"`
}

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

// Migrate tables. It is the responsibility of the CLI of the client to lock before running migrations.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	for _, table := range tables {
		c.logger.Debug().Str("table", table.Name).Msg("Migrating table")
		tableExists, err := c.doesTableExist(ctx, table.Name)
		if err != nil {
			return fmt.Errorf("failed to check if table %s exists: %w", table.Name, err)
		}
		if tableExists {
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

func (c *Client) doesTableExist(ctx context.Context, table string) (bool, error) {
	var cr countRow
	q := c.client.Query(doesTableExistSQL)
	q.Parameters = []bigquery.QueryParameter{
		{Name: "table_name", Value: table},
	}
	ri, err := q.Read(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to check if table %s exists: %w", table, err)
	}
	if ri.TotalRows != 1 {
		return false, fmt.Errorf("failed to check if table %s exists (%d rows in result)", table, ri.TotalRows)
	}
	ri.Next(&cr)
	return cr.Count == 1, nil
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table) error {
	var err error
	var info *tableInfo

	if info, err = c.getTableInfo(ctx, table.Name); err != nil {
		return fmt.Errorf("failed to get table %s columns types: %w", table.Name, err)
	}

	for _, col := range table.Columns {
		columnName := col.Name
		columnType := c.SchemaTypeToBigQuery(col.Type)
		sqliteColumn := info.getColumn(columnName)

		switch {
		case sqliteColumn == nil:
			c.logger.Debug().Str("table", table.Name).Str("column", col.Name).Msg("Column doesn't exist, creating")
			sql := "alter table " + table.Name + " add column \"" + columnName + "\" \"" + columnType + `"`
			job, err := c.client.Query(sql).Run(ctx)
			if err != nil {
				return fmt.Errorf("failed to add column %s on table %s: %w", col.Name, table.Name, err)
			}
			_, err = job.Wait(ctx)
			if err != nil {
				return err
			}
		case sqliteColumn.typ != columnType:
			return fmt.Errorf("column %s on table %s has different type than schema, expected %s got %s. trying dropping table and re-running", col.Name, table.Name, columnType, sqliteColumn.typ)
		}
	}
	return nil
}

func (c *Client) createTableIfNotExist(ctx context.Context, table *schema.Table) error {
	var sb strings.Builder
	// TODO sanitize tablename
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(table.Name)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	for i, col := range table.Columns {
		sqlType := c.SchemaTypeToBigQuery(col.Type)
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
	job, err := c.client.Query(sb.String()).Run(ctx)
	if err != nil {
		return fmt.Errorf("failed to create table with '%s': %w", sb.String(), err)
	}
	_, err = job.Wait(ctx)
	return err
}

func (c *Client) getTableInfo(ctx context.Context, tableName string) (*tableInfo, error) {
	meta, err := c.client.Dataset(c.datasetID).Table(tableName).Metadata(ctx)
	if err != nil {
		return nil, err
	}
	// TODO
	return &tableInfo{
		meta.Name
	}
}
