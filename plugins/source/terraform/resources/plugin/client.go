package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

type Client struct {
	logger    zerolog.Logger
	tables    schema.Tables
	scheduler *scheduler.Scheduler

	backends map[string]*client.TerraformBackend

	plugin.UnimplementedDestination
}

var _ plugin.Client = (*Client)(nil)

func (*Client) Close(context.Context) error { return nil }

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tt, nil
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	if options.BackendOptions != nil {
		c.logger.Warn().Msg("State backend not supported in plugin, skipping")
	}

	return c.scheduler.Sync(ctx, client.New(c.logger, c.backends), tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
}

func configure(ctx context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger: logger,
			tables: getTables(),
		}, nil
	}

	spec := new(client.Spec)
	if err := json.Unmarshal(specBytes, spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	var backends = make(map[string]*client.TerraformBackend, len(spec.Backends))
	for _, config := range spec.Backends {
		logger.Info().Msg("creating new backend")
		// create backend for each backend config
		b, err := client.NewBackend(ctx, &config)
		if err != nil {
			return nil, fmt.Errorf("cannot initialize backend: %w", err)
		}
		backends[b.BackendName] = b
	}

	// Returns the initialized client with requested backends
	return &Client{
		backends: backends,
		logger:   logger,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
			scheduler.WithConcurrency(spec.Concurrency),
		),
		tables: getTables(),
	}, nil
}
