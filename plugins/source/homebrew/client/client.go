package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
	homebrewPlugin "github.com/cloudquery/cloudquery/plugins/source/homebrew/resources/plugin"
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/resources/services/analytics"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 5
	defaultBackoff    = 10 * time.Second
)

type Client struct {
	logger     zerolog.Logger
	Spec       *Spec
	Homebrew   *homebrew.Client
	maxRetries int
	backoff    time.Duration // backoff duration between retries (jitter will be added)

	plugin.UnimplementedDestination
}

func (*Client) ID() string {
	return "homebrew"
}

func (c *Client) GetSpec() any {
	return &Spec{}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.Message) error {
	tables, _ := tables(ctx)
	scheduler := scheduler.NewScheduler(tables, c, scheduler.WithSchedulerStrategy(scheduler.StrategyDFS))
	return scheduler.Sync(ctx, res)
}

func (c *Client) Tables(ctx context.Context) (schema.Tables, error) {
	return tables(ctx)
}

func (c *Client) Close(ctx context.Context) error {
	return nil
}

func Configure(ctx context.Context, logger zerolog.Logger, spec any) (plugin.Client, error) {
	config := spec.(*Spec)
	config.SetDefaults()
	err := config.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	client := homebrew.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create homebrew client: %w", err)
	}

	return &Client{
		logger:     logger,
		Spec:       config,
		Homebrew:   client,
		maxRetries: defaultMaxRetries,
		backoff:    defaultBackoff,
	}, nil
}

func tables(ctx context.Context) (schema.Tables, error) {
	return []*schema.Table{
		analytics.Installs(homebrew.Days30),
		analytics.Installs(homebrew.Days90),
		analytics.Installs(homebrew.Days365),
		analytics.CaskInstalls(homebrew.Days30),
		analytics.CaskInstalls(homebrew.Days90),
		analytics.CaskInstalls(homebrew.Days365),
		analytics.BuildErrors(homebrew.Days30),
		analytics.BuildErrors(homebrew.Days90),
		analytics.BuildErrors(homebrew.Days365),
	}, nil
}

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"homebrew",
		homebrewPlugin.Version,
		Configure,
	)
}
