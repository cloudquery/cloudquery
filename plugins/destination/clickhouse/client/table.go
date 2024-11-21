package client

import (
	"context"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) getTableDefinitions(ctx context.Context, messages message.WriteMigrateTables) (schema.Tables, error) {
	// need proper description without flattened columns
	const flattenNested0 = "SET flatten_nested = 0"
	if err := c.conn.Exec(ctx, flattenNested0); err != nil {
		return nil, err
	}

	query, params := queries.GetTablesSchema(c.database)
	rows, err := c.conn.Query(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return queries.ScanTableSchemas(rows, messages)
}

func (c *Client) getPartitionKeyAndSortingKey(ctx context.Context, table *schema.Table) ([]string, []string, error) {
	sql := queries.GetPartitionKeyAndSortingKeyQuery(c.database, table.Name)
	var partitionKey, sortingKey string
	err := c.conn.QueryRow(ctx, sql).Scan(&partitionKey, &sortingKey)
	if err != nil {
		return nil, nil, err
	}

	splitPartitionKey := []string{}
	if partitionKey != "" {
		splitPartitionKey = strings.Split(partitionKey, ", ")
	}
	splitSortingKey := []string{}
	if sortingKey != "" {
		splitSortingKey = strings.Split(sortingKey, ", ")
	}
	return splitPartitionKey, splitSortingKey, nil
}
