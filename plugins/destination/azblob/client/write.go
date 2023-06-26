package client

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
)

func (c *Client) OpenTable(ctx context.Context, sourceName string, table *schema.Table, syncTime time.Time) (any, error) {
	name := fmt.Sprintf("%s/%s.%s.%s", c.spec.Path, table.Name, c.spec.Format, uuid.NewString())
	if c.spec.NoRotate {
		name = fmt.Sprintf("%s/%s.%s", c.spec.Path, table.Name, c.spec.Format)
	}

	return c.Client.StartStream(table, func(r io.Reader) error {
		_, err := c.storageClient.UploadStream(ctx, c.spec.Container, name, r, nil)
		return err
	})
}

func (*Client) CloseTable(_ context.Context, handle any) error {
	return handle.(*filetypes.Stream).Finish()
}

func (*Client) WriteTableStream(_ context.Context, handle any, msgs []*message.Insert) error {
	if len(msgs) == 0 {
		return nil
	}

	records := make([]arrow.Record, len(msgs))
	for i, msg := range msgs {
		records[i] = msg.Record
	}

	return handle.(*filetypes.Stream).Write(records)
}

func (c *Client) Write(ctx context.Context, options plugin.WriteOptions, msgs <-chan message.Message) error {
	return c.writer.Write(ctx, msgs)
}
