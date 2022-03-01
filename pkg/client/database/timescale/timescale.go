package timescale

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/client/database/postgres"
	"github.com/cloudquery/cloudquery/pkg/client/history"
	pgsdk "github.com/cloudquery/cq-provider-sdk/database/postgres"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/golang-migrate/migrate/v4"
	"github.com/hashicorp/go-hclog"
)

const (
	validateTimescaleInstalled = `SELECT EXISTS(SELECT 1 FROM pg_extension where extname = 'timescaledb')`
)

type Executor struct {
	logger hclog.Logger
	dsn    string
	cfg    *history.Config
	ddl    *DDLManager
}

func New(logger hclog.Logger, dsn string, cfg *history.Config) (*Executor, error) {
	if cfg == nil {
		return nil, fmt.Errorf("missing history config")
	}
	return &Executor{
		logger: logger,
		dsn:    dsn,
		cfg:    cfg,
	}, nil
}

// Setup sets all required history functions and validation checks that it can run cleanly.
func (e *Executor) Setup(ctx context.Context) (string, error) {
	pool, err := pgsdk.Connect(ctx, e.dsn)
	if err != nil {
		return e.dsn, err
	}

	e.ddl, err = NewDDLManager(e.logger, pool, e.cfg, schema.TSDB)
	if err != nil {
		return e.dsn, err
	}

	if err := e.ddl.AddHistoryFunctions(ctx); err != nil {
		return e.dsn, fmt.Errorf("failed to create history functions: %w", err)
	}

	return history.TransformDSN(e.dsn)
}

func (e Executor) Validate(ctx context.Context) (bool, error) {
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

func (e Executor) Prepare(ctx context.Context) error {
	return e.ddl.DropViews(ctx)
}

func (e Executor) Finalize(ctx context.Context, retErr error) error {
	defer e.ddl.Close()

	if retErr != nil && retErr != migrate.ErrNoChange {
		return retErr
	}

	if err := e.ddl.SetupHistory(ctx); err != nil {
		return err
	}

	return retErr // keep migrate.ErrNoChange
}
