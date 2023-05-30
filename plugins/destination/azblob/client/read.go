package client

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- arrow.Record) error {
	if !c.pluginSpec.NoRotate {
		return fmt.Errorf("reading is not supported when `no_rotate` is false. Table: %q; Source: %q", table.Name, sourceName)
	}
	name := fmt.Sprintf("%s/%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format)

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
	return c.Client.Read(byteReader, table, sourceName, res)
}
