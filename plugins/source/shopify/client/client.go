package client

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	spec   Spec

	Services *shopify.Client
	Backend  state.Client
}

func New(logger zerolog.Logger, spec Spec, services *shopify.Client, bk state.Client) *Client {
	return &Client{
		logger:   logger,
		spec:     spec,
		Services: services,
		Backend:  bk,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (Client) ID() string {
	return "shopify"
}
