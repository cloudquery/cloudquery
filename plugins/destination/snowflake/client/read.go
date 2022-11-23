package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
)


func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	return fmt.Errorf("Read not implemented for snowflake. Please use file plugin for writing/reading to snowflake")
}
