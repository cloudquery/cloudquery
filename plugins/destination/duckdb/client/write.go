package client

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/duckdb/json"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"github.com/google/uuid"
)

// func (c *Client) upsert(table *schema.Table, data []any) error {
// 	// At time of writing (March 2023), duckdb does not support updating list columns.
// 	// As a workaround, we delete the row and insert it again. This makes it non-atomic, unfortunately,
// 	// but this is unavoidable until support is added to duckdb itself.
// 	// See https://github.com/duckdb/duckdb/blob/c5d9afb97bbf0be12216f3b89ae3131afbbc3156/src/storage/table/list_column_data.cpp#L243-L251
// 	var sb strings.Builder
// 	if len(table.PrimaryKeys()) > 0 {
// 		sb.WriteString("delete from ")
// 		sb.WriteString(`"` + table.Name + `"`)
// 		sb.WriteString(" where ")
// 		pks := table.PrimaryKeys()
// 		pkData := make([]any, len(pks))
// 		for i, k := range pks {
// 			col := table.Columns.Get(k)
// 			sb.WriteString(`"` + col.Name + `"`)
// 			sb.WriteString(" = ")
// 			sb.WriteString(fmt.Sprintf("$%d ", i+1))
// 			if i < len(pks)-1 {
// 				sb.WriteString("and ")
// 			}
// 			pkData[i] = data[table.Columns.Index(k)]
// 		}
// 		sql := sb.String()
// 		if _, err := c.db.Exec(sql, pkData...); err != nil {
// 			return fmt.Errorf("failed to execute '%s': %w", sql, err)
// 		}
// 		sb.Reset()
// 	}
// 	sb.WriteString("insert into ")
// 	c.insertQuery(&sb, table, data)
// 	expanded := expandData(table, data)
// 	sql := sb.String()
// 	if _, err := c.db.Exec(sql, expanded...); err != nil {
// 		return fmt.Errorf("failed to execute '%s': %w", sql, err)
// 	}
// 	return nil
// }

func nonPkIndices(sc *arrow.Schema) []int {
	var indices []int
	for i, f := range sc.Fields() {
		if !schema.IsPk(f) {
			indices = append(indices, i)
		}
	}
	return indices
}

// At time of writing (March 2023), duckdb does not support updating list columns.
// As a workaround, we delete the row and insert it again. This makes it non-atomic, unfortunately,
// but this is unavoidable until support is added to duckdb itself.
// See https://github.com/duckdb/duckdb/blob/c5d9afb97bbf0be12216f3b89ae3131afbbc3156/src/storage/table/list_column_data.cpp#L243-L251
func containsList(sc *arrow.Schema) bool {
	for _, f := range sc.Fields() {
		if f.Type.ID() == arrow.LIST {
			return true
		}
	}
	return false
}

func (c *Client) upsert(tmpTableName string, tableName string, sc *arrow.Schema) error {
	var sb strings.Builder
	sb.WriteString("insert into " + tableName + " select * from " + tmpTableName + " on conflict (")
	pkIndices := schema.PrimaryKeyIndices(sc)
	for i, indice := range pkIndices {
		col := sc.Field(indice)
		sb.WriteString(col.Name)
		if i < len(pkIndices)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(" ) do update set ")
	indices := nonPkIndices(sc)
	for i, indice := range indices {
		col := sc.Field(indice)
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

func (c *Client) delete_by_pk(tmpTableName string, tableName string, sc *arrow.Schema) error {
	var sb strings.Builder
	sb.WriteString("delete from " + tableName + " using " + tmpTableName + " where ")
	pkIndices := schema.PrimaryKeyIndices(sc)
	for i, indice := range pkIndices {
		col := sc.Field(indice)
		sb.WriteString(tableName + "." + col.Name)
		sb.WriteString(" = ")
		sb.WriteString(tmpTableName + "." + col.Name)
		if i < len(pkIndices)-1 {
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
		sb.WriteString(col.Name)
		if i < len(sc.Fields())-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(") from '" + fileName + "' (timestampformat '%Y-%m-%d %H:%M:%S.%f')")
	if _, err := c.db.Exec(sb.String()); err != nil {
		return err
	}
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, sc *arrow.Schema, records []arrow.Record) error {
	tableName := schema.TableName(sc)
	f, err := os.CreateTemp("", fmt.Sprintf("%s-*.json", tableName))
	if err != nil {
		return err
	}
	// defer os.Remove(f.Name())

	w := json.NewWriter(f, sc)
	for _, r := range records {
		if err := w.Write(r); err != nil {
			return err
		}
	}

	if err := f.Close(); err != nil {
		return err
	}

	fmt.Println(f.Name())
	if c.spec.WriteMode == specs.WriteModeAppend || len(schema.PrimaryKeyIndices(sc)) == 0 {
		if err := c.copy_from_file(tableName, f.Name(), sc); err != nil {
			return err
		}
	} else {
		tmpTableName := tableName + strings.ReplaceAll(uuid.New().String(), "-", "_")
		if err := c.createTableIfNotExist(tmpTableName, sc); err != nil {
			return err
		}
		if err := c.copy_from_file(tmpTableName, f.Name(), sc); err != nil {
			return err
		}

		// At time of writing (March 2023), duckdb does not support updating list columns.
		// As a workaround, we delete the row and insert it again. This makes it non-atomic, unfortunately,
		// but this is unavoidable until support is added to duckdb itself.
		// See https://github.com/duckdb/duckdb/blob/c5d9afb97bbf0be12216f3b89ae3131afbbc3156/src/storage/table/list_column_data.cpp#L243-L251
		if containsList(sc) {
			if err := c.delete_by_pk(tmpTableName, tableName, sc); err != nil {
				return err
			}
			if _, err = c.db.Exec("insert into " + tableName + " from " + tmpTableName); err != nil {
				return err
			}
		} else {
			if err := c.upsert(tmpTableName, tableName, sc); err != nil {
				return err
			}
		}
		if _, err = c.db.Exec("drop table " + tmpTableName); err != nil {
			return err
		}
	}

	return nil
}
