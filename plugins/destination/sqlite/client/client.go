package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/rs/zerolog"

	// Import sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	plugin.UnimplementedSource

	writer *batchwriter.BatchWriter
	db     *sql.DB
	logger zerolog.Logger
	spec   Spec
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "sqlite-dest").Logger(),
	}

	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	c.spec.SetDefaults()
	var err error
	c.writer, err = batchwriter.New(c, batchwriter.WithLogger(c.logger), batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create batch writer: %w", err)
	}

	db, err := sql.Open("sqlite3", c.spec.ConnectionString)
	if err != nil {
		return nil, err
	}
	c.db = db
	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	var err error
	if c.db == nil {
		return errors.New("client already closed or not initialized")
	}
	err = c.db.Close()
	c.db = nil
	return err
}
