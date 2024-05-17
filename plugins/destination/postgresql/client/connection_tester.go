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
		return &plugin.TestConnError{
			Code:    plugin.TestConnFailureCodeInvalidSpec,
			Message: fmt.Errorf("failed to unmarshal spec: %w", err),
		}
	}
	s.SetDefaults()
	if err := s.Validate(); err != nil {
		return &plugin.TestConnError{
			Code:    plugin.TestConnFailureCodeInvalidSpec,
			Message: fmt.Errorf("failed to validate spec: %w", err),
		}
	}

	pgxConfig, err := pgxpool.ParseConfig(s.ConnectionString)
	if err != nil {
		return &plugin.TestConnError{
			Code:    plugin.TestConnFailureCodeInvalidCredentials,
			Message: processError(err),
		}
	}

	c, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		return &plugin.TestConnError{
			Code:    plugin.TestConnFailureCodeInvalidCredentials,
			Message: processError(err),
		}
	}
	defer c.Close()

	_, err = currentDatabase(ctx, c)
	if err != nil {
		return &plugin.TestConnError{
			Code:    plugin.TestConnFailureCodeInvalidCredentials,
			Message: processError(err),
		}
	}

	_, err = currentSchema(ctx, c)
	if err != nil {
		return &plugin.TestConnError{
			Code:    plugin.TestConnFailureCodeInvalidCredentials,
			Message: processError(err),
		}
	}

	return nil
}

func processError(err error) error {
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return fmt.Errorf("no such host %q", dnsErr.Name)
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return errors.New(pgErr.Message)
	}
	return err
}
