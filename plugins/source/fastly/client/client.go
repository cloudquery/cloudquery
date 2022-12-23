package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/services"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/fastly/go-fastly/v7/fastly"
	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 3
	defaultBackoff    = 1 * time.Second
)

type Client struct {
	logger zerolog.Logger
	spec   specs.Source
	Fastly services.FastlyClient

	maxRetries int
	backoff    time.Duration // backoff duration between retries (jitter will be added)
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.spec.Name
}

func (c *Client) withTeamID(teamID string) schema.ClientMeta {
	return &Client{
		logger:     c.logger.With().Str("team_id", teamID).Logger(),
		spec:       c.spec,
		Fastly:     c.Fastly,
		maxRetries: c.maxRetries,
		backoff:    c.backoff,
	}
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	var config Spec
	err := s.UnmarshalSpec(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	client, err := fastly.NewClient(config.FastlyAPIKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create fastly client: %w", err)
	}
	return &Client{
		logger:     logger,
		spec:       s,
		Fastly:     client,
		maxRetries: defaultMaxRetries,
		backoff:    defaultBackoff,
	}, nil
}
