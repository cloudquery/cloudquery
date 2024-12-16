package client

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	if !c.spec.NoRotate {
		return fmt.Errorf("reading is not supported when `no_rotate` is false. Table: %q", table.Name)
	}
	name := fmt.Sprintf("%s/%s.%s%s", c.spec.Path, table.Name, c.spec.Format, c.spec.FileSpec.Compression.Extension())

	response, err := c.storageClient.DownloadStream(ctx, c.spec.Container, name, nil)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	byteReader := bytes.NewReader(b)
	return c.Client.Read(byteReader, table, res)
}
