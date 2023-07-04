package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/filetypes/v4/types"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/google/uuid"
)

func (c *Client) WriteTable(ctx context.Context, msgs <-chan *message.WriteInsert) error {
	var (
		w *storage.Writer
		h types.Handle
	)

	for msg := range msgs {
		if w == nil {
			table := msg.GetTable()

			name := fmt.Sprintf("%s/%s.%s.%s", c.spec.Path, table.Name, c.spec.Format, uuid.NewString())
			if c.spec.NoRotate {
				name = fmt.Sprintf("%s/%s.%s", c.spec.Path, table.Name, c.spec.Format)
			}

			w = c.gcsClient.Bucket(c.spec.Bucket).Object(name).NewWriter(ctx)

			var err error
			h, err = c.Client.WriteHeader(w, table)
			if err != nil {
				return err
			}
		}

		if err := h.WriteContent([]arrow.Record{msg.Record}); err != nil {
			return err
		}
	}

	if err := h.WriteFooter(); err != nil {
		return err
	}
	return w.Close()
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	return c.writer.Write(ctx, msgs)
}
