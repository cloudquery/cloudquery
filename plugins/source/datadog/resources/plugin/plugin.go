package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

var (
	Version = "Development"

	customExceptions = map[string]string{
		"slo":  "SLO",
		"slos": "SLOs",
	}
)

type Client struct {
	plugin.UnimplementedDestination
	client    *client.Client
	options   plugin.NewClientOptions
	scheduler *scheduler.Scheduler
	allTables schema.Tables
}

func newClient(ctx context.Context, logger zerolog.Logger, specBytes []byte, options plugin.NewClientOptions) (plugin.Client, error) {
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
	clientMeta, err := client.Configure(ctx, logger, spec)
	if err != nil {
		return nil, err
	}
	c.client = clientMeta.(*client.Client)
	c.scheduler = scheduler.NewScheduler(scheduler.WithLogger(logger), scheduler.WithConcurrency(spec.Concurrency))
	return nil, nil
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
	if err := c.scheduler.Sync(ctx, c.client.Duplicate(), tables, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID)); err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}
	return nil
}

func Plugin() *plugin.Plugin {
	// here you can append custom non-generated tables
	return plugin.NewPlugin(
		"datadog",
		Version,
		newClient,
	)
}
