package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"

	"github.com/rs/zerolog"
	sf "github.com/snowflakedb/gosnowflake"
)

type Client struct {
	plugin.UnimplementedSource
	db     *sql.DB
	logger zerolog.Logger
	spec   Spec
	writer *batchwriter.BatchWriter
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "sf-dest").Logger(),
	}

	// Read and parse the spec.
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal snowflake spec: %w", err)
	}
	c.spec.SetDefaults()

	// Validate and generate config.
	b64 := "BASE64"
	cfg, err := c.spec.Config(map[string]*string{
		// Force base64-encoded input/output.
		"BINARY_INPUT_FORMAT":  &b64,
		"BINARY_OUTPUT_FORMAT": &b64,
	})
	if err != nil {
		return nil, fmt.Errorf("parsing spec to snowflake config: %w", err)
	}

	c.writer, err = batchwriter.New(c, batchwriter.WithLogger(logger), batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, err
	}

	dsn, err := sf.DSN(cfg)
	if err != nil {
		return nil, fmt.Errorf("building dsn from spec: %w", err)
	}

	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		return nil, err
	}
	c.db = db

	// Setup.
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
