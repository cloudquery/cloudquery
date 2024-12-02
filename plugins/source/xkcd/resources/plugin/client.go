package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"

	"github.com/cloudquery/cloudquery/plugins/source/xkcd/client"
	"github.com/cloudquery/cloudquery/plugins/source/xkcd/internal/xkcd"
	"github.com/cloudquery/cloudquery/plugins/source/xkcd/resources/services"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	maxMsgSize = 100 * 1024 * 1024 // 100 MiB
)

type Client struct {
	logger      zerolog.Logger
	config      client.Spec
	tables      schema.Tables
	options     plugin.NewClientOptions
	scheduler   *scheduler.Scheduler
	backendConn *grpc.ClientConn
	services    *xkcd.Client

	plugin.UnimplementedDestination
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	var stateClient state.Client
	if options.BackendOptions == nil {
		c.logger.Info().Msg("No backend options provided, using no state backend")
		stateClient = &state.NoOpClient{}
		c.backendConn = nil
	} else {
		// TODO: Remove once there's a documented migration path per https://github.com/grpc/grpc-go/issues/7244
		// nolint:staticcheck
		c.backendConn, err = grpc.DialContext(ctx, options.BackendOptions.Connection,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(maxMsgSize),
				grpc.MaxCallSendMsgSize(maxMsgSize),
			),
		)
		if err != nil {
			return fmt.Errorf("failed to dial grpc source plugin at %s: %w", options.BackendOptions.Connection, err)
		}
		stateClient, err = state.NewClient(ctx, c.backendConn, options.BackendOptions.TableName)
		if err != nil {
			return fmt.Errorf("failed to create state client: %w", err)
		}
		c.logger.Info().Str("table_name", options.BackendOptions.TableName).Msg("Connected to state backend")
	}

	schedulerClient := client.New(c.logger, c.config, c.services, stateClient)
	schedulerOptions := []scheduler.SyncOption{scheduler.WithSyncDeterministicCQID(options.DeterministicCQID)}
	if options.Shard != nil {
		schedulerOptions = append(schedulerOptions, scheduler.WithShard(options.Shard.Num, options.Shard.Total))
	}
	err = c.scheduler.Sync(ctx, schedulerClient, tt, res, schedulerOptions...)
	if err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}
	return stateClient.Flush(ctx)
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tt, nil
}

func (c *Client) Close(_ context.Context) error {
	if c.backendConn != nil {
		return c.backendConn.Close()
	}
	return nil
}

func Configure(_ context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger:  logger.With().Str("module", "xkcd").Logger(),
			options: opts,
			tables:  getTables(),
		}, nil
	}

	config := client.Spec{}
	if err := json.Unmarshal(specBytes, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	xkcdClient, err := xkcd.NewClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		logger:  logger.With().Str("module", "xkcd").Logger(),
		options: opts,
		config:  config,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
			scheduler.WithConcurrency(config.Concurrency),
			scheduler.WithInvocationID(opts.InvocationID),
		),
		services: xkcdClient,
		tables:   getTables(),
	}, nil
}

func getTables() schema.Tables {
	tables := []*schema.Table{
		services.ComicsTable(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	if err := transformers.Apply(tables, titleTransformer); err != nil {
		panic(err)
	}
	for _, t := range tables {
		schema.AddCqIDs(t)
	}
	return tables
}

func titleTransformer(table *schema.Table) error {
	if table.Title != "" {
		return nil
	}

	exceptions := maps.Clone(docs.DefaultTitleExceptions)
	exceptions["xkcd"] = "XKCD"
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	table.Title = csr.ToTitle(table.Name)
	return nil
}
