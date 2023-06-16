package values

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"golang.org/x/sync/errgroup"
)

func BatchAddRecords(ctx context.Context, batch driver.Batch, table *schema.Table, records []arrow.Record) error {
	eg, _ := errgroup.WithContext(ctx)
	for n := range table.Columns {
		n := n
		eg.Go(func() error {
			column := batch.Column(n)
			for i := range records {
				data, err := FromArray(records[i].Column(n))
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
