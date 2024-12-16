package client

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/parquet"
	"github.com/apache/arrow-go/v18/parquet/pqarrow"
	"github.com/cenkalti/backoff/v4"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
)

func nonPkIndices(sc *schema.Table) []int {
	var indices []int
	for i, c := range sc.Columns {
		if !c.PrimaryKey {
			indices = append(indices, i)
		}
	}
	return indices
}

// At time of writing (March 2023), duckdb does not support updating list columns.
// As a workaround, we delete the row and insert it again. This makes it non-atomic, unfortunately,
// but this is unavoidable until support is added to duckdb itself.
// See https://github.com/duckdb/duckdb/blob/c5d9afb97bbf0be12216f3b89ae3131afbbc3156/src/storage/table/list_column_data.cpp#L243-L251
func containsList(sc *schema.Table) bool {
	return slices.ContainsFunc(sc.Columns, func(c schema.Column) bool { return dtContainsList(c.Type) })
}

func dtContainsList(dt arrow.DataType) bool {
	switch dt := dt.(type) {
	case *arrow.StructType:
		return slices.ContainsFunc(dt.Fields(), func(f arrow.Field) bool { return dtContainsList(f.Type) })
	case *arrow.MapType:
		return dtContainsList(dt.KeyType()) || dtContainsList(dt.ItemType())
	case arrow.ListLikeType:
		return true
	default:
		return false
	}
}

func (c *Client) upsert(ctx context.Context, tmpTableName string, table *schema.Table) error {
	var sb strings.Builder
	sb.WriteString("INSERT INTO ")
	sb.WriteString(table.Name)
	sb.WriteString("(" + strings.Join(sanitized(table.Columns.Names()), ", ") + ")")
	sb.WriteString(" SELECT ")
	sb.WriteString(strings.Join(sanitized(table.Columns.Names()), ", "))
	sb.WriteString(" FROM ")
	sb.WriteString(tmpTableName)
	sb.WriteString(" ON CONFLICT (" + strings.Join(table.PrimaryKeys(), ", ") + ")")
	indices := nonPkIndices(table)
	if len(indices) == 0 {
		sb.WriteString(" DO NOTHING")
		return c.exec(ctx, sb.String())
	}

	sb.WriteString(" DO UPDATE SET ")

	written := 0
	for _, index := range nonPkIndices(table) {
		col := table.Columns[index]
		if col.Unique {
			// we skip this stuff, as unique constraint can't be updated by DuckDB
			continue
		}
		if written > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(sanitizeID(col.Name))
		sb.WriteString(" = excluded.")
		sb.WriteString(sanitizeID(col.Name))
		written++
	}
	query := sb.String()

	// return c.exec(ctx, query)
	// per https://duckdb.org/docs/sql/indexes#over-eager-unique-constraint-checking we might need some retries
	// as the upsert for tables with PKs is transformed into delete + insert internally
	return backoff.Retry(
		func() error {
			return c.exec(ctx, query)
		},
		backoff.WithContext(backoff.WithMaxRetries(backoff.NewConstantBackOff(50*time.Millisecond), 3), ctx),
	)
}

func (c *Client) deleteByPK(ctx context.Context, tmpTableName string, table *schema.Table) error {
	var sb strings.Builder
	sb.WriteString("delete from " + table.Name + " using " + tmpTableName + " where ")
	for i, col := range table.PrimaryKeys() {
		if i > 0 {
			sb.WriteString(" and ")
		}
		sb.WriteString(table.Name + "." + sanitizeID(col))
		sb.WriteString(" = ")
		sb.WriteString(tmpTableName + "." + sanitizeID(col))
	}

	return c.exec(ctx, sb.String())
}

func (c *Client) copyFromFile(ctx context.Context, tableName string, fileName string, table *schema.Table) error {
	return c.exec(ctx, "copy "+tableName+
		"("+strings.Join(sanitized(table.Columns.Names()), ", ")+
		") from '"+fileName+"' (FORMAT PARQUET)")
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	if err := c.writer.Write(ctx, msgs); err != nil {
		return fmt.Errorf("failed to write messages: %w", err)
	}
	if err := c.writer.Flush(ctx); err != nil {
		return fmt.Errorf("failed to flush messages: %w", err)
	}
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, name string, msgs message.WriteInserts) error {
	if len(msgs) == 0 {
		return nil
	}

	table := msgs[0].GetTable()

	writeStart := time.Now()
	tmpFile, err := writeTMPFile(table, msgs)
	if err != nil {
		return err
	}
	c.logger.Debug().Str("table", table.Name).Str("duration", time.Since(writeStart).String()).Msg("write tmp file")
	defer os.Remove(tmpFile)

	if len(table.PrimaryKeys()) == 0 {
		copyStart := time.Now()
		defer func() {
			c.logger.Debug().Str("table", table.Name).Str("duration", time.Since(copyStart).String()).Msg("copy file to table")
		}()
		return c.copyFromFile(ctx, name, tmpFile, table)
	}

	tmpTableName := name + strings.ReplaceAll(uuid.New().String(), "-", "_")
	if err := c.createTableIfNotExist(ctx, tmpTableName, table); err != nil {
		return fmt.Errorf("failed to create table %s: %w", tmpTableName, err)
	}
	defer func() {
		e := c.exec(ctx, "drop table "+tmpTableName)
		if err == nil {
			// we preserve original error, so update only on nil err
			err = e
		}
	}()

	if err := c.copyFromFile(ctx, tmpTableName, tmpFile, table); err != nil {
		return fmt.Errorf("failed to copy from file %s: %w", tmpFile, err)
	}

	// At time of writing (March 2023), duckdb does not support updating list columns.
	// As a workaround, we delete the row and insert it again. This makes it non-atomic, unfortunately,
	// but this is unavoidable until support is added to duckdb itself.
	// See https://github.com/duckdb/duckdb/blob/c5d9afb97bbf0be12216f3b89ae3131afbbc3156/src/storage/table/list_column_data.cpp#L243-L251
	if containsList(table) {
		return c.deleteInsert(ctx, tmpTableName, table)
	}

	return c.upsert(ctx, tmpTableName, table)
}

func writeTMPFile(table *schema.Table, msgs []*message.WriteInsert) (fileName string, err error) {
	sc := transformSchemaForWriting(table.ToArrowSchema())

	// create temp file
	f, err := os.CreateTemp("", fmt.Sprintf("%s-*.parquet", table.Name))
	if err != nil {
		return "", err
	}
	defer f.Close() // we don't care here, as the happy-path will actually check the error
	fileName = f.Name()

	// prep file writer
	fw, err := pqarrow.NewFileWriter(sc, f,
		parquet.NewWriterProperties(
			parquet.WithVersion(parquet.V2_LATEST),       // use latest
			parquet.WithMaxRowGroupLength(128*1024*1024), // 128M
			// parquet.WithCompression(compress.Codecs.Snappy),
		),
		pqarrow.NewArrowWriterProperties(pqarrow.WithStoreSchema()),
	)
	if err != nil {
		return "", err
	}
	defer fw.Close() // we don't care here either as the happy path will check the error

	// write records
	for _, msg := range msgs {
		if err = fw.WriteBuffered(transformRecord(sc, msg.Record)); err != nil {
			return "", err
		}
	}

	// close file writer (will close the underlying file, too)
	return fileName, fw.Close()
}

func (c *Client) deleteInsert(ctx context.Context, tmpTableName string, table *schema.Table) error {
	if err := c.deleteByPK(ctx, tmpTableName, table); err != nil {
		return err
	}

	sb := new(strings.Builder)
	sb.WriteString("INSERT INTO ")
	sb.WriteString(table.Name)
	sb.WriteString("(" + strings.Join(sanitized(table.Columns.Names()), ", ") + ")")
	sb.WriteString(" SELECT ")
	sb.WriteString(strings.Join(sanitized(table.Columns.Names()), ", "))
	sb.WriteString(" FROM ")
	sb.WriteString(tmpTableName)
	sb.WriteString(" ON CONFLICT DO NOTHING")
	query := sb.String()

	// per https://duckdb.org/docs/sql/indexes#over-eager-unique-constraint-checking we might need to retry
	return backoff.Retry(
		func() error {
			return c.exec(ctx, query)
		},
		backoff.WithContext(backoff.WithMaxRetries(backoff.NewConstantBackOff(50*time.Millisecond), 3), ctx),
	)
}
