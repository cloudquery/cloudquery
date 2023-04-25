package values

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v12/arrow"
	"golang.org/x/sync/errgroup"
)

func BatchAddRecords(ctx context.Context, sc *arrow.Schema, batch driver.Batch, records []arrow.Record) error {
	// all records conform to the same schema
	eg, _ := errgroup.WithContext(ctx)
	for i := range sc.Fields() {
		arrays := nthArray(records, i)
		column := batch.Column(i)
		eg.Go(func() error {
			return batchArrays(column, arrays)
		})
	}
	return eg.Wait()
}

func batchArrays(column driver.BatchColumn, arrays []arrow.Array) error {
	for _, arr := range arrays {
		data, err := FromArray(arr)
		if err != nil {
			return err
		}

		if err := column.Append(data); err != nil {
			return err
		}
	}
	return nil
}

func nthArray(records []arrow.Record, n int) []arrow.Array {
	res := make([]arrow.Array, len(records))
	for i, record := range records {
		res[i] = record.Column(n)
	}
	return res
}
