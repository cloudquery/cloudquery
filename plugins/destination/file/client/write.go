package client

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/filetypes/v4/types"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
)

type filestream struct {
	wc *writeCloser
	h  types.Handle
}

type writeCloser struct {
	io.WriteCloser
	closed bool
}

func (w *writeCloser) Close() error {
	w.closed = true
	return w.WriteCloser.Close()
}

func (c *Client) OpenTable(_ context.Context, sourceName string, table *schema.Table, syncTime time.Time) (any, error) {
	p := replacePathVariables(c.spec.Path, table.Name, c.spec.Format, uuid.NewString(), syncTime)

	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	wc := &writeCloser{WriteCloser: f}
	h, err := c.Client.WriteHeader(wc, table)
	if err != nil {
		_ = f.Close()
		return nil, fmt.Errorf("failed to write header: %w", err)
	}

	return &filestream{
		wc: wc,
		h:  h,
	}, nil
}

func (*Client) CloseTable(_ context.Context, handle any) error {
	fs := handle.(*filestream)
	if err := fs.h.WriteFooter(); err != nil {
		if !fs.wc.closed {
			_ = fs.wc.Close()
		}
		return fmt.Errorf("failed to write footer: %w", err)
	}

	// ParquetWriter likes to close the underlying writer, so we need to check if it's already closed
	if !fs.wc.closed {
		if err := fs.wc.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (*Client) WriteTableStream(_ context.Context, handle any, upsert bool, msgs []*message.Insert) error {
	if len(msgs) == 0 {
		return nil
	}

	records := make([]arrow.Record, len(msgs))
	for i, msg := range msgs {
		records[i] = msg.Record
	}

	return handle.(*filestream).h.WriteContent(records)
}

func (c *Client) Write(ctx context.Context, options plugin.WriteOptions, msgs <-chan message.Message) error {
	return c.writer.Write(ctx, msgs)
}

func replacePathVariables(specPath, table string, format filetypes.FormatType, fileIdentifier string, t time.Time) string {
	name := strings.ReplaceAll(specPath, PathVarTable, table)
	name = strings.ReplaceAll(name, PathVarFormat, string(format))
	name = strings.ReplaceAll(name, PathVarUUID, fileIdentifier)
	name = strings.ReplaceAll(name, YearVar, t.Format("2006"))
	name = strings.ReplaceAll(name, MonthVar, t.Format("01"))
	name = strings.ReplaceAll(name, DayVar, t.Format("02"))
	name = strings.ReplaceAll(name, HourVar, t.Format("15"))
	name = strings.ReplaceAll(name, MinuteVar, t.Format("04"))
	return filepath.Clean(name)
}
