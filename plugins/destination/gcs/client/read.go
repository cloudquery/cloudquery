package client

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) Read(ctx context.Context, arrowSchema *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	tableName := schema.TableName(arrowSchema)
	if !c.pluginSpec.NoRotate {
		return fmt.Errorf("reading is not supported when `no_rotate` is false. Table: %q; Source: %q", tableName, sourceName)
	}
	name := fmt.Sprintf("%s/%s.%s", c.pluginSpec.Path, tableName, c.pluginSpec.Format)
	r, err := c.bucket.Object(name).NewReader(ctx)
	if err != nil {
		return err
	}
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	byteReader := bytes.NewReader(b)
	return c.Client.Read(byteReader, arrowSchema, sourceName, res)
}
