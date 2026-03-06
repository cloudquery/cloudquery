package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/tenable/client"
	"github.com/cloudquery/cloudquery/plugins/source/tenable/internal/tenable"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

type Client struct {
	logger    zerolog.Logger
	config    client.Spec
	tables    schema.Tables
	scheduler *scheduler.Scheduler
	services  *tenable.Client

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

	schedulerClient := client.New(c.logger, c.config, c.services, nil)
	schedulerOptions := []scheduler.SyncOption{scheduler.WithSyncDeterministicCQID(options.DeterministicCQID)}
	if options.Shard != nil {
		schedulerOptions = append(schedulerOptions, scheduler.WithShard(options.Shard.Num, options.Shard.Total))
	}

	if err := c.scheduler.Sync(ctx, schedulerClient, tt, res, schedulerOptions...); err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}
	return nil
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tt, nil
}

func (*Client) Close(_ context.Context) error { return nil }

func getTables() schema.Tables {
	return schema.Tables{}
}

func Configure(_ context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger: logger,
			tables: getTables(),
		}, nil
	}

	var spec client.Spec
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	tenableClient := tenable.NewClient(spec.AccessKey, spec.SecretKey)

	return &Client{
		logger:  logger,
		config:  spec,
		tables:  getTables(),
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
			scheduler.WithConcurrency(spec.Concurrency),
			scheduler.WithInvocationID(opts.InvocationID),
		),
		services: tenableClient,
	}, nil
}

func TestConnection(ctx context.Context, _ zerolog.Logger, specBytes []byte) error {
	var spec client.Spec
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		return plugin.NewTestConnError("INVALID_SPEC", fmt.Errorf("failed to unmarshal spec: %w", err))
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return plugin.NewTestConnError("INVALID_SPEC", fmt.Errorf("failed to validate spec: %w", err))
	}

	tenableClient := tenable.NewClient(spec.AccessKey, spec.SecretKey)
	if _, err := tenableClient.GetServerStatus(ctx); err != nil {
		if errors.Is(err, tenable.ErrUnauthorized) {
			return plugin.NewTestConnError("UNAUTHORIZED", fmt.Errorf("invalid API keys: %w", err))
		}
		return plugin.NewTestConnError("CONNECTION_FAILED", fmt.Errorf("failed to connect to Tenable.io: %w", err))
	}

	return nil
}
