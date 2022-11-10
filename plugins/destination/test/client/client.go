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
	plugins.DefaultReverseTransformer
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

func (*Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	return nil
}

func (*Client) Migrate(ctx context.Context, tables schema.Tables) error {
	return nil
}

func (*Client) Write(ctx context.Context, tables schema.Tables, res <-chan *plugins.ClientResource) error {
	for range res {
		// do nothing
	}
	return nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}

func (*Client) DeleteStale(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	return nil
}

func (*Client) ReverseTransformValues(table *schema.Table, values []interface{}) (schema.CQTypes, error) {
	return nil, nil
}
