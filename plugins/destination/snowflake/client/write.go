package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

const (
	createOrReplaceFileFormat = `create or replace file format cq_plugin_json_format type = 'JSON'`
	createOrReplaceStage      = `create or replace stage cq_plugin_stage file_format = cq_plugin_json_format;`
	putFileIntoStage          = `put file://%s @cq_plugin_stage auto_compress=true`
	copyIntoTable             = `copy into %s from @cq_plugin_stage/%s file_format = (format_name = cq_plugin_json_format) match_by_column_name = case_insensitive`
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources []arrow.Record) error {
	tableName := table.Name
	f, err := os.CreateTemp(os.TempDir(), tableName+".json.*")
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	for _, r := range resources {
		arr := array.RecordToStructArray(r)
		enc := json.NewEncoder(f)
		enc.SetEscapeHTML(false)
		for i := 0; i < arr.Len(); i++ {
			if err := enc.Encode(arr.GetOneForMarshal(i)); err != nil {
				return err
			}
		}
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("failed to close temp file with last resource %s: %w", f.Name(), err)
	}
	sql := fmt.Sprintf(putFileIntoStage, f.Name())
	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to put file into stage with last resource %s: %w", sql, err)
	}
	sql = fmt.Sprintf(copyIntoTable, tableName, path.Base(f.Name()))
	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to copy file into table with last resource %s: %w", sql, err)
	}
	return err
}
