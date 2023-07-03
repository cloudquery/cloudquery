package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rs/zerolog"
)

const maxMsgSize = 100 * 1024 * 1024 // 100 MiB

type Client struct {
	plugin.UnimplementedDestination
	schduler   *scheduler.Scheduler
	syncClient *client.Client
}

func NewClient(ctx context.Context, logger zerolog.Logger, specBytes []byte) (plugin.Client, error) {
	c := &Client{}
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
	c.schduler = scheduler.NewScheduler(scheduler.WithLogger(logger), scheduler.WithConcurrency(uint64(spec.Concurrency)))
	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	return nil
}

func (c *Client) Tables(ctx context.Context) (schema.Tables, error) {
	allTables := PluginAutoGeneratedTables()
	return allTables, nil
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	tables := PluginAutoGeneratedTables()
	tables, err := tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}
	syncClient := c.syncClient
	if options.BackendOptions != nil {
		conn, err := grpc.DialContext(ctx, options.BackendOptions.Connection,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(maxMsgSize),
				grpc.MaxCallSendMsgSize(maxMsgSize),
			),
		)
		if err != nil {
			return fmt.Errorf("failed to dial grpc source plugin at %s: %w", options.BackendOptions.Connection, err)
		}
		stateClient, err := state.NewClient(ctx, conn, options.BackendOptions.TableName)
		if err != nil {
			return fmt.Errorf("failed to create state client: %w", err)
		}
		syncClient = c.syncClient.WithBackend(stateClient)
	}
	return c.schduler.Sync(ctx, syncClient, tables, res)
}
