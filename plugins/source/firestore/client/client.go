package client

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"google.golang.org/api/option"
)

type Client struct {
	plugin.UnimplementedDestination
	logger         zerolog.Logger
	tables         schema.Tables
	options        plugin.NewClientOptions
	client         *firestore.Client
	maxBatchSize   int
	orderBy        string
	orderDirection string
}

var _ schema.ClientMeta = (*Client)(nil)

func (*Client) ID() string {
	return "source-firestore"
}

func Configure(ctx context.Context, logger zerolog.Logger, spec []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger:  logger,
			options: opts,
			tables:  schema.Tables{},
		}, nil
	}
	var firestoreSpec Spec
	err := json.Unmarshal(spec, &firestoreSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	if err := firestoreSpec.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	firestoreSpec.SetDefaults()

	var client *firestore.Client
	if firestoreSpec.ServiceAccountJSON == "" {
		// Use Application Default Credentials
		client, err = firestore.NewClient(ctx, firestoreSpec.ProjectID)
		if err != nil {
			return nil, fmt.Errorf("failed to create firestore client: %w", err)
		}
	} else {
		creds := option.WithCredentialsJSON([]byte(firestoreSpec.ServiceAccountJSON))
		client, err = firestore.NewClient(ctx, firestoreSpec.ProjectID, creds)
		if err != nil {
			return nil, fmt.Errorf("failed to create firestore client: %w", err)
		}
	}
	zctx := logger.With().Str("module", "firestore-source")
	c := &Client{
		logger:         zctx.Logger(),
		client:         client,
		maxBatchSize:   firestoreSpec.MaxBatchSize,
		orderBy:        firestoreSpec.OrderBy,
		orderDirection: firestoreSpec.OrderDirection,
		options:        opts,
	}

	c.tables, err = c.listTables(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}
	if len(c.tables) == 0 {
		return nil, fmt.Errorf("no tables found")
	}

	return c, nil
}

func (c Client) Tables(ctx context.Context, opts plugin.TableOptions) (schema.Tables, error) {
	if c.options.NoConnection {
		return schema.Tables{}, nil
	}
	return c.tables.FilterDfs(opts.Tables, opts.SkipTables, opts.SkipDependentTables)
}

func (c Client) Close(_ context.Context) error {
	return c.client.Close()
}
