package client

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/filetypes/v4/types"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/google/uuid"
)

func (c *Client) WriteTable(_ context.Context, msgs <-chan *message.WriteInsert) error {
	var (
		f *os.File
		h types.Handle
	)
	for msg := range msgs {
		if f == nil {
			table := msg.GetTable()
			p := c.replacePathVariables(table.Name, uuid.NewString(), time.Now().UTC())
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
	}

	if err := h.WriteFooter(); err != nil {
		return err
	}
	return f.Close()
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	return c.writer.Write(ctx, msgs)
}

func (c *Client) replacePathVariables(table string, fileIdentifier string, t time.Time) string {
	name := strings.ReplaceAll(c.spec.Path, PathVarTable, table)
	if strings.Contains(name, PathVarFormat) {
		e := string(c.spec.Format) + c.spec.Compression.Extension()
		name = strings.ReplaceAll(name, PathVarFormat, e)
	}
	name = strings.ReplaceAll(name, PathVarUUID, fileIdentifier)
	name = strings.ReplaceAll(name, YearVar, t.Format("2006"))
	name = strings.ReplaceAll(name, MonthVar, t.Format("01"))
	name = strings.ReplaceAll(name, DayVar, t.Format("02"))
	name = strings.ReplaceAll(name, HourVar, t.Format("15"))
	name = strings.ReplaceAll(name, MinuteVar, t.Format("04"))
	return filepath.Clean(name)
}
