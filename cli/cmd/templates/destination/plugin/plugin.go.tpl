package plugin

import (
	"fmt"
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "Development"
)

type Client struct {
	logger zerolog.Logger
}

func New() * Client {
	return &Client{
		logger: log.With().Str("module", "{{.Name}}").Logger(),
	}
}

func (*Client) Name() string {
	return "{{.Name}}"
}

func (*Client) Version() string {
	return Version
}

func (p *Client) Initialize(ctx context.Context, spec specs.Destination) error {
	return fmt.Errorf("not implemented")
}

func (p *Client) SetLogger(logger zerolog.Logger) {
	p.logger = logger
}

func (p *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	return fmt.Errorf("not implemented")
}

func (p *Client) Write(ctx context.Context, table string, data map[string]interface{}) error {
	return fmt.Errorf("not implemented")
}