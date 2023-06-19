package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/client"
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
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

var tables = []*schema.Table{
	analytics.Installs(homebrew.Days30),
	analytics.Installs(homebrew.Days90),
	analytics.Installs(homebrew.Days365),
	analytics.CaskInstalls(homebrew.Days30),
	analytics.CaskInstalls(homebrew.Days90),
	analytics.CaskInstalls(homebrew.Days365),
	analytics.BuildErrors(homebrew.Days30),
	analytics.BuildErrors(homebrew.Days90),
	analytics.BuildErrors(homebrew.Days365),
}

type Client struct {
	SchedulerClient *client.Client
	logger          zerolog.Logger

	plugin.UnimplementedDestination
}

func (c *Client) GetSpec() any {
	return &client.Spec{}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.Message) error {
	scheduler := scheduler.NewScheduler(c.SchedulerClient, scheduler.WithSchedulerStrategy(scheduler.StrategyDFS))
	return scheduler.Sync(ctx, tables, res)
}

func (c *Client) Tables(ctx context.Context) (schema.Tables, error) {
	return tables, nil
}

func (c *Client) Close(ctx context.Context) error {
	return nil
}

func Configure(ctx context.Context, logger zerolog.Logger, spec []byte) (plugin.Client, error) {
	config := &client.Spec{}
	if err := json.Unmarshal(spec, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	err := config.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	homebrewClient := homebrew.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create homebrew client: %w", err)
	}

	schedulerClient := &client.Client{
		Homebrew:   homebrewClient,
		Logger:     logger,
		Spec:       config,
		MaxRetries: defaultMaxRetries,
		Backoff:    defaultBackoff,
	}

	return &Client{
		logger:          logger,
		SchedulerClient: schedulerClient,
	}, nil
}
