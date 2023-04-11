package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/specs"
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
		return nil, fmt.Errorf("failed to unmarshal neo4j spec: %w", err)
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	c.pluginSpec = spec
	c.client, err = neo4j.NewDriverWithContext(c.pluginSpec.ConnectionString, neo4j.BasicAuth(c.pluginSpec.Username, c.pluginSpec.Password, ""), func(c *neo4j.Config) {
		c.Log = &Logger{Base: logger}
	})
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

func (c *Client) LoggedSession(ctx context.Context, cf neo4j.SessionConfig) neo4j.SessionWithContext {
	if c.logger.GetLevel() <= zerolog.DebugLevel {
		cf.BoltLogger = &Logger{Base: c.logger}
	}
	return c.client.NewSession(ctx, cf)
}
