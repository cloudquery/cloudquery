package client

import (
	"context"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/queries"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) getPartitionKeyAndSortingKey(ctx context.Context, table *schema.Table) ([]string, []string, error) {
	sql := queries.GetPartitionKeyAndSortingKeyQuery(c.database, table.Name)
	var partitionKey, sortingKey string

	err := retryQueryRowAndScan(ctx, c.logger, c.conn, sql, []any{}, []any{&partitionKey, &sortingKey})
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

func (c *Client) getTTL(ctx context.Context, table *schema.Table) (string, error) {
	sql := queries.GetTTLQuery(c.database, table.Name)
	var statement string

	err := retryQueryRowAndScan(ctx, c.logger, c.conn, sql, []any{}, []any{&statement})
	if err != nil {
		return "", err
	}

	ttl := ""
	for _, line := range strings.Split(statement, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "TTL") {
			// Extract the TTL statement, which is everything after "TTL"
			ttl = strings.TrimPrefix(line, "TTL")
			ttl = strings.TrimSpace(ttl)
			break
		}
	}

	return ttl, nil
}
