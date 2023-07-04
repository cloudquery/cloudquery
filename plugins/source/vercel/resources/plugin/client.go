package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/deployment"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/domain"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/project"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/team"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
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
	services    *vercel.Client

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

	schedulerClient := client.New(c.logger, c.config, c.services, c.config.TeamIDs, stateClient)
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
	return c.backendConn.Close()
}

func Configure(ctx context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
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

	services, err := getServiceClient(logger.With().Str("source", "vercel-client").Logger(), config, "")
	if err != nil {
		return nil, err
	}
	if len(config.TeamIDs) == 0 {
		config.TeamIDs, err = getTeamIDs(ctx, services)
		if err != nil {
			return nil, fmt.Errorf("failed to discover team ids: %w", err)
		}
	}

	return &Client{
		config: config,
		logger: logger,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
		),
		services: services,
		tables:   getTables(),
	}, nil
}

func getTables() schema.Tables {
	tables := []*schema.Table{
		domain.Domains(),
		team.Teams(),
		project.Projects(),
		deployment.Deployments(),
	}
	for _, t := range tables {
		if err := t.Transform(t); err != nil {
			panic(err)
		}
		for _, rel := range t.Relations {
			if err := rel.Transform(rel); err != nil {
				panic(err)
			}
		}

		schema.AddCqIDs(t)
	}
	return tables
}

func getServiceClient(logger zerolog.Logger, spec client.Spec, teamID string) (*vercel.Client, error) {
	return vercel.New(
		logger,
		&http.Client{
			Timeout: time.Duration(spec.Timeout) * time.Second,
		},
		spec.EndpointURL,
		spec.AccessToken,
		teamID,
		spec.MaxRetries,
		spec.MaxWait,
		spec.PageSize,
	), nil
}

func getTeamIDs(ctx context.Context, svc *vercel.Client) ([]string, error) {
	var pg vercel.Paginator
	var teams []string

	for {
		list, p, err := svc.ListTeams(ctx, &pg)
		if err != nil {
			return nil, err
		}
		for _, t := range list {
			teams = append(teams, t.ID)
		}

		if p.Next == nil {
			break
		}
		pg.Next = p.Next
	}

	return teams, nil
}
