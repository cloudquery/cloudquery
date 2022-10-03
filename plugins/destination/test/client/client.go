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

func (*Client) Migrate(ctx context.Context, tables schema.Tables) error {
	return nil
}

func (*Client) Write(ctx context.Context, table string, data map[string]interface{}) error {
	return nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}

func (*Client) DeleteStale(ctx context.Context, tables string, sourceName string, syncTime time.Time) error {
	return nil
}
