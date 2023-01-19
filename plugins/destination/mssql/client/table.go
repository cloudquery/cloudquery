package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/client/queries"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) tableExists(ctx context.Context, table *schema.Table) (bool, error) {
	var tableExist int

	query, params := queries.TableExists(c.schemaName, table)
	if err := c.db.QueryRowContext(ctx, query, params...).Scan(&tableExist); err != nil {
		return false, fmt.Errorf("failed to check if table %s exists: %w",
			queries.SanitizedTableName(c.schemaName, table), err)
	}
	return tableExist == 1, nil
}
