package client

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/goccy/go-json"
)

const (
	createOrReplaceFileFormat = `create or replace file format cq_plugin_json_format type = 'JSON'`
	createOrReplaceStage      = `create or replace stage cq_plugin_stage file_format = cq_plugin_json_format;`
	putFileIntoStage          = `put 'file://%v' @cq_plugin_stage auto_compress=true`
	copyIntoTable             = `copy into %s from '@cq_plugin_stage' files=('%s.gz') on_error = ABORT_STATEMENT file_format = (format_name = cq_plugin_json_format) match_by_column_name = case_insensitive`
	mergeIntoTable            = `MERGE INTO %s AS dest USING (SELECT %s FROM @cq_plugin_stage/%s.gz t) AS source ON %s WHEN MATCHED THEN %s WHEN NOT MATCHED THEN %s`
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

	table := msgs[0].GetTable()

	// if the table has primary keys, use MERGE in order to upsert rather than just append
	if len(table.PrimaryKeys()) > 0 {
		return c.mergeIntoTable(ctx, table, f)
	}

	return c.copyIntoTable(ctx, table, f)
}

func (c *Client) copyIntoTable(ctx context.Context, table *schema.Table, f *os.File) error {
	sql := fmt.Sprintf(copyIntoTable, table.Name, escapePath(filepath.Base(f.Name())))

	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to copy file into table with last resource %s: %w", sql, err)
	}

	return nil
}

func (c *Client) mergeIntoTable(ctx context.Context, table *schema.Table, f *os.File) error {
	// https://docs.snowflake.com/en/sql-reference/sql/merge#syntax
	sql := fmt.Sprintf(mergeIntoTable, table.Name, createColumnsList(table), escapePath(filepath.Base(f.Name())), createPrimaryKeyList(table), updateColumnsList(table), insertColumnsList(table))

	if _, err := c.db.ExecContext(ctx, sql); err != nil {
		return fmt.Errorf("failed to merge file into table: %s: %w", sql, err)
	}

	return nil
}

func createColumnsList(table *schema.Table) string {
	// creates a string like:
	// $1:COL1::TEXT as COL1, $1:COL2::NUMBER as COL2, $1:COL3::TIMESTAMP_TZ as COL3
	columns := make([]string, 0, len(table.Columns))
	for _, col := range table.Columns {
		columns = append(columns, fmt.Sprintf("$1:%s::%s as %s", col.Name, SchemaTypeToSnowflake(col.Type), col.Name))
	}
	return strings.Join(columns, ",")
}

func createPrimaryKeyList(table *schema.Table) string {
	// creates a string like:
	// source.COL1=dest.COL1 AND source.COL2=dest.COL2 AND source.COL3=dest.COL3
	columns := make([]string, 0, len(table.PrimaryKeys()))
	for _, col := range table.PrimaryKeys() {
		columns = append(columns, fmt.Sprintf("source.%s=dest.%s", col, col))
	}
	return strings.Join(columns, " AND ")
}

func updateColumnsList(table *schema.Table) string {
	// creates an update string like:
	// UPDATE SET COL1=source.COL1, COL2=source.COL2, COL3=source.COL3
	columns := make([]string, 0, len(table.Columns))
	for _, col := range table.Columns {
		columns = append(columns, fmt.Sprintf("%s=source.%s", strings.ToUpper(col.Name), strings.ToUpper(col.Name)))
	}
	return fmt.Sprintf(" UPDATE SET %s ", strings.Join(columns, ","))
}

func insertColumnsList(table *schema.Table) string {
	// creates a string like:
	// INSERT (COL1, COL2, COL3) VALUES (source.COL1, source.COL2, source.COL3)
	names, values := make([]string, 0, len(table.Columns)), make([]string, 0, len(table.Columns))
	for _, col := range table.Columns {
		names = append(names, strings.ToUpper(col.Name))
		values = append(values, "source."+strings.ToUpper(col.Name))
	}

	return fmt.Sprintf("INSERT (%s) VALUES (%s)", strings.Join(names, ", "), strings.Join(values, ", "))
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
