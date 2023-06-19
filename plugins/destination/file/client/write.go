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
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
)

func (c *Client) WriteTableBatch(_ context.Context, tableName string, upsert bool, msgs []*message.Insert) error {
	if len(msgs) == 0 {
		return nil
	}
	table, err := schema.NewTableFromArrowSchema(msgs[0].Record.Schema())
	if err != nil {
		return err
	}

	timeNow := time.Now().UTC()
	p := replacePathVariables(c.spec.Path, tableName, c.spec.Format, uuid.NewString(), timeNow)

	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	records := make([]arrow.Record, len(msgs))
	for i, msg := range msgs {
		records[i] = msg.Record
	}

	return c.Client.WriteTableBatchFile(f, table, records)
}

func (c *Client) Write(ctx context.Context, options plugin.WriteOptions, msgs <-chan message.Message) error {
	if err := c.writer.Write(ctx, msgs); err != nil {
		return err
	}
	if err := c.writer.Flush(ctx); err != nil {
		return err
	}
	return nil
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
