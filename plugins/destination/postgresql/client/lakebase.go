package client

import (
	"context"
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
	w, err := databricks.NewWorkspaceClient(&databricks.Config{
		Host:         lb.Host,
		ClientID:     lb.ClientID,
		ClientSecret: lb.ClientSecret,
	})
	if err != nil {
		return fmt.Errorf("failed to create databricks workspace client: %w", err)
	}

	pgxConfig.BeforeConnect = func(ctx context.Context, connConfig *pgx.ConnConfig) error {
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
