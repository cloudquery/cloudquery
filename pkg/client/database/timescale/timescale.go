package timescale

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/client/database/postgres"
	"github.com/cloudquery/cloudquery/pkg/client/history"
	pgsdk "github.com/cloudquery/cq-provider-sdk/database/postgres"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-hclog"
)

type Executor struct {
	logger hclog.Logger
	dsn    string
	cfg    *history.Config
}

func New(logger hclog.Logger, dsn string, cfg *history.Config) Executor {
	return Executor{
		logger: logger,
		dsn:    dsn,
		cfg:    cfg,
	}
}

// Setup sets all required history functions and validation checks that it can run cleanly.
func (e Executor) Setup(ctx context.Context) (string, error) {
	pool, err := pgsdk.Connect(ctx, e.dsn)
	if err != nil {
		return e.dsn, err
	}
	defer pool.Close()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return e.dsn, err
	}
	defer conn.Release()

	if err := history.AddHistoryFunctions(ctx, conn); err != nil {
		return e.dsn, fmt.Errorf("failed to create history functions: %w", err)
	}

	return history.TransformDSN(e.dsn)
}

func (e Executor) Validate(ctx context.Context) (bool, error) {
	const (
		validateTimescaleInstalled = `SELECT EXISTS(SELECT 1 FROM pg_extension where extname = 'timescaledb')`
	)

	pool, err := pgsdk.Connect(ctx, e.dsn)
	if err != nil {
		return false, err
	}
	defer pool.Close()

	if err := postgres.ValidatePostgresVersion(ctx, pool, postgres.MinPostgresVersion); err != nil {
		return false, err
	}

	var installed bool
	if err := pgxscan.Get(ctx, pool, &installed, validateTimescaleInstalled); err != nil {
		return false, err
	}
	if !installed {
		return false, fmt.Errorf("timescaledb extension not installed, `CREATE EXTENSION IF NOT EXISTS timescaledb;`")
	}

	return true, nil
}

func (e Executor) Finalize(ctx context.Context) error {
	pool, err := pgsdk.Connect(ctx, e.dsn)
	if err != nil {
		return err
	}
	defer pool.Close()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	ddl, err := history.NewDDLManager(e.logger, conn, e.cfg, schema.TSDB)
	if err != nil {
		return err
	}
	return ddl.SetupHistory(ctx, conn)
}
