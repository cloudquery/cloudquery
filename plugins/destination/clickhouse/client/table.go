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

func (c *Client) equalTTLs(table *schema.Table, ttl1, ttl2 string) (bool, error) {
	// ClickHouse allows different syntaxes for the same TTL expression,
	// and a user query like "TTL col + INTERVAL 1 DAY" gets converted to SHOW CREATE TABLE as "TTL col + toIntervalDay(1)".
	// Therefore this function runs a query to compare two given TTLs and returns true if they are equivalent according
	// to ClickHouse.
	if ttl1 == "" && ttl2 != "" || ttl1 != "" && ttl2 == "" {
		// If one TTL is empty and the other is not, they are not equal.
		return false, nil
	}
	if ttl1 == ttl2 {
		// If both TTLs are exactly the same, they are equal.
		return true, nil
	}

	// If they are different, we need to compare them using a query.
	sql := queries.EqualTTLsQuery(table, ttl1, ttl2)
	var result *uint8
	err := retryQueryRowAndScan(context.Background(), c.logger, c.conn, sql, []any{}, []any{&result})
	if err != nil {
		return false, err
	}
	return *result == 1, nil
}
