package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	destination.DefaultReverseTransformer
	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec
	client     *mongo.Client
}

func New(ctx context.Context, logger zerolog.Logger, destSpec specs.Destination) (destination.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "mongo-dest").Logger(),
		spec:   destSpec,
	}

	var spec Spec
	if err := destSpec.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal MongoDB spec: %w", err)
	}
	if err := spec.Validate(); err != nil {
		return nil, err
	}
	c.client, err = mongo.NewClient(options.Client().ApplyURI(spec.ConnectionString))
	if err != nil {
		return nil, err
	}
	if err := c.client.Connect(context.Background()); err != nil {
		return nil, err
	}
	c.pluginSpec = spec

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}
