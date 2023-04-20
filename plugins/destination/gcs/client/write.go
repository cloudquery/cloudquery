package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/google/uuid"
)

func (c *Client) WriteTableBatch(ctx context.Context, arrowSchema *arrow.Schema, data []arrow.Record) error {
	if len(data) == 0 {
		return nil
	}
	tableName := schema.TableName(arrowSchema)
	name := fmt.Sprintf("%s/%s.%s.%s", c.pluginSpec.Path, tableName, c.pluginSpec.Format, uuid.NewString())
	if c.pluginSpec.NoRotate {
		name = fmt.Sprintf("%s/%s.%s", c.pluginSpec.Path, tableName, c.pluginSpec.Format)
	}
	w := c.gcsClient.Bucket(c.pluginSpec.Bucket).Object(name).NewWriter(ctx)
	defer w.Close()

	return c.Client.WriteTableBatchFile(w, arrowSchema, data)
}
