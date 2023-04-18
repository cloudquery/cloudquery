package client

import (
	"context"
	"os"
	"path"
	"strings"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/filetypes/v2"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/google/uuid"
)

func (c *Client) WriteTableBatch(ctx context.Context, arrowSchema *arrow.Schema, data []arrow.Record) error {
	tableName := schema.TableName(arrowSchema)
	timeNow := time.Now().UTC()
	p := replacePathVariables(c.pluginSpec.Path, tableName, c.pluginSpec.Format, uuid.NewString(), timeNow)
	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return c.Client.WriteTableBatchFile(f, arrowSchema, data)
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
	return path.Clean(name)
}
