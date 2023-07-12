package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"

	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedDestination
	scheduler  *scheduler.Scheduler
	syncClient *client.Client
	options    plugin.NewClientOptions
	allTables  schema.Tables
}

func NewClient(ctx context.Context, logger zerolog.Logger, specBytes []byte, options plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		options:   options,
		allTables: getTables(),
	}
	if options.NoConnection {
		return c, nil
	}
	spec := &client.Spec{}
	if err := json.Unmarshal(specBytes, spec); err != nil {
		return nil, err
	}
	spec.SetDefaults()
	syncClient, err := client.New(ctx, logger, spec)
	if err != nil {
		return nil, err
	}
	c.syncClient = syncClient.(*client.Client)
	c.scheduler = scheduler.NewScheduler(scheduler.WithLogger(logger), scheduler.WithConcurrency(spec.Concurrency))
	return c, nil
}

func (*Client) Close(_ context.Context) error {
	return nil
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	return c.allTables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	if c.options.NoConnection {
		return fmt.Errorf("no connection")
	}
	tables, err := c.allTables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}
	return c.scheduler.Sync(ctx, c.syncClient.Duplicate(), tables, res)
}
