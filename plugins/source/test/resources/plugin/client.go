package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/test/client"
	"github.com/cloudquery/cloudquery/plugins/source/test/resources/services"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

type Client struct {
	logger    zerolog.Logger
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
	return c.scheduler.Sync(ctx, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
}

func (c *Client) Tables(context.Context) (schema.Tables, error) {
	return c.tables, nil
}

func (*Client) Close(_ context.Context) error {
	return nil
}

func Configure(_ context.Context, logger zerolog.Logger, spec []byte) (plugin.Client, error) {
	config := &client.Spec{}
	if err := json.Unmarshal(spec, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	schedulerClient := &client.TestClient{
		Logger: logger,
		Spec:   *config,
	}

	return &Client{
		logger: logger,
		scheduler: scheduler.NewScheduler(schedulerClient,
			scheduler.WithLogger(logger),
		),
		tables: getTables(),
	}, nil
}

func getTables() schema.Tables {
	tables := schema.Tables{
		services.TestSomeTable(),
		services.TestDataTable(),
	}
	for i := range tables {
		tables[i].Columns = append([]schema.Column{schema.CqIDColumn, schema.CqParentIDColumn}, tables[i].Columns...)
		if tables[i].Transform != nil {
			if err := tables[i].Transform(tables[i]); err != nil {
				panic(err)
			}
		}
	}
	return tables
}
