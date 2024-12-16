package client

import (
	"context"
	"fmt"
	"io"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/google/uuid"
)

func (c *Client) WriteTable(ctx context.Context, msgs <-chan *message.WriteInsert) error {
	var s *filetypes.Stream

	for msg := range msgs {
		if s == nil {
			table := msg.GetTable()

			name := fmt.Sprintf("%s/%s.%s%s", c.spec.Path, table.Name, c.spec.Format, c.spec.FileSpec.Compression.Extension())
			if !c.spec.NoRotate {
				name += "." + uuid.NewString()
			}

			var err error
			s, err = c.Client.StartStream(table, func(r io.Reader) error {
				_, err := c.storageClient.UploadStream(ctx, c.spec.Container, name, r, nil)
				return err
			})
			if err != nil {
				return err
			}
		}

		if err := s.Write([]arrow.Record{msg.Record}); err != nil {
			_ = s.FinishWithError(err)
			return err
		}
	}

	return s.Finish()
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	return c.writer.Write(ctx, msgs)
}
