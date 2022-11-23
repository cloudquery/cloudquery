package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) Write(ctx context.Context, tables schema.Tables, res <-chan *plugins.ClientResource) error {
	return fmt.Errorf("Write not implemented for snowflake. Please use file destination plugin for writing/reading to snowflake")
}
