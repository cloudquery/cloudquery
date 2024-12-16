package client

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/filetypes/v4/types"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/google/uuid"
)

var syncAfterWrite bool // used in testing

func (c *Client) WriteTable(_ context.Context, msgs <-chan *message.WriteInsert) error {
	var (
		f *os.File
		h types.Handle
	)
	for msg := range msgs {
		if f == nil {
			table := msg.GetTable()
			p := c.spec.ReplacePathVariables(table.Name, uuid.NewString(), time.Now().UTC())
			if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}

			var err error
			f, err = os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}

			h, err = c.Client.WriteHeader(f, table)
			if err != nil {
				return err
			}
		}

		if err := h.WriteContent([]arrow.Record{msg.Record}); err != nil {
			return err
		}
		if syncAfterWrite {
			_ = f.Sync()
		}
	}

	if err := h.WriteFooter(); err != nil {
		return err
	}
	return f.Close()
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	return c.writer.Write(ctx, msgs)
}
