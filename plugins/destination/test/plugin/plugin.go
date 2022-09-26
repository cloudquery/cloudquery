package plugin

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	Version = "Development"
)

type Client struct {
	logger zerolog.Logger
}

func New() *Client {
	return &Client{
		logger: log.With().Str("module", "test").Logger(),
	}
}

func (*Client) Name() string {
	return "test"
}

func (*Client) Version() string {
	return Version
}

func (*Client) Initialize(ctx context.Context, spec specs.Destination) error {
	return nil
}

func (p *Client) SetLogger(logger zerolog.Logger) {
	p.logger = logger
}

func (*Client) Migrate(ctx context.Context, tables schema.Tables) error {
	return nil
}

func (*Client) Write(ctx context.Context, table string, data map[string]interface{}) error {
	return nil
}
