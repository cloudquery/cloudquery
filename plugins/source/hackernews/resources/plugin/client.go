package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/client"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/resources/services/items"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/hermanschaaf/hackernews"
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
	hnClient := hackernews.NewClient()

	stateClient, err := state.NewConnectedClient(ctx, options.BackendOptions)
	if err != nil {
		return fmt.Errorf("failed to create state client: %w", err)
	}
	defer stateClient.Close()

	schedulerClient, err := client.New(c.logger, c.config, hnClient, stateClient)
	if err != nil {
		return fmt.Errorf("failed to create scheduler client: %w", err)
	}

	schedulerOptions := []scheduler.SyncOption{scheduler.WithSyncDeterministicCQID(options.DeterministicCQID)}
	if options.Shard != nil {
		schedulerOptions = append(schedulerOptions, scheduler.WithShard(options.Shard.Num, options.Shard.Total))
	}
	err = c.scheduler.Sync(ctx, schedulerClient, tt, res, schedulerOptions...)
	if err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}
	return stateClient.Flush(ctx)
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

func getTables() []*schema.Table {
	tables := []*schema.Table{
		items.Items(),
	}

	for i := range tables {
		tables[i].Columns = append([]schema.Column{schema.CqIDColumn, schema.CqParentIDColumn}, tables[i].Columns...)
		err := tables[i].Transform(tables[i])
		if err != nil {
			panic(err)
		}
	}
	return tables
}

func Configure(_ context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger: logger,
			tables: getTables(),
		}, nil
	}
	config := client.Spec{}
	if err := json.Unmarshal(specBytes, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	err := config.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	return &Client{
		config: config,
		logger: logger,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
			scheduler.WithInvocationID(opts.InvocationID),
		),
		tables: getTables(),
	}, nil
}
