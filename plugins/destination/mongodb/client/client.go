package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	plugin.UnimplementedSource
	logger     zerolog.Logger
	spec *Spec
	client     *mongo.Client
	writer     *writers.BatchWriter
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte) (plugin.Client, error) {
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
	c.writer, err = writers.NewBatchWriter(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}
