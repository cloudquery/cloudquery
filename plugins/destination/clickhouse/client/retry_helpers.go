package client

import (
	"context"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	retry "github.com/avast/retry-go/v4"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/queries"
	arrowvalues "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/typeconv/arrow/values"
	chvalues "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/typeconv/ch/values"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

func getRetryOptions(logger zerolog.Logger, query string) []retry.Option {
	onRetryFunc := func(logger zerolog.Logger, query string) func(n uint, err error) {
		return func(attempt uint, err error) {
			logger.Warn().Err(err).Uint("attempt", attempt).Msgf("Retrying query %s", query)
		}
	}

	commonRetryOptions := []retry.Option{
		retry.Attempts(5),
		retry.Delay(3 * time.Second),
		retry.MaxJitter(1 * time.Second),
		retry.LastErrorOnly(true),
		retry.RetryIf(func(err error) bool {
			return strings.Contains(err.Error(), "Too many simultaneous queries")
		}),
	}
	return append(commonRetryOptions, retry.OnRetry(onRetryFunc(logger, query)))
}

func retryQueryRowAndScan(ctx context.Context, logger zerolog.Logger, conn clickhouse.Conn, query string, args []any, dest []any) error {
	err := retry.Do(
		func() error {
			return conn.QueryRow(ctx, query, args...).Scan(dest...)
		},
		getRetryOptions(logger, query)...,
	)

	return err
}

func retryExec(ctx context.Context, logger zerolog.Logger, conn clickhouse.Conn, query string, args ...any) error {
	err := retry.Do(
		func() error {
			return conn.Exec(ctx, query, args...)
		},
		getRetryOptions(logger, query)...,
	)
	return err
}

func retryBatchSend(ctx context.Context, logger zerolog.Logger, conn clickhouse.Conn, table *schema.Table, records []arrow.Record) error {
	err := retry.Do(
		func() error {
			batch, err := conn.PrepareBatch(ctx, queries.Insert(table))
			if err != nil {
				return err
			}

			if err := chvalues.BatchAddRecords(ctx, batch, table.ToArrowSchema(), records); err != nil {
				_ = batch.Abort()
				return err
			}

			return batch.Send()
		},
		getRetryOptions(logger, "batch.Send()")...,
	)
	return err
}

func retryGetTableDefinitions(ctx context.Context, logger zerolog.Logger, database string, conn clickhouse.Conn, messages message.WriteMigrateTables) (schema.Tables, error) {
	schemas, err := retry.DoWithData(
		func() (schema.Tables, error) {
			const flattenNested0 = "SET flatten_nested = 0"
			if err := conn.Exec(ctx, flattenNested0); err != nil {
				return nil, err
			}

			query, params := queries.GetTablesSchema(database)
			rows, err := conn.Query(ctx, query, params...)
			if err != nil {
				return nil, err
			}
			defer rows.Close()

			return queries.ScanTableSchemas(rows, messages)
		},
		getRetryOptions(logger, "getTableDefinitions")...,
	)

	return schemas, err
}

func retryRead(ctx context.Context, logger zerolog.Logger, conn clickhouse.Conn, table *schema.Table) (arrow.Record, error) {
	record, err := retry.DoWithData(
		func() (arrow.Record, error) {
			rows, err := conn.Query(ctx, queries.Read(table))
			if err != nil {
				return nil, err
			}
			defer rows.Close()

			columnsTypes := rows.ColumnTypes()
			builder := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
			for rows.Next() {
				row := rowArr(columnsTypes)
				if err = rows.Scan(row...); err != nil {
					return nil, err
				}

				if err = arrowvalues.AppendToRecordBuilder(builder, row); err != nil {
					return nil, err
				}
			}

			return builder.NewRecord(), nil
		},
		getRetryOptions(logger, "read")...,
	)
	return record, err
}
