package timescale

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/client/database/postgres"
	"github.com/cloudquery/cloudquery/pkg/client/history"
	pgsdk "github.com/cloudquery/cq-provider-sdk/database/postgres"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/golang-migrate/migrate/v4"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	validateTimescaleInstalled = `SELECT EXISTS(SELECT 1 FROM pg_extension where extname = 'timescaledb')`
	timescaleVersionQuery      = `SELECT extversion FROM pg_catalog.pg_extension WHERE extname='timescaledb';`
)

var MinTimescaleVersion = version.Must(version.NewVersion("2.0"))

// queryRower helps with unit tests
type queryRower interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

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

	if err := postgres.ValidatePostgresConnection(ctx, pool); err != nil {
		return false, err
	}

	if err := postgres.ValidatePostgresVersion(ctx, pool); err != nil {
		return false, err
	}

	var installed bool
	if err := pgxscan.Get(ctx, pool, &installed, validateTimescaleInstalled); err != nil {
		return false, err
	}
	if !installed {
		return false, fmt.Errorf("timescaledb extension not installed, `CREATE EXTENSION IF NOT EXISTS timescaledb;`")
	}

	if err := ValidateTimescaleVersion(ctx, pool); err != nil {
		return false, err
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

func runningTimescaleVersion(ctx context.Context, q queryRower) (*version.Version, error) {
	row := q.QueryRow(ctx, timescaleVersionQuery)
	var result string
	if err := row.Scan(&result); err != nil {
		return nil, err
	}
	fields := strings.Fields(result)
	if len(fields) != 1 {
		return nil, fmt.Errorf("failed to parse version: %s, %q", result, fields)
	}
	return version.NewVersion(fields[0])
}

// ValidateTimescaleVersion checks that Timescale plugin version available through pool is not lower than wanted version.
// In this case it returns nil. Otherwise returns error describing current and desired version or any other error encountered
// during the check.
func ValidateTimescaleVersion(ctx context.Context, pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	return doValidateTimescaleVersion(ctx, conn, MinTimescaleVersion)
}

func doValidateTimescaleVersion(ctx context.Context, q queryRower, want *version.Version) error {
	got, err := runningTimescaleVersion(ctx, q)
	if err != nil {
		return fmt.Errorf("error getting Timescale version: %w", err)
	}
	if got.LessThan(want) {
		return fmt.Errorf("unsupported Timescale version: %s. (should be >= %s)", got.String(), want.String())
	}
	return nil
}
