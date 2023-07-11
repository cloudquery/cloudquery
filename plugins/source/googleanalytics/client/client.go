package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
	"google.golang.org/api/option"
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
	service   *analyticsdata.Service
	backend   state.Client

	reports []*Report

	PropertyID string
	StartDate  string

	logger zerolog.Logger
}

var _ schema.ClientMeta = (*Client)(nil)

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return "googleanalytics:property-id:{" + c.PropertyID + "}"
}

func (*Client) Close(ctx context.Context) error {
	return nil
}

func (c *Client) Tables(ctx context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tables := make(schema.Tables, len(c.reports))
	for i, r := range c.reports {
		tables[i] = r.table(c.PropertyID)
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
	c.backend = stateClient

	tables, err := c.Tables(ctx, plugin.TableOptions{
		Tables:              options.Tables,
		SkipTables:          options.SkipTables,
		SkipDependentTables: options.SkipDependentTables,
	})
	if err != nil {
		return err
	}
	if err := c.scheduler.Sync(ctx, c, tables, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID)); err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}
	return nil
}

func Configure(ctx context.Context, logger zerolog.Logger, specBytes []byte, options plugin.NewClientOptions) (plugin.Client, error) {
	spec := new(Spec)
	if err := json.Unmarshal(specBytes, spec); err != nil {
		return nil, err
	}

	spec.setDefaults()
	if err := spec.validate(); err != nil {
		return nil, err
	}

	opts := []option.ClientOption{
		option.WithScopes(analyticsdata.AnalyticsReadonlyScope),
		option.WithRequestReason("cloudquery resource fetch"),
		// we disable telemetry to boost performance and be on the safe side with telemetry
		option.WithTelemetryDisabled(),
	}

	if spec.OAuth != nil {
		tokenSource, err := spec.OAuth.getTokenSource(ctx)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(tokenSource))
	}

	svc, err := analyticsdata.NewService(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	svc.UserAgent = "cloudquery:source-googleanalytics"

	c := &Client{
		service:    svc,
		StartDate:  spec.StartDate,
		PropertyID: spec.PropertyID,
		reports:    spec.Reports,
		logger: logger.With().
			Str("plugin", "googleanalytics").
			Str("property_id", spec.PropertyID).
			Logger(),
	}
	c.scheduler = scheduler.NewScheduler(scheduler.WithConcurrency(spec.Concurrency))
	return c, nil
}
