package client

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	ftypes "github.com/cloudquery/filetypes/v4/types"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
)

type stream struct {
	h    ftypes.Handle
	wc   *writeCloser
	done chan error
}

type writeCloser struct {
	*io.PipeWriter
	closed bool
}

func (w *writeCloser) Close() error {
	w.closed = true
	return w.PipeWriter.Close()
}

func (c *Client) OpenTable(ctx context.Context, sourceName string, table *schema.Table, syncTime time.Time) (any, error) {
	name := fmt.Sprintf("%s/%s.%s.%s", c.spec.Path, table.Name, c.spec.Format, uuid.NewString())
	if c.spec.NoRotate {
		name = fmt.Sprintf("%s/%s.%s", c.spec.Path, table.Name, c.spec.Format)
	}

	pr, pw := io.Pipe()
	doneCh := make(chan error)

	go func() {
		_, err := c.storageClient.UploadStream(ctx, c.spec.Container, name, pr, nil)
		_ = pr.CloseWithError(err)
		doneCh <- err
		close(doneCh)
	}()

	wc := &writeCloser{PipeWriter: pw}
	h, err := c.Client.WriteHeader(wc, table)
	if err != nil {
		_ = pw.CloseWithError(err)
		<-doneCh
		return nil, err
	}

	return &stream{
		h:    h,
		wc:   wc,
		done: doneCh,
	}, nil
}

func (*Client) CloseTable(_ context.Context, handle any) error {
	s := handle.(*stream)
	if err := s.h.WriteFooter(); err != nil {
		if !s.wc.closed {
			_ = s.wc.CloseWithError(err)
		}
		return fmt.Errorf("failed to write footer: %w", <-s.done)
	}

	// ParquetWriter likes to close the underlying writer, so we need to check if it's already closed
	if !s.wc.closed {
		if err := s.wc.Close(); err != nil {
			return err
		}
	}

	return <-s.done
}

func (*Client) WriteTableStream(_ context.Context, handle any, msgs []*message.Insert) error {
	if len(msgs) == 0 {
		return nil
	}

	records := make([]arrow.Record, len(msgs))
	for i, msg := range msgs {
		records[i] = msg.Record
	}

	return handle.(*stream).h.WriteContent(records)
}

func (c *Client) Write(ctx context.Context, options plugin.WriteOptions, msgs <-chan message.Message) error {
	return c.writer.Write(ctx, msgs)
}
