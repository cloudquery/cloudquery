package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	Interface
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

func Configure(_ context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
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
		Interface: client,
		tailnet:   tsSpec.Tailnet,
		logger:    logger,
	}, nil
}
