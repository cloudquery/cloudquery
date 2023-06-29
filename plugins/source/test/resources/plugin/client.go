package plugin

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/test/client"
	"github.com/cloudquery/cloudquery/plugins/source/test/resources/services"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

var tables = []*schema.Table{
	services.TestSomeTable(),
	services.TestDataTable(),
}

type Client struct {
	SchedulerClient *client.TestClient
	logger          zerolog.Logger

	plugin.UnimplementedDestination
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	s := scheduler.NewScheduler(c.SchedulerClient, scheduler.WithStrategy(scheduler.StrategyDFS))
	return s.Sync(ctx, tables, res)
}

func (*Client) Tables(_ context.Context) (schema.Tables, error) {
	return tables, nil
}

func (*Client) Close(_ context.Context) error {
	return nil
}

func Configure(_ context.Context, logger zerolog.Logger, spec []byte) (plugin.Client, error) {
	schedulerClient := &client.TestClient{
		Logger: logger,
	}

	return &Client{
		logger:          logger,
		SchedulerClient: schedulerClient,
	}, nil
}
