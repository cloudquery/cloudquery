package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudquery/cloudquery/plugins/source/googleanalytics/client"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	maxMsgSize = 100 * 1024 * 1024 // 100 MiB
)

type Client struct {
	plugin.UnimplementedDestination
	scheduler *scheduler.Scheduler
	options   plugin.NewClientOptions
	spec      client.Spec

	logger zerolog.Logger
}

func (*Client) Close(_ context.Context) error {
	return nil
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	if c.options.NoConnection {
		return schema.Tables{}, nil
	}
	tables := make(schema.Tables, len(c.spec.Reports))
	for i, r := range c.spec.Reports {
		tables[i] = r.Table(c.spec.PropertyID)
	}
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

	tables, err := c.Tables(ctx, plugin.TableOptions{
		Tables:              options.Tables,
		SkipTables:          options.SkipTables,
		SkipDependentTables: options.SkipDependentTables,
	})
	if err != nil {
		return err
	}
	syncClient, err := client.New(ctx, c.logger, c.spec, stateClient)
	if err != nil {
		return err
	}
	err = c.scheduler.Sync(ctx, syncClient, tables, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
	if err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}
	return stateClient.Flush(ctx)
}

func Configure(ctx context.Context, logger zerolog.Logger, specBytes []byte, options plugin.NewClientOptions) (plugin.Client, error) {
	if options.NoConnection {
		return &Client{
			logger:  logger,
			options: options,
		}, nil
	}
	spec := new(client.Spec)
	if err := json.Unmarshal(specBytes, spec); err != nil {
		return nil, err
	}

	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	c := &Client{
		spec:    *spec,
		options: options,
		logger: logger.With().
			Str("plugin", "googleanalytics").
			Str("property_id", spec.PropertyID).
			Logger(),
	}
	c.scheduler = scheduler.NewScheduler(scheduler.WithConcurrency(spec.Concurrency))
	return c, nil
}
