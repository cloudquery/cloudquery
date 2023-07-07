package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/access_groups"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/accounts"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/certificate_packs"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/dns_records"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/images"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/ips"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/waf_overrides"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/waf_packages"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/worker_meta_data"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/worker_routes"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/zones"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
)

var (
	Version = "development"
)

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
	syncClient, err := client.Configure(ctx, logger, spec)
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
	return c.scheduler.Sync(ctx, c.syncClient, tables, res)
}

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"cloudflare",
		Version,
		newClient,
	)
}

func getTables() schema.Tables {
	tables := []*schema.Table{
		access_groups.AccessGroups(),
		accounts.Accounts(),
		certificate_packs.CertificatePacks(),
		dns_records.DNSRecords(),
		images.Images(),
		ips.IPs(),
		waf_packages.WAFPackages(),
		waf_overrides.WAFOverrides(),
		worker_meta_data.WorkerMetaData(),
		worker_routes.WorkerRoutes(),
		zones.Zones(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	for _, table := range tables {
		schema.AddCqIDs(table)
	}
	return tables
}
