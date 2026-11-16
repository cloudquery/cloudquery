package values

import (
	"context"
	"fmt"

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
	for n, fieldIndex := range fieldIndices {
		column, chunks := batch.Column(n), table.Column(fieldIndex).Data().Chunks()
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

func schemaFieldIndices(batch driver.Batch, sc *arrow.Schema) ([]int, error) {
	columns := batch.Columns()
	fieldIndices := make([]int, len(columns))
	for n, column := range columns {
		indices := sc.FieldIndices(column.Name())
		if len(indices) == 0 {
			return nil, fmt.Errorf("column %q is not in the table schema", column.Name())
		}
		fieldIndices[n] = indices[0]
	}
	return fieldIndices, nil
}
