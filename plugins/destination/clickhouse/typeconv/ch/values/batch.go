package values

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"golang.org/x/sync/errgroup"
)

func BatchAddRecords(ctx context.Context, batch driver.Batch, reader array.RecordReader) error {
	for reader.Next() {
		if err := appendRecord(ctx, batch, reader.Record()); err != nil {
			return err
		}
	}
	return nil
}

func appendRecord(ctx context.Context, batch driver.Batch, record arrow.Record) error {
	eg, _ := errgroup.WithContext(ctx)
	for i, col := range record.Columns() {
		i, col := i, col
		eg.Go(func() error {
			if record.Schema().Field(i).Name == `time32s_map` {
				fmt.Println("need to investigate")
			}
			return appendArray(batch.Column(i), col)
		})
	}
	return eg.Wait()
}

func appendArray(column driver.BatchColumn, arr arrow.Array) error {
	data, err := FromArray(arr)
	if err != nil {
		return err
	}

	return column.Append(data)
}
