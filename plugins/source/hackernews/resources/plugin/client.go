package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/resources/services/items"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hermanschaaf/hackernews"
	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 5
	defaultBackoff    = 10 * time.Second
	maxMsgSize        = 100 * 1024 * 1024 // 100 MiB
)

type Client struct {
	logger      zerolog.Logger
	tables      schema.Tables
	scheduler   *scheduler.Scheduler
	backendConn *grpc.ClientConn
	plugin.UnimplementedDestination
}

func (c *Client) GetSpec() any {
	return &client.Spec{}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
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
	return c.backendConn.Close()
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

func Configure(ctx context.Context, logger zerolog.Logger, specBytes []byte) (plugin.Client, error) {
	config := client.Spec{}
	if err := json.Unmarshal(specBytes, &config); err != nil {
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

	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		d := &net.Dialer{}
		return d.DialContext(ctx, "unix", addr)
	}
	conn, err := grpc.DialContext(ctx, config.Backend.Connection,
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxMsgSize),
			grpc.MaxCallSendMsgSize(maxMsgSize),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial grpc source plugin at %s: %w", config.Backend.Connection, err)
	}
	stateClient, err := state.NewClient(ctx, conn, config.Backend.Table)
	if err != nil {
		return nil, fmt.Errorf("failed to create state client: %w", err)
	}
	schedulerClient, err := client.New(logger, config, hnClient, stateClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create scheduler client: %w", err)
	}
	scheduler := scheduler.NewScheduler(schedulerClient,
		scheduler.WithLogger(logger),
	)
	return &Client{
		backendConn: conn,
		logger:      logger,
		scheduler:   scheduler,
		tables:      getTables(),
	}, nil
}
