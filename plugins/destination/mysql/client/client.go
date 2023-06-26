package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers"
	"github.com/rs/zerolog"

	mysql "github.com/go-sql-driver/mysql"
)

type Client struct {
	plugin.UnimplementedSource
	logger zerolog.Logger
	spec Spec
	db *sql.DB
	writer *writers.BatchWriter
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte) (plugin.Client, error) {
	c :=  &Client{logger: logger.With().Str("module", "mysql").Logger()}
	var err error
	
	if err := json.Unmarshal(spec, &c.spec) ; err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	c.spec.SetDefaults()
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.writer, err = writers.NewBatchWriter(c, writers.WithLogger(logger), writers.WithBatchSize(c.spec.BatchSize), writers.WithBatchSizeBytes(c.spec.BatchSizeBytes))
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

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.db.Close()
}
