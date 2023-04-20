package client

import (
	"bytes"
	"context"
	"fmt"
	"io"

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
	var b bytes.Buffer
	w := io.Writer(&b)
	if err := c.Client.WriteTableBatchFile(w, arrowSchema, data); err != nil {
		return err
	}
	r := io.Reader(&b)
	_, err := c.storageClient.UploadStream(ctx, c.pluginSpec.Container, name, r, nil)
	if err != nil {
		return err
	}
	return nil
}
