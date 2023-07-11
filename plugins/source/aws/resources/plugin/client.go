package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const maxMsgSize = 100 * 1024 * 1024 // 100 MiB

type Client struct {
	plugin.UnimplementedDestination
	scheduler *scheduler.Scheduler
	client    schema.ClientMeta
	logger    zerolog.Logger
	options   plugin.NewClientOptions
	allTables schema.Tables
}

func New(ctx context.Context, logger zerolog.Logger, specBytes []byte, options plugin.NewClientOptions) (plugin.Client, error) {
	var spec client.Spec
	c := &Client{
		options:   options,
		logger:    logger,
		allTables: getTables(),
	}
	if options.NoConnection {
		return c, nil
	}
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
		scheduler.WithConcurrency(spec.Concurrency),
		scheduler.WithLogger(logger),
	)
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
	tt, err := c.allTables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	var stateClient state.Client
	if options.BackendOptions == nil {
		c.logger.Info().Msg("No backend options provided, using no state backend")
		stateClient = &state.NoOpClient{}
	} else {
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
		stateClient, err = state.NewClient(ctx, conn, options.BackendOptions.TableName)
		if err != nil {
			return fmt.Errorf("failed to create state client: %w", err)
		}
		c.logger.Info().Str("table_name", options.BackendOptions.TableName).Msg("Connected to state backend")
	}
	awsClient := c.client.(*client.Client)
	// for each sync we want to create a copy of the client so they won't share state
	awsClient = awsClient.Duplicate()
	awsClient.Backend = stateClient
	return c.scheduler.Sync(ctx, awsClient, tt, res)
}
