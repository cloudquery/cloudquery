package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/groups"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/projects"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/settings"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/users"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
)

var Version = "Development"

type Client struct {
	plugin.UnimplementedDestination
	scheduler  *scheduler.Scheduler
	syncClient *client.Client
	options    plugin.NewClientOptions
	allTables  schema.Tables
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
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}
	syncClient, err := client.Configure(ctx, logger, *spec)
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
	tables, err := c.allTables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	if c.options.NoConnection {
		return fmt.Errorf("no connection")
	}
	tables, err := c.allTables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}
	return c.scheduler.Sync(ctx, c.syncClient, tables, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
}

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"gitlab",
		Version,
		newClient,
	)
}

func getTables() schema.Tables {
	tables := schema.Tables{
		groups.Groups(),
		projects.Projects(),
		settings.Settings(),
		users.Users(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	if err := transformers.Apply(tables, func(t *schema.Table) error {
		t.Title = docs.DefaultTitleTransformer(t)
		return nil
	}); err != nil {
		panic(err)
	}
	for _, table := range tables {
		schema.AddCqIDs(table)
	}
	return tables
}
