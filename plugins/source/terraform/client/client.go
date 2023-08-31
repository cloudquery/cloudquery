package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

type Client struct {
	Backends map[string]*TerraformBackend
	logger   zerolog.Logger

	// CurrentBackend set by client multiplexer
	CurrentBackend string
}

func New(logger zerolog.Logger, backends map[string]*TerraformBackend) schema.ClientMeta {
	return &Client{
		Backends: backends,
		logger:   logger,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.CurrentBackend
}

func (c *Client) Backend() *TerraformBackend {
	if c.CurrentBackend != "" {
		backend := c.Backends[c.CurrentBackend]
		return backend
	}
	for _, backend := range c.Backends {
		return backend
	}
	return nil
}

// Sets the current backend to working with
func (c *Client) withSpecificBackend(backendName string) *Client {
	return &Client{
		Backends:       c.Backends,
		logger:         c.logger,
		CurrentBackend: backendName,
	}
}
