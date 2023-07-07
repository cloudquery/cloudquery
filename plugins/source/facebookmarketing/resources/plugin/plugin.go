package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/resources/services"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
	"golang.org/x/exp/maps"
)

var (
	Version = "development"
)

var customExceptions = map[string]string{
	"facebookmarketing": "Facebook Marketing",
	"adaccounts":        "Ad Accounts",
	"adcreatives":       "Ad Creatives",
	"adimages":          "Ad Images",
	"adlabels":          "Ad Labels",
	"adplayables":       "Ad Playables",
	"adcloudplayables":  "Ad Cloud Playables",
	"advideos":          "Ad Videos",
	"adspixels":         "Ads Pixels",
	"adsets":            "Ad Sets",
	"customaudiences":   "Custom Audiences",
	"customconversions": "Custom Conversions",
}

func titleTransformer(table *schema.Table) {
	if table.Title == "" {
		exceptions := maps.Clone(docs.DefaultTitleExceptions)
		for k, v := range customExceptions {
			exceptions[k] = v
		}
		csr := caser.New(caser.WithCustomExceptions(exceptions))
		t := csr.ToTitle(table.Name)
		table.Title = strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
	}
	for _, rel := range table.Relations {
		titleTransformer(rel)
	}
}

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
	syncClient, err := client.New(ctx, logger, *spec)
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
	return c.scheduler.Sync(ctx, c.syncClient, tables, res)
}

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"facebookmarketing",
		Version,
		newClient,
	)
}

func getTables() schema.Tables {
	tables := schema.Tables{
		services.Campaigns(),
		services.Adsets(),
		services.Ads(),
		services.Adcreatives(),
		services.Adimages(),
		services.Advideos(),
		services.AdStudies(),
		services.Customaudiences(),
		services.Users(),
		services.Adaccounts(),
		services.AdPlacePageSets(),
		services.Adcloudplayables(),
		services.Adlabels(),
		services.Adplayables(),
		services.Adrules(),
		services.Adspixels(),
		services.AdvertisableApplications(),
		services.Businesses(),
		services.BroadTargetingCategoriess(),
		services.ConnectedInstagramAccounts(),
		services.Customconversions(),
		services.MaxBids(),
		services.OfflineConversionDataSets(),
		services.PromotePages(),
		services.PublisherBlockLists(),
		services.ReachFrequencyPredictions(),
		services.SavedAudiences(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	for _, table := range tables {
		titleTransformer(table)
		schema.AddCqIDs(table)
	}
	return tables
}
