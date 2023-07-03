package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/rs/zerolog"

	// Import the oracle driver
	_ "github.com/sijms/go-ora/v2"
)

type Client struct {
	logger      zerolog.Logger
	metrics     *source.Metrics
	Tables      schema.Tables
	db          *sql.DB
	Concurrency uint64
}

var _ schema.ClientMeta = (*Client)(nil)

func (*Client) ID() string {
	return "oracledb"
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	var oracleDBSpec Spec
	err := spec.UnmarshalSpec(&oracleDBSpec)
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

	c := &Client{logger: logger.With().Str("module", "oracledb-source").Logger(), db: db, Concurrency: spec.Concurrency}
	c.Tables, err = c.listTables(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}
	c.Tables, err = c.Tables.FilterDfs(spec.Tables, spec.SkipTables, spec.SkipDependentTables)
	if err != nil {
		return nil, fmt.Errorf("failed to apply config to tables: %w", err)
	}

	return c, nil
}
