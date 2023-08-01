package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"

	"github.com/rs/zerolog"
)

type Client struct {
	logger    zerolog.Logger
	config    client.Spec
	tables    schema.Tables
	scheduler *scheduler.Scheduler

	plugin.UnimplementedDestination
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	if options.BackendOptions != nil {
		c.logger.Warn().Msg("State backend not supported in plugin, skipping")
	}

	schedulerClient, err := client.New(ctx, c.logger, c.config)
	if err != nil {
		return err
	}

	return c.scheduler.Sync(ctx, schedulerClient, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tt, nil
}

func (*Client) Close(_ context.Context) error {
	return nil
}

func Configure(_ context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger: logger,
			tables: getTables(),
		}, nil
	}

	config := &client.Spec{}
	if err := json.Unmarshal(specBytes, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	return &Client{
		config: *config,
		logger: logger,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
			scheduler.WithConcurrency(config.Concurrency),
		),
		tables: getTables(),
	}, nil
}
