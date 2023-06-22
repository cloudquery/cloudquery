package plugin

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedDestination
	scheduler *scheduler.Scheduler
	client    schema.ClientMeta
}

func New(ctx context.Context, logger zerolog.Logger, specBytes []byte) (plugin.Client, error) {
	var spec client.Spec
	c := &Client{}
	var err error
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		return nil, err
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}
	c.client, err = client.Configure(ctx, logger, spec)
	if err != nil {
		return nil, err
	}

	c.scheduler = scheduler.NewScheduler(
		c.client,
		scheduler.WithConcurrency(spec.Concurrency),
		scheduler.WithLogger(logger),
		scheduler.WithDeterministicCQId(true),
		scheduler.WithSchedulerStrategy(spec.Scheduler),
	)
	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	return nil
}

func (c *Client) Tables(ctx context.Context) (schema.Tables, error) {
	return tables(), nil
}
func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.Message) error {
	tables := tables()
	tt, err := tables.FilterDfs(options.Tables, options.SkipTables, false)
	if err != nil {
		return err
	}
	return c.scheduler.Sync(ctx, tt, res)
}
