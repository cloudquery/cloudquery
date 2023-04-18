package client

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

const (
	createOrReplaceFileFormat = `create or replace file format cq_plugin_json_format type = 'JSON'`
	createOrReplaceStage      = `create or replace stage cq_plugin_stage file_format = cq_plugin_json_format;`
	putFileIntoStage          = `put file://%s @cq_plugin_stage auto_compress=true`
	copyIntoTable             = `copy into %s from @cq_plugin_stage/%s file_format = (format_name = cq_plugin_json_format) match_by_column_name = case_insensitive`
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources []arrow.Record) error {
	f, err := os.CreateTemp(os.TempDir(), table.Name+".json.*")
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	for _, r := range resources {
		b, err := r.MarshalJSON()
		if err != nil {
			return err
		}
		// b = append(b, '\n')
		if _, err := f.Write(b); err != nil {
			return err
		}
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("failed to close temp file with last resource %s: %w", f.Name(), err)
	}
	sql := fmt.Sprintf(putFileIntoStage, f.Name())
	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to put file into stage with last resource %s: %w", sql, err)
	}
	sql = fmt.Sprintf(copyIntoTable, table.Name, path.Base(f.Name()))
	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to copy file into table with last resource %s: %w", sql, err)
	}
	return err
}
