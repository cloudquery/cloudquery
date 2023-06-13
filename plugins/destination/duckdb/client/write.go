package client

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/parquet"
	"github.com/apache/arrow/go/v13/parquet/pqarrow"
	"github.com/cenkalti/backoff/v4"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
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
	sb.WriteString("(" + strings.Join(table.Columns.Names(), ", ") + ")")
	sb.WriteString(" SELECT ")
	sb.WriteString(strings.Join(table.Columns.Names(), ", "))
	sb.WriteString(" FROM ")
	sb.WriteString(tmpTableName)
	sb.WriteString(" ON CONFLICT (" + strings.Join(table.PrimaryKeys(), ", ") + ")")
	sb.WriteString(" DO UPDATE SET ")

	written := 0
	indices := nonPkIndices(table)
	for _, index := range indices {
		col := table.Columns[index]
		if col.Unique {
			// we skip this stuff, as unique constraint can't be updated by DuckDB
			continue
		}
		if written > 0 {
			sb.WriteString(",\n")
		}
		sb.WriteString(col.Name)
		sb.WriteString(" = excluded.")
		sb.WriteString(col.Name)
		written++
	}

	query := sb.String()
	c.logger.Info().Str("query", query).Msg("upsert")
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
		sb.WriteString(table.Name + "." + col)
		sb.WriteString(" = ")
		sb.WriteString(tmpTableName + "." + col)
	}

	return c.exec(ctx, sb.String())
}

func (c *Client) copyFromFile(ctx context.Context, tableName string, fileName string, sc *arrow.Schema) error {
	var sb strings.Builder
	sb.WriteString("copy " + tableName + "(")
	for i, col := range sc.Fields() {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(sanitizeID(col.Name))
	}
	sb.WriteString(") from '" + fileName + "' (FORMAT PARQUET)")
	return c.exec(ctx, sb.String())
}

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, records []arrow.Record) (err error) {
	tmpFile, err := c.writeTMPFile(table, records)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile)

	sc := transformSchemaForWriting(table.ToArrowSchema())
	if !c.enabledPks() || len(table.PrimaryKeys()) == 0 {
		c.logger.Warn().Str("table", table.Name).Msg("simple copyFromFile")
		return c.copyFromFile(ctx, table.Name, tmpFile, sc)
	}

	tmpTableName := table.Name + strings.ReplaceAll(uuid.New().String(), "-", "_")
	c.logger.Warn().Str("table", table.Name).Msg("copy to tmp table")
	if err := c.createTableIfNotExist(ctx, tmpTableName, table); err != nil {
		return fmt.Errorf("failed to create table %s: %w", tmpTableName, err)
	}
	defer func() {
		e := c.exec(ctx, "drop table "+tmpTableName)
		c.logger.Err(e).Str("table", table.Name).Msg("dropped " + tmpTableName)
		if err == nil {
			// we preserve original error, so update only on nil err
			err = e
		}
	}()
	if err := c.copyFromFile(ctx, tmpTableName, tmpFile, sc); err != nil {
		return fmt.Errorf("failed to copy from file %s: %w", tmpFile, err)
	}
	c.logger.Warn().Str("table", table.Name).Msg("copied to " + tmpTableName)

	// At time of writing (March 2023), duckdb does not support updating list columns.
	// As a workaround, we delete the row and insert it again. This makes it non-atomic, unfortunately,
	// but this is unavoidable until support is added to duckdb itself.
	// See https://github.com/duckdb/duckdb/blob/c5d9afb97bbf0be12216f3b89ae3131afbbc3156/src/storage/table/list_column_data.cpp#L243-L251
	if containsList(table) {
		c.logger.Warn().Str("table", table.Name).Msg("will call deleteInsert " + tmpTableName)
		return c.deleteInsert(ctx, tmpTableName, table)
	}

	c.logger.Warn().Str("table", table.Name).Msg("will call upsert " + tmpTableName)
	return c.upsert(ctx, tmpTableName, table)
}

func (c *Client) writeTMPFile(table *schema.Table, records []arrow.Record) (fileName string, err error) {
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
			parquet.WithVersion(parquet.V2_4),
			parquet.WithMaxRowGroupLength(128*1024*1024), // 128M
		),
		pqarrow.NewArrowWriterProperties(pqarrow.WithStoreSchema()),
	)
	if err != nil {
		return "", err
	}
	defer fw.Close() // we don't care here either as the happy path will check the error

	// write records
	for _, r := range records {
		if err = fw.Write(transformRecord(sc, r)); err != nil {
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
	c.logger.Warn().Str("table", table.Name).Msg("del by PK OK " + tmpTableName)

	sb := new(strings.Builder)
	sb.WriteString("INSERT INTO ")
	sb.WriteString(table.Name)
	sb.WriteString("(" + strings.Join(table.Columns.Names(), ", ") + ")")
	sb.WriteString(" SELECT ")
	sb.WriteString(strings.Join(table.Columns.Names(), ", "))
	sb.WriteString(" FROM ")
	sb.WriteString(tmpTableName)
	sb.WriteString(" ON CONFLICT (" + strings.Join(table.PrimaryKeys(), ", ") + ")")
	sb.WriteString(" DO NOTHING")
	query := sb.String()
	c.logger.Info().Str("table", table.Name).Str("query", query).Msg("will call insert")

	// per https://duckdb.org/docs/sql/indexes#over-eager-unique-constraint-checking we might need to retry
	return backoff.Retry(
		func() error {
			c.logger.Info().Str("table", table.Name).Str("query", query).Msg("calling insert")
			err := c.exec(ctx, query)
			c.logger.Err(err).Str("table", table.Name).Str("query", query).Msg("inserted")
			return err
		},
		backoff.WithContext(backoff.WithMaxRetries(backoff.NewConstantBackOff(50*time.Millisecond), 3), ctx),
	)
}
