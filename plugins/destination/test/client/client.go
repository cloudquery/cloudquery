package client

import (
	"context"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Client struct {
	logger zerolog.Logger
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (plugins.DestinationClient, error) {
	return &Client{
		logger: log.With().Str("module", "test").Logger(),
	}, nil
}

func (*Client) Metrics() plugins.DestinationMetrics {
	return plugins.DestinationMetrics{}
}

func (*Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- *schema.DestinationResource) error {
	return nil
}

func (*Client) Migrate(ctx context.Context, tables schema.Tables) error {
	return nil
}

func (*Client) Write(ctx context.Context, table schema.Tables, res <-chan *schema.DestinationResource) error {
	return nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}

func (*Client) DeleteStale(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	return nil
}
