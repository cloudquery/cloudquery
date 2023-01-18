package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	destination.DefaultReverseTransformer
	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec
	client     neo4j.DriverWithContext
}

func New(ctx context.Context, logger zerolog.Logger, destSpec specs.Destination) (destination.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "neo4j").Logger(),
		spec:   destSpec,
	}
	var spec Spec
	if err := destSpec.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BigQuery spec: %w", err)
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	c.pluginSpec = spec
	c.client, err = neo4j.NewDriverWithContext(c.pluginSpec.ConnectionString, neo4j.BasicAuth(c.pluginSpec.Username, c.pluginSpec.Password, ""))
	if err != nil {
		return nil, err
	}
	if err := c.client.VerifyConnectivity(ctx); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.client.Close(ctx)
}
