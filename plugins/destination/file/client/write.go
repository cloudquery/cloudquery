package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/google/uuid"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
	name := fmt.Sprintf("%s/%s.%s.%s", c.pluginSpec.Directory, table.Name, c.pluginSpec.Format, uuid.NewString())
	if c.pluginSpec.NoRotate {
		name = fmt.Sprintf("%s/%s.%s", c.pluginSpec.Directory, table.Name, c.pluginSpec.Format)
	}
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return c.Client.WriteTableBatchFile(f, table, data)
}
