package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client/services"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/hermanschaaf/hackernews"
	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 5
	defaultBackoff    = 10 * time.Second
)

type Client struct {
	logger     zerolog.Logger
	spec       specs.Source
	HackerNews services.HackernewsClient
	Spec       Spec
	maxRetries int
	backoff    time.Duration // backoff duration between retries (jitter will be added)
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.spec.Name
}

func Configure(ctx context.Context, logger zerolog.Logger, sourceSpec specs.Source) (schema.ClientMeta, error) {
	var config Spec
	err := sourceSpec.UnmarshalSpec(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	client := hackernews.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create hackernews client: %w", err)
	}

	return &Client{
		logger:     logger,
		spec:       sourceSpec,
		HackerNews: client,
		maxRetries: defaultMaxRetries,
		backoff:    defaultBackoff,
	}, nil
}
