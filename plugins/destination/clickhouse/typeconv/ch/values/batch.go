package values

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/ClickHouse/ch-go/proto"
	chcolumn "github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v8/util"
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
			header, rows := columns[n], table.NumRows()
			eg.Go(func() error { return appendEmptyRows(column, header, rows) })
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

func appendEmptyRows(column driver.BatchColumn, header chcolumn.Interface, rows int64) error {
	value, err := emptyRowValue(header)
	if err != nil {
		return fmt.Errorf("column %q is not in the table schema: %w", header.Name(), err)
	}
	for range rows {
		if err := column.AppendRow(value); err != nil {
			return fmt.Errorf("column %q is not in the table schema: %w", header.Name(), err)
		}
	}
	return nil
}

func emptyRowValue(header chcolumn.Interface) (any, error) {
	scratch, err := scratchColumn(header)
	if err != nil {
		return nil, err
	}
	if scratch.AppendRow(nil) == nil {
		return nil, nil
	}
	if arr, ok := header.(*chcolumn.Array); ok {
		return reflect.MakeSlice(arr.ScanType(), 0, 0).Interface(), nil
	}
	if scratch, err = scratchColumn(header); err != nil {
		return nil, err
	}
	if err := scratch.Decode(proto.NewReader(zeroBytes{}), 1); err != nil {
		return nil, err
	}
	return scratch.Row(0, false), nil
}

func scratchColumn(header chcolumn.Interface) (chcolumn.Interface, error) {
	return header.Type().Column(header.Name(), &chcolumn.ServerContext{Timezone: time.UTC})
}

type zeroBytes struct{}

func (zeroBytes) Read(p []byte) (int, error) {
	clear(p)
	return len(p), nil
}

func schemaFieldIndices(batch driver.Batch, sc *arrow.Schema) ([]int, error) {
	fieldIndexByName := make(map[string]int, sc.NumFields())
	for i, field := range sc.Fields() {
		fieldIndexByName[util.StoredName(field.Name)] = i
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
