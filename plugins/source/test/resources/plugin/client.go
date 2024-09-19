package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/test/client"
	"github.com/cloudquery/cloudquery/plugins/source/test/resources/services"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/premium"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
)

type Client struct {
	logger    zerolog.Logger
	options   plugin.NewClientOptions
	config    client.Spec
	tables    schema.Tables
	scheduler *scheduler.Scheduler

	plugin.UnimplementedDestination
	usage premium.UsageClient
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func hasPaidTables(tt schema.Tables) bool {
	flattenedTables := tt.FlattenTables()
	for _, t := range flattenedTables {
		if t.IsPaid {
			return true
		}
	}
	return false
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	if hasPaidTables(tt) {
		c.usage, err = premium.NewUsageClient(c.options.PluginMeta, premium.WithLogger(c.logger))
		if err != nil {
			return fmt.Errorf("failed to initialize usage client: %w", err)
		}
		ctx, err = premium.WithCancelOnQuotaExceeded(ctx, c.usage)
		if err != nil {
			return fmt.Errorf("failed to configure quota monitor: %w", err)
		}
	}

	schedulerClient := &client.Client{
		Logger: c.logger,
		Spec:   c.config,
	}

	schedulerOptions := []scheduler.SyncOption{scheduler.WithSyncDeterministicCQID(options.DeterministicCQID)}
	if options.Shard != nil {
		schedulerOptions = append(schedulerOptions, scheduler.WithShard(options.Shard.Num, options.Shard.Total))
	}

	return c.scheduler.Sync(ctx, schedulerClient, tt, res, schedulerOptions...)
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}

	return tt, nil
}

func (*Client) Close(_ context.Context) error {
	return nil
}

func Configure(_ context.Context, logger zerolog.Logger, spec []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		config := &client.Spec{}
		config.SetDefaults()

		return &Client{
			logger:  logger,
			options: opts,
			tables:  getTables(*config),
		}, nil
	}

	config := &client.Spec{}
	if err := json.Unmarshal(spec, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	for _, env := range config.RequiredEnv {
		parts := strings.Split(env, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid environment variable: %s", env)
		}
		key, value := parts[0], parts[1]
		if os.Getenv(key) != value {
			return nil, fmt.Errorf("environment variable did not match expectation: %s (want %q, but got %q)", key, value, os.Getenv(key))
		}
	}

	return &Client{
		logger:  logger,
		options: opts,
		config:  *config,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
			scheduler.WithInvocationID(opts.InvocationID),
		),
		tables: getTables(*config),
	}, nil
}

func getTables(config client.Spec) schema.Tables {
	tables := schema.Tables{
		services.TestSomeTable(config),
		services.TestDataTable(),
		services.TestPaidTable(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	for _, t := range tables {
		schema.AddCqIDs(t)
	}
	return tables
}

// OnBeforeSend increases the usage count for every message. If some messages should not be counted,
// they can be ignored here.
func (c *Client) OnBeforeSend(_ context.Context, msg message.SyncMessage) (message.SyncMessage, error) {
	if c.usage == nil {
		return msg, nil
	}

	si, ok := msg.(*message.SyncInsert)
	if !ok {
		return msg, nil
	}

	// now we need to determine whether the table used for sync was paid
	isPaid, ok := si.Record.Schema().Metadata().GetValue(schema.MetadataTableIsPaid)
	if !ok || isPaid != schema.MetadataTrue {
		return msg, nil
	}

	if err := c.usage.Increase(uint32(si.Record.NumRows())); err != nil {
		return msg, fmt.Errorf("failed to increase usage: %w", err)
	}

	return msg, nil
}

// OnSyncFinish is used to ensure the final usage count gets reported
func (c *Client) OnSyncFinish(_ context.Context) error {
	if c.usage != nil {
		return c.usage.Close()
	}
	return nil
}
