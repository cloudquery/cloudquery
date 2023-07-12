package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/services/checkout"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/services/customer"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/services/order"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/services/price_rule"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/services/product"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rs/zerolog"
)

const (
	maxMsgSize = 100 * 1024 * 1024 // 100 MiB
)

type Client struct {
	logger      zerolog.Logger
	config      client.Spec
	tables      schema.Tables
	scheduler   *scheduler.Scheduler
	backendConn *grpc.ClientConn
	services    *shopify.Client

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
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	services, err := shopify.New(shopify.ClientOptions{
		Log: logger.With().Str("source", "shopify-client").Logger(),
		HC: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
		ApiKey:      config.APIKey,
		ApiSecret:   config.APISecret,
		AccessToken: config.AccessToken,
		ShopURL:     config.ShopURL,
		MaxRetries:  config.MaxRetries,
		PageSize:    int(config.PageSize),
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		config: config,
		logger: logger,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
			scheduler.WithConcurrency(config.Concurrency),
		),
		services: services,
		tables:   getTables(),
	}, nil
}

func getTables() schema.Tables {
	tables := []*schema.Table{
		customer.Customers(),
		order.Orders(),
		product.Products(),
		price_rule.PriceRules(),
		checkout.AbandonedCheckouts(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	for _, t := range tables {
		schema.AddCqIDs(t)
	}
	return tables
}
