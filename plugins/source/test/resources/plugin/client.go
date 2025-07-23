package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/source/test/v4/client"
	"github.com/cloudquery/cloudquery/plugins/source/test/v4/resources/services"
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

	err = c.scheduler.Sync(ctx, schedulerClient, tt, res, schedulerOptions...)
	if err != nil {
		return fmt.Errorf("failed to run sync: %w", err)
	}

	if len(c.config.DeleteRecords) > 0 {
		c.deleteRecords(res)
	}
	return nil
}

func (c *Client) deleteRecords(res chan<- message.SyncMessage) {
	predicates := make([]message.Predicate, 0, len(c.config.DeleteRecords))
	for _, deleteID := range c.config.DeleteRecords {
		deleteRecord := array.NewRecordBuilder(memory.DefaultAllocator, (&schema.Table{
			Name: "test_some_table",
			Columns: schema.ColumnList{
				schema.Column{Name: "resource_id", Type: arrow.PrimitiveTypes.Int64},
			},
		}).ToArrowSchema())
		deleteRecord.Field(0).(*array.Int64Builder).Append(deleteID)
		deleteValue := deleteRecord.NewRecord()
		predicates = append(predicates, message.Predicate{
			Operator: "eq",
			Column:   "resource_id",
			Record:   deleteValue,
		})
	}
	res <- &message.SyncDeleteRecord{
		DeleteRecord: message.DeleteRecord{
			TableName: "test_some_table",
			WhereClause: message.PredicateGroups{
				{
					GroupingType: "OR",
					Predicates:   predicates,
				},
			},
		},
	}
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

func TestConnection(ctx context.Context, _ zerolog.Logger, specBytes []byte) error {
	var s client.Spec
	if err := json.Unmarshal(specBytes, &s); err != nil {
		return &plugin.TestConnError{
			Code:    "INVALID_SPEC",
			Message: fmt.Errorf("failed to unmarshal spec: %w", err),
		}
	}
	s.SetDefaults()

	return nil
}
