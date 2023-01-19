package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/client/queries"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) tableName(table *schema.Table) string {
	return queries.SanitizeID(c.schemaName, table.Name)
}

func (c *Client) tableExists(ctx context.Context, table string) (bool, error) {
	const tableExistsQuery = `SELECT COUNT(*)
FROM information_schema.tables
WHERE [table_name] = @tableName AND [table_schema] = @tableSchema`

	var tableExist int
	row := c.db.QueryRowContext(ctx, tableExistsQuery,
		sql.Named("tableName", table),
		sql.Named("tableSchema", c.schemaName),
	)
	if err := row.Scan(&tableExist); err != nil {
		return false, fmt.Errorf("failed to check if table %s exists: %w", table, err)
	}
	return tableExist == 1, nil
}
