package client

import (
	"context"

	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	Spec   Spec
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "ID"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	// TODO: Add your client initialization here
	c := Client{
		logger: logger,
		Spec:   *s,
	}

	return c, nil
}
