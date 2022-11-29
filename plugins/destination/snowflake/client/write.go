package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
)

const batchFileSize = uint64(1024 * 1024 * 4) // 4MB

const (
	createOrReplaceFileFormat = `create or replace file format cq_plugin_json_format type = 'JSON'`
	createOrReplaceStage      = `create or replace stage cq_plugin_stage file_format = cq_plugin_json_format;`
	putFileIntoStage          = `put file://%s @cq_plugin_stage auto_compress=true`
	copyIntoTable             = `copy into %s from @cq_plugin_stage/%s file_format = (format_name = cq_plugin_json_format) match_by_column_name = case_insensitive`
)

type worker struct {
	writeChan chan []interface{}
}

func (c *Client) writeResource(ctx context.Context, table *schema.Table, resources <-chan []interface{}) error {
	f, err := os.CreateTemp(os.TempDir(), table.Name+".json.*")
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	totalWritten := uint64(0)
	for r := range resources {
		jsonObj := make(map[string]interface{}, len(table.Columns))
		for i := range r {
			jsonObj[table.Columns[i].Name] = r[i]
		}

		b, err := json.Marshal(jsonObj)
		if err != nil {
			return err
		}
		b = append(b, '\n')
		if _, err := f.Write(b); err != nil {
			return err
		}
		totalWritten += uint64(len(b))
		if totalWritten >= batchFileSize {
			if err := f.Close(); err != nil {
				return err
			}
			sql := fmt.Sprintf(putFileIntoStage, f.Name())
			if _, err := c.db.ExecContext(ctx, sql); err != nil {
				return fmt.Errorf("failed to put file into stage %s: %w", sql, err)
			}
			if err := os.Remove(f.Name()); err != nil {
				return err
			}
			sql = fmt.Sprintf(copyIntoTable, table.Name, path.Base(f.Name()))
			if _, err := c.db.ExecContext(ctx, sql); err != nil {
				return fmt.Errorf("failed to copy file into table %s: %w", sql, err)
			}
			f, err = os.CreateTemp(os.TempDir(), table.Name+".json.*")
			if err != nil {
				return err
			}
			totalWritten = 0
		}
	}
	if totalWritten > 0 {
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
	}

	return err
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *plugins.ClientResource) error {
	eg := errgroup.Group{}
	workers := make(map[string]*worker, len(tables))

	if _, err := c.db.ExecContext(ctx, createOrReplaceFileFormat); err != nil {
		return fmt.Errorf("failed to create file format %s: %w", createOrReplaceFileFormat, err)
	}

	if _, err := c.db.ExecContext(ctx, createOrReplaceStage); err != nil {
		return fmt.Errorf("failed to create stage %s: %w", createOrReplaceStage, err)
	}

	for _, t := range tables.FlattenTables() {
		t := t
		workers[t.Name] = &worker{
			writeChan: make(chan []interface{}),
		}
		eg.Go(func() error {
			return c.writeResource(ctx, t, workers[t.Name].writeChan)
		})
	}

	for r := range res {
		workers[r.TableName].writeChan <- r.Data
	}
	for _, w := range workers {
		close(w.writeChan)
	}

	return eg.Wait()
}
