package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/services/bss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/services/ecs"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/services/oss"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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

	schedulerClient, err := client.New(c.logger, c.config)
	if err != nil {
		return err
	}

	return c.scheduler.Sync(ctx, schedulerClient, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
}

func (c *Client) Tables(ctx context.Context, options plugin.TableOptions) (schema.Tables, error) {
	return c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
}

func (*Client) Close(ctx context.Context) error {
	return nil
}
func getTables() []*schema.Table {
	tables := []*schema.Table{
		bss.BillOverview(),
		bss.Bill(),
		bss.BillDetails(),
		ecs.Instances(),
		oss.Buckets(),
	}
	err := transformers.TransformTables(tables)
	if err != nil {
		panic(err)
	}
	for i := range tables {
		schema.AddCqIDs(tables[i])
	}
	return tables
}

func Configure(ctx context.Context, logger zerolog.Logger, spec []byte, options plugin.NewClientOptions) (plugin.Client, error) {
	if options.NoConnection {
		return &Client{
			logger: logger,
			tables: getTables(),
		}, nil
	}

	config := client.Spec{}
	if err := json.Unmarshal(spec, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	err := config.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}
	sc := scheduler.NewScheduler(
		scheduler.WithLogger(logger),
		scheduler.WithConcurrency(config.Concurrency),
	)
	return &Client{
		logger:                   logger,
		config:                   config,
		tables:                   getTables(),
		scheduler:                sc,
		UnimplementedDestination: plugin.UnimplementedDestination{},
	}, nil
}
