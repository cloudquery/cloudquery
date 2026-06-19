package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// maxLakebaseConnLifetime caps how long a pooled connection is reused before it
// is recycled. Lakebase OAuth tokens are valid for ~60 minutes, so connections
// are recycled well before expiry to ensure BeforeConnect mints a fresh token
// (Databricks recommends a maximum connection lifetime of 45 minutes).
const maxLakebaseConnLifetime = 45 * time.Minute

// configureLakebase wires up Databricks Lakebase authentication on the given pool
// config. It creates a Databricks workspace client and installs a BeforeConnect
// callback that generates a fresh short-lived OAuth database credential and uses
// it as the connection password for every new connection.
func configureLakebase(pgxConfig *pgxpool.Config, lb *spec.LakebaseSpec) error {
	// The Lakebase database credential is sensitive and Lakebase requires TLS, so
	// reject any connection string that could connect without it. Fail fast with a
	// clear error rather than leaking the token over plaintext or failing later in
	// a less obvious way.
	if !requiresTLS(pgxConfig.ConnConfig) {
		return errors.New("lakebase requires a TLS connection: set `sslmode=require` (or `verify-ca`/`verify-full`) in `connection_string`")
	}

	w, err := databricks.NewWorkspaceClient(&databricks.Config{
		Host:         lb.Host,
		ClientID:     lb.ClientID,
		ClientSecret: lb.ClientSecret,
	})
	if err != nil {
		return fmt.Errorf("failed to create databricks workspace client: %w", err)
	}

	// Preserve any previously configured BeforeConnect hook and run it first, so
	// Lakebase auth composes with other hooks instead of discarding them. The
	// Lakebase token is set last as it must be the connection password.
	prevBeforeConnect := pgxConfig.BeforeConnect
	pgxConfig.BeforeConnect = func(ctx context.Context, connConfig *pgx.ConnConfig) error {
		if prevBeforeConnect != nil {
			if err := prevBeforeConnect(ctx, connConfig); err != nil {
				return err
			}
		}
		cred, err := w.Postgres.GenerateDatabaseCredential(ctx, postgres.GenerateDatabaseCredentialRequest{
			Endpoint: lb.Endpoint,
		})
		if err != nil {
			return fmt.Errorf("failed to generate lakebase database credential: %w", err)
		}
		connConfig.Password = cred.Token
		return nil
	}

	// Recycle connections before the OAuth token expires so that the next
	// connection picks up a fresh credential via BeforeConnect.
	if pgxConfig.MaxConnLifetime <= 0 || pgxConfig.MaxConnLifetime > maxLakebaseConnLifetime {
		pgxConfig.MaxConnLifetime = maxLakebaseConnLifetime
	}

	return nil
}

// requiresTLS reports whether every connection path in the pgx config uses TLS.
// pgx represents sslmode=prefer (the default) and sslmode=allow via Fallbacks
// that can silently downgrade to a plaintext connection, so every path (the
// primary config plus all fallbacks) must have a non-nil TLSConfig. Only
// sslmode=require, verify-ca, and verify-full satisfy this.
func requiresTLS(connConfig *pgx.ConnConfig) bool {
	if connConfig.TLSConfig == nil {
		return false
	}
	for _, fb := range connConfig.Fallbacks {
		if fb.TLSConfig == nil {
			return false
		}
	}
	return true
}
