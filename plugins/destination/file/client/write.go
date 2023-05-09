package client

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/filetypes/v3"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination/batchingwriter"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/google/uuid"
)

var _ batchingwriter.OpenCloseWriter = (*Client)(nil)

func (c *Client) OpenTable(_ context.Context, sourceSpec specs.Source, table *schema.Table, syncTime time.Time) error {
	c.logger.Debug().Str("source", sourceSpec.Name).Str("table", table.Name).Msg("OpenTable")

	c.tableFilesMu.Lock()
	defer c.tableFilesMu.Unlock()

	p := replacePathVariables(c.pluginSpec.Path, table.Name, c.pluginSpec.Format, uuid.NewString(), syncTime)

	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	h, err := c.Client.WriteHeader(f, table)
	if err != nil {
		return err
	}

	key := mapKey(sourceSpec, table)
	c.tableFiles[key] = writer{
		f: f,
		h: h,
	}
	c.lastSpec = sourceSpec

	return nil
}

func (c *Client) CloseTable(_ context.Context, sourceSpec specs.Source, table *schema.Table, _ time.Time) error {
	c.logger.Debug().Str("source", sourceSpec.Name).Str("table", table.Name).Msg("CloseTable")

	c.tableFilesMu.Lock()
	defer c.tableFilesMu.Unlock()

	key := mapKey(sourceSpec, table)
	err := c.tableFiles[key].h.WriteFooter()
	err2 := c.tableFiles[key].f.Close()
	delete(c.tableFiles, key)
	if err != nil {
		return err
	}
	return err2
}

func (c *Client) WriteTableBatch(_ context.Context, sourceSpec specs.Source, table *schema.Table, _ time.Time, data []arrow.Record) error {
	c.tableFilesMu.RLock()
	key := mapKey(sourceSpec, table)
	h := c.tableFiles[key].h
	c.tableFilesMu.RUnlock()

	return h.WriteContent(data)
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

func mapKey(sourceSpec specs.Source, table *schema.Table) string {
	return sourceSpec.Name + ":" + table.Name
}
