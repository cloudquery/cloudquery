package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
	"github.com/{{.Org}}/cq-source-{{.Name}}/resources"
)

type Client struct {
	logger    zerolog.Logger
	config    Spec
	tables    schema.Tables
	scheduler *scheduler.Scheduler

	plugin.UnimplementedDestination
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "ID"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}


func New(_ context.Context, logger zerolog.Logger, spec []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger: logger,
			tables: getTables(),
		}, nil
	}

	config := &Spec{}
	if err := json.Unmarshal(spec, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	// TODO: Add your client initialization here

	return &Client{logger: logger, config: *config, scheduler: scheduler.NewScheduler(scheduler.WithLogger(logger)), tables: getTables()}, nil
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	return c.scheduler.Sync(ctx, &Client{}, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}

	return tt, nil
}

func (*Client) Close(_ context.Context) error {
	// TODO: Add your client cleanup here
	return nil
}

func getTables() schema.Tables {
	tables := schema.Tables{
		resources.SampleTable(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	for _, t := range tables {
		schema.AddCqIDs(t)
	}
	return tables
}