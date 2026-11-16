package values

import (
	"context"
	"fmt"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"golang.org/x/sync/errgroup"
)

func BatchAddRecords(ctx context.Context, batch driver.Batch, sc *arrow.Schema, records []arrow.RecordBatch) error {
	fieldIndices, err := schemaFieldIndices(batch, sc)
	if err != nil {
		return err
	}
	table := array.NewTableFromRecords(sc, records)
	eg, _ := errgroup.WithContext(ctx)
	columns := batch.Columns()
	for n, fieldIndex := range fieldIndices {
		column := batch.Column(n)
		if fieldIndex == columnNotInSchema {
			name, rows := columns[n].Name(), table.NumRows()
			eg.Go(func() error { return appendNulls(column, name, rows) })
			continue
		}
		chunks := table.Column(fieldIndex).Data().Chunks()
		eg.Go(func() error {
			for _, chunk := range chunks {
				data, err := FromArray(chunk)
				if err != nil {
					return err
				}
				if err := column.Append(data); err != nil {
					return err
				}
			}
			return nil
		})
	}
	return eg.Wait()
}

const columnNotInSchema = -1

func appendNulls(column driver.BatchColumn, name string, rows int64) error {
	for range rows {
		if err := column.AppendRow(nil); err != nil {
			return fmt.Errorf("column %q is not in the table schema: %w", name, err)
		}
	}
	return nil
}

func schemaFieldIndices(batch driver.Batch, sc *arrow.Schema) ([]int, error) {
	fieldIndexByName := make(map[string]int, sc.NumFields())
	for i, field := range sc.Fields() {
		fieldIndexByName[strings.ReplaceAll(field.Name, "\x00", "")] = i
	}
	columns := batch.Columns()
	fieldIndices := make([]int, len(columns))
	matched := 0
	for n, column := range columns {
		fieldIndex, ok := fieldIndexByName[column.Name()]
		if !ok {
			fieldIndices[n] = columnNotInSchema
			continue
		}
		fieldIndices[n] = fieldIndex
		matched++
	}
	if missing := sc.NumFields() - matched; missing > 0 {
		return nil, fmt.Errorf("table is missing %d schema columns", missing)
	}
	return fieldIndices, nil
}
