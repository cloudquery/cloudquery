package values

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"golang.org/x/sync/errgroup"
)

func BatchAddRecords(ctx context.Context, batch driver.Batch, sc *arrow.Schema, records []arrow.Record) error {
	table := array.NewTableFromRecords(sc, records)
	eg, _ := errgroup.WithContext(ctx)
	for n := 0; n < int(table.NumCols()); n++ {
		column, chunks := batch.Column(n), table.Column(n).Data().Chunks()
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
