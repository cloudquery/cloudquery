package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/google/uuid"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data []arrow.Record) error {
	if len(data) == 0 {
		return nil
	}
	name := fmt.Sprintf("%s/%s.%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format, uuid.NewString())
	if c.pluginSpec.NoRotate {
		name = fmt.Sprintf("%s/%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format)
	}
	w := c.gcsClient.Bucket(c.pluginSpec.Bucket).Object(name).NewWriter(ctx)
	defer w.Close()

	return c.Client.WriteTableBatchFile(w, table, data)
}
