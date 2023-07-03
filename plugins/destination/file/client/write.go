package client

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/filetypes/v4"
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

			p := replacePathVariables(c.spec.Path, table.Name, c.spec.Format, uuid.NewString(), time.Now().UTC())

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
