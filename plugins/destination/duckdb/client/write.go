package client

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/parquet"
	"github.com/apache/arrow/go/v13/parquet/pqarrow"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/schema"
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
	for _, f := range sc.Columns {
		if f.Type.ID() == arrow.LIST {
			return true
		}
	}
	return false
}

func (c *Client) upsert(tmpTableName string, tableName string, table *schema.Table) error {
	var sb strings.Builder
	sb.WriteString("insert into " + tableName + " select * from " + tmpTableName + " on conflict (")
	pks := table.PrimaryKeys()
	for i, col := range pks {
		sb.WriteString(col)
		if i < len(pks)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(" ) do update set ")
	indices := nonPkIndices(table)
	for i, indice := range indices {
		col := table.Columns[indice]
		sb.WriteString(col.Name)
		sb.WriteString(" = excluded.")
		sb.WriteString(col.Name)
		if i < len(indices)-1 {
			sb.WriteString(", ")
		}
	}
	if _, err := c.db.Exec(sb.String()); err != nil {
		return err
	}
	return nil
}

func (c *Client) delete_by_pk(tmpTableName string, tableName string, table *schema.Table) error {
	var sb strings.Builder
	sb.WriteString("delete from " + tableName + " using " + tmpTableName + " where ")
	pks := table.PrimaryKeys()
	for i, col := range pks {
		sb.WriteString(tableName + "." + col)
		sb.WriteString(" = ")
		sb.WriteString(tmpTableName + "." + col)
		if i < len(pks)-1 {
			sb.WriteString(" and ")
		}
	}
	if _, err := c.db.Exec(sb.String()); err != nil {
		return err
	}
	return nil
}

func (c *Client) copy_from_file(tableName string, fileName string, sc *arrow.Schema) error {
	var sb strings.Builder
	sb.WriteString("copy " + tableName + "(")
	for i, col := range sc.Fields() {
		sb.WriteString("\"" + col.Name + "\"")
		if i < len(sc.Fields())-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(") from '" + fileName + "' (FORMAT PARQUET)")
	if _, err := c.db.Exec(sb.String()); err != nil {
		return err
	}
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, records []arrow.Record) error {
	f, err := os.CreateTemp("", fmt.Sprintf("%s-*.parquet", table.Name))
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	sc := table.ToArrowSchema()

	props := parquet.NewWriterProperties(
		parquet.WithVersion(parquet.V2_4),
		// parquet.WithMaxRowGroupLength(128*1024*1024), // 128M
		// parquet.WithCompression(compress.Codecs.Snappy),
	)
	arrprops := pqarrow.NewArrowWriterProperties(
		pqarrow.WithStoreSchema(),
	)

	fw, err := pqarrow.NewFileWriter(sc, f, props, arrprops)
	if err != nil {
		return err
	}

	// w := json.NewWriter(f, sc)
	for _, r := range records {
		err := fw.Write(r)
		if err != nil {
			return err
		}
		// if err := w.Write(r); err != nil {
		// 	return err
		// }
	}
	if err := fw.Close(); err != nil {
		return err
	}
	// if err := f.Close(); err != nil {
	// 	return err
	// }

	if c.spec.WriteMode == specs.WriteModeAppend || len(table.PrimaryKeys()) == 0 {
		if err := c.copy_from_file(table.Name, f.Name(), sc); err != nil {
			return err
		}
	} else {
		tmpTableName := table.Name + strings.ReplaceAll(uuid.New().String(), "-", "_")
		if err := c.createTableIfNotExist(tmpTableName, table); err != nil {
			return fmt.Errorf("failed to create table %s: %w", tmpTableName, err)
		}
		if err := c.copy_from_file(tmpTableName, f.Name(), sc); err != nil {
			return fmt.Errorf("failed to copy from file %s: %w", f.Name(), err)
		}

		// At time of writing (March 2023), duckdb does not support updating list columns.
		// As a workaround, we delete the row and insert it again. This makes it non-atomic, unfortunately,
		// but this is unavoidable until support is added to duckdb itself.
		// See https://github.com/duckdb/duckdb/blob/c5d9afb97bbf0be12216f3b89ae3131afbbc3156/src/storage/table/list_column_data.cpp#L243-L251
		if containsList(table) {
			if err := c.delete_by_pk(tmpTableName, table.Name, table); err != nil {
				return err
			}
			if _, err = c.db.Exec("insert into " + table.Name + " from " + tmpTableName); err != nil {
				return err
			}
		} else {
			if err := c.upsert(tmpTableName, table.Name, table); err != nil {
				return err
			}
		}
		if _, err = c.db.Exec("drop table " + tmpTableName); err != nil {
			return err
		}
	}

	return nil
}
