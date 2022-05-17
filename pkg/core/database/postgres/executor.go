package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	sdkpg "github.com/cloudquery/cq-provider-sdk/database/postgres"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-version"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Executor struct {
	dsn  string
	dbId string
}

var MinPostgresVersion = version.Must(version.NewVersion("11.0"))

func New(dsn string) *Executor {
	return &Executor{
		dsn: dsn,
	}
}

func (e *Executor) Validate(ctx context.Context) (bool, error) {
	pool, err := sdkpg.Connect(ctx, e.dsn)
	if err != nil {
		return false, err
	}

	if err := ValidatePostgresConnection(ctx, pool); err != nil {
		return false, err
	}

	// Not returning the error immediately as this error should not block anything
	var dbIdErr error
	e.dbId, dbIdErr = GetDatabaseId(ctx, pool)

	if err := ValidatePostgresVersion(ctx, pool); err != nil {
		return true, err
	}

	return true, dbIdErr
}

func (e Executor) Identifier(_ context.Context) (string, bool) {
	if e.dbId == "" {
		return "", false
	}
	return e.dbId, true
}

// ValidatePostgresConnection validates that we can actually connect to the postgres database.
func ValidatePostgresConnection(ctx context.Context, pool *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return pool.Ping(ctx)
}

// ValidatePostgresVersion checks that PostgreSQL instance version available through pool is not lower than wanted version.
// In this case it returns nil. Otherwise returns error describing current and desired version or any other error encountered
// during the check.
func ValidatePostgresVersion(ctx context.Context, pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	return doValidatePostgresVersion(ctx, conn, MinPostgresVersion)
}

func GetDatabaseId(ctx context.Context, q pgxscan.Querier) (string, error) {
	var result string
	err := pgxscan.Get(ctx, q, &result, `SELECT system_identifier::varchar AS id FROM pg_control_system()`)
	return result, err
}

func doValidatePostgresVersion(ctx context.Context, q pgxscan.Querier, want *version.Version) error {
	got, err := runningPostgresVersion(ctx, q)
	if err != nil {
		return fmt.Errorf("error getting PostgreSQL version: %w", err)
	}
	if got.LessThan(want) {
		return fmt.Errorf("unsupported PostgreSQL version: %s. (should be >= %s)", got.String(), want.String())
	}
	return nil
}

func runningPostgresVersion(ctx context.Context, q pgxscan.Querier) (*version.Version, error) {
	var result string
	if err := pgxscan.Get(ctx, q, &result, `SELECT version()`); err != nil {
		return nil, err
	}

	fields := strings.Fields(result)
	if len(fields) < 2 {
		return nil, fmt.Errorf("failed to parse version: %s", result)
	}
	return version.NewVersion(fields[1])
}
