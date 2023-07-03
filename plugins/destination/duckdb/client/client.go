package client

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/rs/zerolog"

	// import duckdb driver
	"github.com/marcboeker/go-duckdb"
)

type Client struct {
	plugin.UnimplementedSource
	db        *sql.DB
	connector driver.Connector
	logger    zerolog.Logger
	spec      Spec
	writer    *batchwriter.BatchWriter
}

var _ plugin.Client = (*Client)(nil)

func New(ctx context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "duckdb-dest").Logger(),
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	c.spec.SetDefaults()
	c.writer, err = batchwriter.New(c, batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create batch writer: %w", err)
	}
	c.connector, err = duckdb.NewConnector(c.spec.ConnectionString, nil)
	if err != nil {
		return nil, err
	}

	c.db = sql.OpenDB(c.connector)

	err = c.exec(ctx, "INSTALL 'json'; LOAD 'json';")
	if err != nil {
		return nil, err
	}
	err = c.exec(ctx, "INSTALL 'parquet'; LOAD 'parquet';")
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	var err error

	if c.db == nil {
		return fmt.Errorf("client already closed or not initialized")
	}

	if err := c.writer.Close(ctx); err != nil {
		_ = c.db.Close()
		c.db = nil
		return fmt.Errorf("failed to close writer: %w", err)
	}

	err = c.db.Close()
	c.db = nil
	return err
}

func (c *Client) exec(ctx context.Context, query string, args ...any) error {
	r, err := c.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	if c.spec.Debug {
		rowsAffected, rowsErr := r.RowsAffected()
		if rowsErr == nil {
			c.logger.Debug().Str("query", query).Any("values", args).Int64("rowsAffected", rowsAffected).Msg("exec query")
		} else {
			c.logger.Debug().Str("query", query).Any("values", args).Err(rowsErr).Msg("exec query")
		}
	}
	return nil
}
