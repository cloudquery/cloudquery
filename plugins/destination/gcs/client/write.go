package client

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/filetypes/v4/types"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
)

// objectWriter is the subset of *storage.Writer used by WriteTable.
type objectWriter interface {
	io.Writer
	Close() error
}

// onceCloser closes the underlying writer exactly once and remembers the
// result, so a writer can be released on every error path without risking a
// double close.
type onceCloser struct {
	objectWriter
	once sync.Once
	err  error
}

func newOnceCloser(w objectWriter) *onceCloser {
	return &onceCloser{objectWriter: w}
}

func (w *onceCloser) Close() error {
	w.once.Do(func() {
		w.err = w.objectWriter.Close()
	})
	return w.err
}

func (c *Client) WriteTable(ctx context.Context, msgs <-chan *message.WriteInsert) error {
	return writeTable(msgs, c.Client, func(table *schema.Table) objectWriter {
		name := c.spec.ReplacePathVariables(table.Name, uuid.NewString(), time.Now().UTC(), c.syncID)
		return c.gcsClient.Bucket(c.spec.Bucket).Object(name).NewWriter(ctx)
	})
}

// writeTable streams msgs into a single object created by open. The writer is
// closed on every path, including failures in the header, content and footer
// writes, so the in-progress upload and its connection are not leaked.
func writeTable(msgs <-chan *message.WriteInsert, ft types.FileType, open func(*schema.Table) objectWriter) error {
	var (
		w *onceCloser
		h types.Handle
	)

	for msg := range msgs {
		if w == nil {
			table := msg.GetTable()
			w = newOnceCloser(open(table))

			var err error
			h, err = ft.WriteHeader(w, table)
			if err != nil {
				return closeOnError(w, err)
			}
		}

		if err := h.WriteContent([]arrow.RecordBatch{msg.Record}); err != nil {
			return closeOnError(w, err)
		}
	}

	if w == nil {
		// No messages arrived, so no object was created and there is nothing to close.
		return nil
	}

	if err := h.WriteFooter(); err != nil {
		return closeOnError(w, err)
	}
	return w.Close()
}

// closeOnError releases the writer and returns the write error, joined with
// the close error if closing also failed.
func closeOnError(w *onceCloser, err error) error {
	if closeErr := w.Close(); closeErr != nil {
		return errors.Join(err, closeErr)
	}
	return err
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	return c.writer.Write(ctx, msgs)
}
