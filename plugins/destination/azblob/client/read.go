package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	if !c.pluginSpec.NoRotate {
		return fmt.Errorf("reading is not supported when no_rotate is false. Table: %q; Source: %q", table.Name, sourceName)
	}
	name := fmt.Sprintf("%s/%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format)

	response, err := c.storageClient.DownloadStream(ctx, c.pluginSpec.Container, name, nil)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return c.Client.Read(response.Body, table, sourceName, res)
}
