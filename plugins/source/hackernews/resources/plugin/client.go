package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/resources/services/items"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"

	"github.com/hermanschaaf/hackernews"
	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 5
	defaultBackoff    = 10 * time.Second
)

type Client struct {
	logger    zerolog.Logger
	tables    schema.Tables
	scheduler *scheduler.Scheduler
	plugin.UnimplementedDestination
}

func (c *Client) GetSpec() any {
	return &client.Spec{}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.Message) error {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}
	return c.scheduler.Sync(ctx, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
}

func (c *Client) Tables(ctx context.Context) (schema.Tables, error) {
	return c.tables, nil
}

func (c *Client) Close(ctx context.Context) error {
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

func Configure(ctx context.Context, logger zerolog.Logger, spec []byte) (plugin.Client, error) {
	config := &client.Spec{}
	if err := json.Unmarshal(spec, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	err := config.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	hnClient := hackernews.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create homebrew client: %w", err)
	}

	backendConfig := managedplugin.Config{
		Name:     config.Backend.Name,
		Registry: config.Backend.Registry,
		Path:     config.Backend.Path,
		Version:  config.Backend.Version,
	}
	backendPlugin, err := managedplugin.NewClient(ctx, managedplugin.PluginDestination, backendConfig,
		managedplugin.WithLogger(logger),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create managed plugin client for backend: %w", err)
	}
	backendSpecBytes, err := json.Marshal(config.Backend.Spec)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal backend spec: %w", err)
	}
	stateClient, err := state.NewClient(ctx, backendPlugin.Conn, backendSpecBytes, config.Backend.Table)
	if err != nil {
		return nil, fmt.Errorf("failed to create state client: %w", err)
	}

	schedulerClient, err := client.New(logger, *config, hnClient, stateClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create scheduler client: %w", err)
	}
	scheduler := scheduler.NewScheduler(schedulerClient,
		scheduler.WithLogger(logger),
	)
	return &Client{
		logger:    logger,
		scheduler: scheduler,
		tables:    getTables(),
	}, nil
}
