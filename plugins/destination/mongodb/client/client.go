package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	plugin.UnimplementedSource
	logger zerolog.Logger
	spec   *Spec
	client *mongo.Client
	writer *batchwriter.BatchWriter
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "mongo-dest").Logger(),
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal MongoDB spec: %w", err)
	}
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.client, err = mongo.NewClient(options.Client().ApplyURI(c.spec.ConnectionString).SetRegistry(getRegistry()))
	if err != nil {
		return nil, err
	}
	if err := c.client.Connect(context.Background()); err != nil {
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
		_ = c.client.Disconnect(ctx)
		return fmt.Errorf("failed to close writer: %w", err)
	}
	return c.client.Disconnect(ctx)
}
