package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"

	// Import the oracle driver
	_ "github.com/sijms/go-ora/v2"
)

type Client struct {
	plugin.UnimplementedDestination
	logger      zerolog.Logger
	tables      schema.Tables
	options     plugin.NewClientOptions
	db          *sql.DB
	concurrency int
}

var _ schema.ClientMeta = (*Client)(nil)

func (*Client) ID() string {
	return "oracledb"
}

func Configure(ctx context.Context, logger zerolog.Logger, spec []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger:  logger,
			options: opts,
			tables:  schema.Tables{},
		}, nil
	}
	var oracleDBSpec Spec
	err := json.Unmarshal(spec, &oracleDBSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	oracleDBSpec.SetDefaults()
	if err := oracleDBSpec.Validate(); err != nil {
		return nil, err
	}

	db, err := sql.Open("oracle", oracleDBSpec.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open oracle DB: %w", err)
	}

	// Sanity connection check
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get oracle DB connection: %w", err)
	}
	defer conn.Close()

	c := &Client{logger: logger.With().Str("module", "oracledb-source").Logger(), db: db, concurrency: oracleDBSpec.Concurrency, options: opts}
	c.tables, err = c.listTables(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
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
	return c.db.Close()
}
