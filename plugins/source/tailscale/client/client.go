package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

type Client struct {
	*tailscale.Client
	tailnet string
	logger  zerolog.Logger
}

var _ schema.ClientMeta = (*Client)(nil)

func (c *Client) ID() string {
	return c.tailnet
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func Configure(_ context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	tsSpec := new(Spec)
	err := spec.UnmarshalSpec(tsSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	client, err := tsSpec.getClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create Tailscale client: %w", err)
	}
	return &Client{
		Client:  client,
		tailnet: tsSpec.Tailnet,
		logger:  logger,
	}, nil
}
