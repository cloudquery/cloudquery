package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func ConnectionTester(ctx context.Context, _ zerolog.Logger, specBytes []byte) error {
	var s spec.Spec
	if err := json.Unmarshal(specBytes, &s); err != nil {
		return plugin.NewTestConnError("INVALID_SPEC", fmt.Errorf("failed to unmarshal spec: %w", err))
	}
	s.SetDefaults()
	if err := s.Validate(); err != nil {
		return plugin.NewTestConnError("INVALID_SPEC", fmt.Errorf("failed to validate spec: %w", err))
	}

	pgxConfig, err := pgxpool.ParseConfig(s.ConnectionString)
	if err != nil {
		return processError(err, "INVALID_CONFIG")
	}

	c, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return processError(err, "CONN_FAILED")
	}
	defer c.Close()

	_, err = currentDatabase(ctx, c)
	if err != nil {
		return processError(err, "DATABASE_FAILED")
	}

	_, err = currentSchema(ctx, c)
	if err != nil {
		return processError(err, "SCHEMA_FAILED")
	}

	return nil
}

func processError(err error, preferredErrorCode string) error {
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return plugin.NewTestConnError("DNS_FAILED", fmt.Errorf("no such host %q", dnsErr.Name))
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return plugin.NewTestConnError(preferredErrorCode, errors.New(pgErr.Message))
	}
	return plugin.NewTestConnError(preferredErrorCode, err)
}
