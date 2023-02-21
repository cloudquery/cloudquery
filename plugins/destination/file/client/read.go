package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	if !c.pluginSpec.NoRotate {
		return fmt.Errorf("reading is not supported when no_rotate is false. Table: %q; Source: %q", table.Name, sourceName)
	}
	name := fmt.Sprintf("%s/%s.%s", c.pluginSpec.Directory, table.Name, c.pluginSpec.Format)
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	return c.formatClient.Read(f, table, sourceName, res)
}
