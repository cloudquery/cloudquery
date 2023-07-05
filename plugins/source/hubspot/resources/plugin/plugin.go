package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/resources/services/crm"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

var (
	Version = "development"
)

var customExceptions = map[string]string{
	"crm":     "CRM",
	"hubspot": "HubSpot",
}

func titleTransformer(table *schema.Table) string {
	if table.Title != "" {
		return table.Title
	}
	exceptions := make(map[string]string)
	for k, v := range docs.DefaultTitleExceptions {
		exceptions[k] = v
	}
	for k, v := range customExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	return csr.ToTitle(table.Name)
}

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
	syncClient, err := client.New(ctx, logger, *spec)
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

func getTables() schema.Tables {
	tables := schema.Tables{
		crm.Contacts(),
		crm.Companies(),
		crm.Deals(),
		crm.LineItems(),
		crm.Products(),
		crm.Tickets(),
		crm.Quotes(),
		crm.Owners(),
		crm.Pipelines(),
	}
	return tables
}

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"cloudquery-hubspot",
		Version,
		nil,
	)
}
