package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/jira/resources/boards"
	"github.com/cloudquery/cloudquery/plugins/source/jira/resources/fields"
	"github.com/cloudquery/cloudquery/plugins/source/jira/resources/issues"
	"github.com/cloudquery/cloudquery/plugins/source/jira/resources/priorities"
	"github.com/cloudquery/cloudquery/plugins/source/jira/resources/projects"
	"github.com/cloudquery/cloudquery/plugins/source/jira/sync"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedDestination
	syncClient *sync.Client
	options    plugin.NewClientOptions
	scheduler  *scheduler.Scheduler
}

func New(ctx context.Context, logger zerolog.Logger, specBytes []byte, options plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		options: options,
	}
	if options.NoConnection {
		return c, nil
	}
	spec := &sync.Spec{}
	if err := json.Unmarshal(specBytes, spec); err != nil {
		return nil, err
	}
	syncClient, err := sync.New(ctx, logger, spec)
	if err != nil {
		return nil, err
	}
	c.syncClient = syncClient
	c.scheduler = scheduler.NewScheduler(scheduler.WithLogger(logger), scheduler.WithConcurrency(spec.Concurrency))
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
	if err := c.scheduler.Sync(ctx, c.syncClient, tables, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID)); err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}
	return nil
}

func getTables() schema.Tables {
	tables := schema.Tables{
		boards.Boards(),
		projects.Projects(),
		issues.Issues(),
		priorities.Priorities(),
		fields.Fields(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	return tables
}
