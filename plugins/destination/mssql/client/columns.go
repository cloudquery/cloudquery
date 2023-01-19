package client

import (
	"context"
	"database/sql"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/client/queries"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"golang.org/x/exp/slices"
)

func (c *Client) pkEnabled() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}

func (c *Client) getTableColumns(ctx context.Context, table *schema.Table) (queries.Definitions, error) {
	query, params := queries.GetTableSchema(c.schemaName, table)
	var tc queries.Definitions

	rows, err := c.db.QueryContext(ctx, query, params...)
	if err != nil {
		c.logErr(err)
		return nil, err
	}

	if err := processRows(rows, func(row *sql.Rows) error {
		var name string
		var typ string
		var nullable bool

		if err := row.Scan(&name, &typ, &nullable); err != nil {
			return err
		}

		tc = append(tc, queries.NewDefinition(name, typ, nullable))

		return nil
	}); err != nil {
		c.logErr(err)
		return nil, err
	}

	return tc, nil
}

func (c *Client) getTablePK(ctx context.Context, table *schema.Table) ([]string, error) {
	query, params := queries.GetTablePK(c.schemaName, table)

	rows, err := c.db.QueryContext(ctx, query, params...)
	if err != nil {
		c.logErr(err)
		return nil, err
	}

	var result []string
	if err := processRows(rows, func(row *sql.Rows) error {
		var name string
		var idx int

		if err := rows.Scan(&name, &idx); err != nil {
			return err
		}

		result = append(result, name)

		return nil
	}); err != nil {
		c.logErr(err)
		return nil, err
	}

	return result, nil
}

func (c *Client) getStalePks(table *schema.Table, primaryKey []string) []string {
	if !c.pkEnabled() {
		return nil
	}

	schemaPK := table.PrimaryKeys()

	var stale []string
	for _, key := range primaryKey {
		if !slices.Contains(schemaPK, key) {
			stale = append(stale, key)
		}
	}

	return stale
}

func (c *Client) getDropNotNullQuery(table *schema.Table, stalePK []string) string {
	tableName := c.tableName(table)
	statements := make([]string, len(stalePK))
	for i, col := range stalePK {
		statements[i] = "ALTER TABLE " + tableName + " ALTER COLUMN " + queries.SanitizeID(col) + " DROP NOT NULL;"
	}

	return strings.Join(statements, "\n")
}
