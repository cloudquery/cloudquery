package postgres

import (
	"context"
	"fmt"
	"strings"

	sdkpg "github.com/cloudquery/cq-provider-sdk/database/postgres"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Executor struct {
	logger hclog.Logger
	dsn    string
}

var MinPostgresVersion = version.Must(version.NewVersion("11.0"))

func New(logger hclog.Logger, dsn string) Executor {
	return Executor{
		logger: logger,
		dsn:    dsn,
	}
}

func (e Executor) Setup(ctx context.Context) (string, error) {
	return e.dsn, nil
}

func (e Executor) Validate(ctx context.Context) (bool, error) {
	pool, err := sdkpg.Connect(ctx, e.dsn)
	if err != nil {
		return false, err
	}

	if err := ValidatePostgresVersion(ctx, pool, MinPostgresVersion); err != nil {
		return true, err
	}

	return true, nil
}

func (e Executor) Finalize(ctx context.Context) error {
	return nil
}

// queryRower helps with unit tests
type queryRower interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func runningPostgresVersion(ctx context.Context, q queryRower) (*version.Version, error) {
	row := q.QueryRow(ctx, "SELECT version()")
	var result string
	if err := row.Scan(&result); err != nil {
		return nil, err
	}
	fields := strings.Fields(result)
	if len(fields) < 2 {
		return nil, fmt.Errorf("failed to parse version: %s", result)
	}
	return version.NewVersion(fields[1])
}

// ValidatePostgresVersion checks that PostgreSQL instance version available through pool is not lower than wanted version.
// In this case it returns nil. Otherwise returns error describing current and desired version or any other error encountered
// during the check.
func ValidatePostgresVersion(ctx context.Context, pool *pgxpool.Pool, want *version.Version) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	return doValidatePostgresVersion(ctx, conn, want)
}

func doValidatePostgresVersion(ctx context.Context, q queryRower, want *version.Version) error {
	got, err := runningPostgresVersion(ctx, q)
	if err != nil {
		return fmt.Errorf("error getting PostgreSQL version: %w", err)
	}
	if got.LessThan(want) {
		return fmt.Errorf("unsupported PostgreSQL version: %v. (should be >= %v)", got, want)
	}
	return nil
}
