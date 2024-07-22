package client

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/goccy/go-json"
)

const (
	createOrReplaceFileFormat = `create or replace file format cq_plugin_json_format type = 'JSON'`
	createOrReplaceStage      = `create or replace stage cq_plugin_stage file_format = cq_plugin_json_format;`
	putFileIntoStage          = `put 'file://%v' @cq_plugin_stage auto_compress=true`
	copyIntoTable             = `copy into %s from '@cq_plugin_stage' files=('%s.gz') on_error = ABORT_STATEMENT file_format = (format_name = cq_plugin_json_format) match_by_column_name = case_insensitive`
)

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	if err := c.writer.Write(ctx, msgs); err != nil {
		return err
	}
	if err := c.writer.Flush(ctx); err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, name string, msgs message.WriteInserts) error {
	if err := c.setupWrite(ctx); err != nil {
		return err
	}

	tableName := name
	f, err := os.CreateTemp(os.TempDir(), tableName+".json.*")
	if err != nil {
		return err
	}
	if c.spec.LeaveStageFiles {
		c.logger.Info().Str("filename", f.Name()).Str("table", name).Msg("Created stage file")
	}
	defer func() {
		f.Close()
		if !c.spec.LeaveStageFiles {
			os.Remove(f.Name())
		}
	}()

	enc := json.NewEncoder(f)
	enc.SetEscapeHTML(false)

	for _, r := range msgs {
		arr := array.RecordToStructArray(r.Record)
		for i := 0; i < arr.Len(); i++ {
			if err := enc.Encode(arr.GetOneForMarshal(i)); err != nil {
				return err
			}
		}
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("failed to close temp file with last resource %s: %w", f.Name(), err)
	}

	sql := fmt.Sprintf(putFileIntoStage, escapePath(f.Name()))

	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to put file into stage with last resource %s: %w", sql, err)
	}

	sql = fmt.Sprintf(copyIntoTable, tableName, escapePath(filepath.Base(f.Name())))
	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to copy file into table with last resource %s: %w", sql, err)
	}
	return err
}

func (c *Client) setupWrite(ctx context.Context) error {
	var setupErr error
	c.setupWriteOnce.Do(func() {
		if _, err := c.db.ExecContext(ctx, createOrReplaceFileFormat); err != nil {
			setupErr = fmt.Errorf("failed to create file format %s: %w", createOrReplaceFileFormat, err)
			return
		}
		if _, err := c.db.ExecContext(ctx, createOrReplaceStage); err != nil {
			setupErr = fmt.Errorf("failed to create stage %s: %w", createOrReplaceStage, err)
		}
	})
	return setupErr
}
