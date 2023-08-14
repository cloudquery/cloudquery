package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/resources/services/items"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/hermanschaaf/hackernews"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 5
	defaultBackoff    = 10 * time.Second
	maxMsgSize        = 100 * 1024 * 1024 // 100 MiB
)

type Client struct {
	logger      zerolog.Logger
	config      client.Spec
	tables      schema.Tables
	scheduler   *scheduler.Scheduler
	backendConn *grpc.ClientConn
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
	hnClient := hackernews.NewClient()
	if err != nil {
		return fmt.Errorf("failed to create hackernews client: %w", err)
	}

	var stateClient state.Client
	if options.BackendOptions == nil {
		c.logger.Info().Msg("No backend options provided, using no state backend")
		stateClient = &state.NoOpClient{}
		c.backendConn = nil
	} else {
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

	schedulerClient, err := client.New(c.logger, c.config, hnClient, stateClient)
	if err != nil {
		return fmt.Errorf("failed to create scheduler client: %w", err)
	}

	return c.scheduler.Sync(ctx, schedulerClient, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
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

func getTables() []*schema.Table {
	tables := []*schema.Table{
		items.Items(),
	}

	for i := range tables {
		tables[i].Columns = append([]schema.Column{schema.CqIDColumn, schema.CqParentIDColumn}, tables[i].Columns...)
		err := tables[i].Transform(tables[i])
		if err != nil {
			panic(err)
		}
	}
	return tables
}

func Configure(_ context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger: logger,
			tables: getTables(),
		}, nil
	}
	config := client.Spec{}
	if err := json.Unmarshal(specBytes, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	err := config.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	return &Client{
		config: config,
		logger: logger,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
		),
		tables: getTables(),
	}, nil
}
