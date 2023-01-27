package client

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/uuid"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
	name := fmt.Sprintf("%s/%s.%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format, uuid.NewString())
	if c.pluginSpec.NoRotate {
		name = fmt.Sprintf("%s/%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format)
	}
	var b bytes.Buffer
	w := io.Writer(&b)
	if err := c.Client.WriteTableBatchFile(w, table, data); err != nil {
		return err
	}
	r := io.Reader(&b)
	_, err := c.storageClient.UploadStream(ctx, c.pluginSpec.Container, name, r, nil)
	if err != nil {
		return err
	}
	return nil
}
