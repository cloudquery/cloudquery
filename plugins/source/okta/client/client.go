package client

import (
	"github.com/okta/okta-sdk-golang/v3/okta"
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	spec   Spec

	*okta.APIClient
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (Client) ID() string {
	return "okta"
}

func New(logger zerolog.Logger, s Spec, okt *okta.APIClient) *Client {
	return &Client{
		APIClient: okt,

		logger: logger,
		spec:   s,
	}
}
