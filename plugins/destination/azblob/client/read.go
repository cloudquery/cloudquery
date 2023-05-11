package client

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) Read(ctx context.Context, arrowSchema *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	tableName := schema.TableName(arrowSchema)
	if !c.pluginSpec.NoRotate {
		return fmt.Errorf("reading is not supported when `no_rotate` is false. Table: %q; Source: %q", tableName, sourceName)
	}
	name := fmt.Sprintf("%s/%s.%s", c.pluginSpec.Path, tableName, c.pluginSpec.Format)

	response, err := c.storageClient.DownloadStream(ctx, c.pluginSpec.Container, name, nil)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	byteReader := bytes.NewReader(b)
	return c.Client.Read(byteReader, arrowSchema, sourceName, res)
}
