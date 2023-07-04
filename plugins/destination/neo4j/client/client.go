package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	logger zerolog.Logger
	spec   *Spec
	client neo4j.DriverWithContext
	writer *batchwriter.BatchWriter
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "neo4j").Logger(),
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal neo4j spec: %w", err)
	}
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.spec.SetDefaults()

	var err error
	c.client, err = neo4j.NewDriverWithContext(c.spec.ConnectionString, neo4j.BasicAuth(c.spec.Username, c.spec.Password, ""), func(c *neo4j.Config) {
		c.Log = &Logger{Base: logger}
	})
	if err != nil {
		return nil, err
	}
	if err := c.client.VerifyConnectivity(ctx); err != nil {
		return nil, err
	}

	c.writer, err = batchwriter.New(c, batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		_ = c.client.Close(ctx)
		return fmt.Errorf("failed to close writer: %w", err)
	}
	return c.client.Close(ctx)
}

func (c *Client) LoggedSession(ctx context.Context, cf neo4j.SessionConfig) neo4j.SessionWithContext {
	if c.logger.GetLevel() <= zerolog.DebugLevel {
		cf.BoltLogger = &Logger{Base: c.logger}
	}
	return c.client.NewSession(ctx, cf)
}
