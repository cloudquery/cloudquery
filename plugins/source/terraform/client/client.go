package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	Backends map[string]*TerraformBackend
	logger   zerolog.Logger

	// CurrentBackend set by client multiplexer
	CurrentBackend string
}

func New(logger zerolog.Logger, backends map[string]*TerraformBackend) Client {
	return Client{
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

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	tfSpec := &Spec{}
	if err := s.UnmarshalSpec(tfSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal terraform spec: %w", err)
	}

	if len(tfSpec.Backends) == 0 {
		return nil, errors.New("no backends were provided")
	}

	var backends = make(map[string]*TerraformBackend)
	for _, config := range tfSpec.Backends {
		config := config

		logger.Info().Msg("creating new backend")
		// create backend for each backend config
		b, err := NewBackend(ctx, &config)
		if err != nil {
			return nil, fmt.Errorf("cannot initialize backend: %w", err)
		}
		backends[b.BackendName] = b
	}

	client := New(logger, backends)

	// Returns the initialized client with requested backends
	return &client, nil
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
