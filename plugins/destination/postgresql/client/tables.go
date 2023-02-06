package client

import "context"

func (c *Client) getPgTables(ctx context.Context) error {
	tag, err := c.conn.Exec(ctx, sqlSelectAllTables)
}