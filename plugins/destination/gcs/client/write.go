package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/uuid"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
	name := fmt.Sprintf("%s/%s.%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format, uuid.NewString())
	if c.pluginSpec.NoRotate {
		name = fmt.Sprintf("%s/%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format)
	}
	w := c.gcsClient.Bucket(c.pluginSpec.Bucket).Object(name).NewWriter(ctx)
	defer w.Close()

	return c.Client.WriteTableBatchFile(w, table, data)
}
