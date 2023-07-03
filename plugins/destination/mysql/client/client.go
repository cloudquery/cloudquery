package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/rs/zerolog"

	mysql "github.com/go-sql-driver/mysql"
)

type Client struct {
	plugin.UnimplementedSource
	logger zerolog.Logger
	spec   Spec
	db     *sql.DB
	writer *batchwriter.BatchWriter
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte) (plugin.Client, error) {
	c := &Client{logger: logger.With().Str("module", "mysql").Logger()}
	var err error

	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	c.spec.SetDefaults()
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.writer, err = batchwriter.New(c, batchwriter.WithLogger(logger), batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create batch writer: %w", err)
	}

	dsn, err := mysql.ParseDSN(c.spec.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("invalid MySQL connection string: %w", err)
	}
	if dsn.Params == nil {
		dsn.Params = map[string]string{}
	}
	dsn.Params["parseTime"] = "true"
	db, err := sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql connection: %w", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	c.db = db

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		_ = c.db.Close()
		return fmt.Errorf("failed to close writer: %w", err)
	}
	return c.db.Close()
}
