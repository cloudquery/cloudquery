package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client/spec"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/resources/services/crm"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
	"golang.org/x/exp/maps"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	maxMsgSize = 100 * 1024 * 1024 // 100 MiB
)

var customExceptions = map[string]string{
	"crm":     "CRM",
	"hubspot": "HubSpot",
}

type Client struct {
	logger    zerolog.Logger
	config    spec.Spec
	scheduler *scheduler.Scheduler
	options   plugin.NewClientOptions
	allTables schema.Tables
	plugin.UnimplementedDestination
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

	tables, err := c.allTables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	var stateClient state.Client
	if options.BackendOptions == nil {
		c.logger.Info().Msg("No backend options provided, using no state backend")
		stateClient = &state.NoOpClient{}
	} else {
		backendConn, err := grpc.DialContext(ctx, options.BackendOptions.Connection,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(maxMsgSize),
				grpc.MaxCallSendMsgSize(maxMsgSize),
			),
		)
		if err != nil {
			return fmt.Errorf("failed to dial grpc source plugin at %s: %w", options.BackendOptions.Connection, err)
		}
		defer backendConn.Close()
		stateClient, err = state.NewClient(ctx, backendConn, options.BackendOptions.TableName)
		if err != nil {
			return fmt.Errorf("failed to create state client: %w", err)
		}
		c.logger.Info().Str("table_name", options.BackendOptions.TableName).Msg("Connected to state backend")
	}

	schedulerClient, err := client.New(ctx, c.logger, c.config, stateClient)
	if err != nil {
		return err
	}

	syncErr := c.scheduler.Sync(ctx, schedulerClient, tables, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))

	if err = stateClient.Flush(ctx); err != nil {
		return errors.Join(fmt.Errorf("failed to save state: %w", err), syncErr)
	}

	return syncErr
}

func Configure(_ context.Context, logger zerolog.Logger, specBytes []byte, options plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger:    logger,
		options:   options,
		allTables: getTables(),
	}
	if options.NoConnection {
		return c, nil
	}
	s := spec.Spec{}
	if err := json.Unmarshal(specBytes, &s); err != nil {
		return nil, err
	}
	s.SetDefaults()
	if err := s.Validate(); err != nil {
		return nil, err
	}
	c.config = s
	c.scheduler = scheduler.NewScheduler(scheduler.WithLogger(logger), scheduler.WithConcurrency(s.Concurrency))
	return c, nil
}

func getTables() schema.Tables {
	tables := schema.Tables{
		crm.Contacts(),
		crm.Companies(),
		crm.Deals(),
		crm.LineItems(),
		crm.Products(),
		crm.Tickets(),
		crm.Quotes(),
		crm.Owners(),
		crm.Pipelines(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	if err := transformers.Apply(tables, titleTransformer()); err != nil {
		panic(err)
	}
	for _, table := range tables {
		schema.AddCqIDs(table)
	}
	return tables
}

func titleTransformer() func(table *schema.Table) error {
	exceptions := maps.Clone(docs.DefaultTitleExceptions)
	for k, v := range customExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	return func(table *schema.Table) error {
		if table.Title == "" {
			table.Title = csr.ToTitle(table.Name)
		}
		return nil
	}
}
