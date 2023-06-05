package client

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/parquet"
	"github.com/apache/arrow/go/v13/parquet/pqarrow"
	backoff "github.com/cenkalti/backoff/v4"
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
	sb.WriteString("insert into " + table.Name + " select * from " + tmpTableName + " on conflict (")
	sb.WriteString(strings.Join(table.PrimaryKeys(), ", "))
	sb.WriteString(" ) do update set ")
	indices := nonPkIndices(table)
	for i, index := range indices {
		if i > 0 {
			sb.WriteString(", ")
		}
		col := table.Columns[index]
		sb.WriteString(col.Name)
		sb.WriteString(" = excluded.")
		sb.WriteString(col.Name)
	}
	return c.exec(ctx, sb.String())
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
	f, err := os.CreateTemp("", fmt.Sprintf("%s-*.parquet", table.Name))
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	defer f.Close()
	sc := transformSchemaForWriting(table.ToArrowSchema())

	props := parquet.NewWriterProperties(
		parquet.WithVersion(parquet.V2_4),
		parquet.WithMaxRowGroupLength(128*1024*1024), // 128M
	)
	arrprops := pqarrow.NewArrowWriterProperties(
		pqarrow.WithStoreSchema(),
	)
	fw, err := pqarrow.NewFileWriter(sc, f, props, arrprops)
	if err != nil {
		return err
	}
	defer fw.Close()

	for _, r := range records {
		transformedRec := transformRecord(sc, r)
		err := fw.Write(transformedRec)
		if err != nil {
			return err
		}
	}
	if err := fw.Close(); err != nil {
		return err
	}

	if !c.enabledPks() || len(table.PrimaryKeys()) == 0 {
		return c.copyFromFile(ctx, table.Name, f.Name(), sc)
	}

	tmpTableName := table.Name + strings.ReplaceAll(uuid.New().String(), "-", "_")
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
	if err := c.copyFromFile(ctx, tmpTableName, f.Name(), sc); err != nil {
		return fmt.Errorf("failed to copy from file %s: %w", f.Name(), err)
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

func (c *Client) deleteInsert(ctx context.Context, tmpTableName string, table *schema.Table) error {
	if err := c.deleteByPK(ctx, tmpTableName, table); err != nil {
		return err
	}

	// per https://duckdb.org/docs/sql/indexes#over-eager-unique-constraint-checking we might need to retry
	return backoff.Retry(
		func() error {
			return c.exec(ctx, "insert into "+table.Name+" from "+tmpTableName)
		},
		backoff.WithContext(backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3), ctx),
	)
}
