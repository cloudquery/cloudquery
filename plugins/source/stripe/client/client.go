package client

import (
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
	"github.com/stripe/stripe-go/v74/client"
)

type Client struct {
	logger zerolog.Logger
	spec   Spec

	Services *client.API
	Backend  state.Client
}

func New(logger zerolog.Logger, spec Spec, services *client.API, bk state.Client) *Client {
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
	return "stripe"
}
