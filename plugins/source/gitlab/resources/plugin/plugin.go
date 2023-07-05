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
	schduler   *scheduler.Scheduler
	syncClient *client.Client
	options    plugin.NewClientOptions
}

func newClient(ctx context.Context, logger zerolog.Logger, specBytes []byte, options plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		options: options,
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
	c.schduler = scheduler.NewScheduler(scheduler.WithLogger(logger), scheduler.WithConcurrency(uint64(spec.Concurrency)))
	return c, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}

func (*Client) Tables(ctx context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tables := getTables()
	tables, err := tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	if c.options.NoConnection {
		return fmt.Errorf("no connection")
	}
	tables := getTables()
	tables, err := tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}
	return c.schduler.Sync(ctx, c.syncClient, tables, res)
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
	for _, table := range tables {
		schema.AddCqIDs(table)
	}
	return tables
}
