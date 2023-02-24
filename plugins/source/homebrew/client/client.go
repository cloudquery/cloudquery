package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 5
	defaultBackoff    = 10 * time.Second
)

type Client struct {
	logger     zerolog.Logger
	sourceSpec specs.Source
	Spec       Spec
	Homebrew   *homebrew.Client
	maxRetries int
	backoff    time.Duration // backoff duration between retries (jitter will be added)
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (*Client) ID() string {
	return "homebrew"
}

func Configure(ctx context.Context, logger zerolog.Logger, sourceSpec specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var config Spec
	err := sourceSpec.UnmarshalSpec(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	err = config.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	client := homebrew.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create homebrew client: %w", err)
	}

	return &Client{
		logger:     logger,
		sourceSpec: sourceSpec,
		Spec:       config,
		Homebrew:   client,
		maxRetries: defaultMaxRetries,
		backoff:    defaultBackoff,
	}, nil
}
