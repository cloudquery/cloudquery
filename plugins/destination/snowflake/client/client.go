package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/rs/zerolog"

	_ "github.com/snowflakedb/gosnowflake" // "snowflake" database/sql driver.
)

type Client struct {
	plugin.UnimplementedSource
	db     *sql.DB
	logger zerolog.Logger
	spec   Spec
	writer *batchwriter.BatchWriter
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "sf-dest").Logger(),
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal snowflake spec: %w", err)
	}
	c.spec.SetDefaults()
	c.writer, err = batchwriter.New(c, batchwriter.WithLogger(logger), batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, err
	}
	dsn, err := c.spec.DSN()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("snowflake", dsn+"&BINARY_INPUT_FORMAT=BASE64&BINARY_OUTPUT_FORMAT=BASE64")
	if err != nil {
		return nil, err
	}
	c.db = db
	if _, err := c.db.ExecContext(ctx, createOrReplaceFileFormat); err != nil {
		return nil, fmt.Errorf("failed to create file format %s: %w", createOrReplaceFileFormat, err)
	}
	if _, err := c.db.ExecContext(ctx, createOrReplaceStage); err != nil {
		return nil, fmt.Errorf("failed to create stage %s: %w", createOrReplaceStage, err)
	}
	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	if c.db == nil {
		return fmt.Errorf("client already closed or not initialized")
	}

	if err := c.writer.Close(ctx); err != nil {
		_ = c.db.Close()
		c.db = nil
		return fmt.Errorf("failed to close writer: %w", err)
	}

	err := c.db.Close()
	c.db = nil
	return err
}
