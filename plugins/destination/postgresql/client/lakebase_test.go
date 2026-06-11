package client

import (
	"context"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestConfigureLakebase(t *testing.T) {
	cases := []struct {
		name            string
		connString      string
		wantMaxLifetime time.Duration
	}{
		{
			name:            "caps unset connection lifetime",
			connString:      "postgres://u@localhost:5432/db?sslmode=require",
			wantMaxLifetime: maxLakebaseConnLifetime,
		},
		{
			name:            "caps too-long connection lifetime",
			connString:      "postgres://u@localhost:5432/db?sslmode=require&pool_max_conn_lifetime=2h",
			wantMaxLifetime: maxLakebaseConnLifetime,
		},
		{
			name:            "preserves shorter connection lifetime",
			connString:      "postgres://u@localhost:5432/db?sslmode=require&pool_max_conn_lifetime=10m",
			wantMaxLifetime: 10 * time.Minute,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := pgxpool.ParseConfig(tc.connString)
			if err != nil {
				t.Fatalf("failed to parse connection string: %v", err)
			}

			err = configureLakebase(cfg, &spec.LakebaseSpec{
				Endpoint:     "projects/p/branches/b/endpoints/e",
				Host:         "https://example.cloud.databricks.com",
				ClientID:     "dummy-id",
				ClientSecret: "dummy-secret",
			})
			if err != nil {
				t.Fatalf("configureLakebase returned error: %v", err)
			}

			if cfg.BeforeConnect == nil {
				t.Error("expected BeforeConnect callback to be set")
			}
			if cfg.MaxConnLifetime != tc.wantMaxLifetime {
				t.Errorf("MaxConnLifetime = %v, want %v", cfg.MaxConnLifetime, tc.wantMaxLifetime)
			}
		})
	}
}

func TestConfigureLakebase_RequiresTLS(t *testing.T) {
	cases := []struct {
		name       string
		connString string
		wantErr    bool
	}{
		{name: "require ok", connString: "postgres://u@localhost:5432/db?sslmode=require"},
		{name: "verify-ca ok", connString: "postgres://u@localhost:5432/db?sslmode=verify-ca"},
		{name: "verify-full ok", connString: "postgres://u@localhost:5432/db?sslmode=verify-full"},
		{name: "disable rejected", connString: "postgres://u@localhost:5432/db?sslmode=disable", wantErr: true},
		// allow connects in plaintext first, then falls back to TLS.
		{name: "allow rejected", connString: "postgres://u@localhost:5432/db?sslmode=allow", wantErr: true},
		// prefer (the default when sslmode is unset) can fall back to plaintext.
		{name: "prefer rejected", connString: "postgres://u@localhost:5432/db?sslmode=prefer", wantErr: true},
		{name: "default (unset) rejected", connString: "postgres://u@localhost:5432/db", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := pgxpool.ParseConfig(tc.connString)
			if err != nil {
				t.Fatalf("failed to parse connection string: %v", err)
			}
			err = configureLakebase(cfg, &spec.LakebaseSpec{
				Endpoint:     "projects/p/branches/b/endpoints/e",
				Host:         "https://example.cloud.databricks.com",
				ClientID:     "dummy-id",
				ClientSecret: "dummy-secret",
			})
			if tc.wantErr {
				if err == nil {
					t.Error("expected an error for a non-TLS connection string, got nil")
				}
				if cfg.BeforeConnect != nil {
					t.Error("expected BeforeConnect to remain unset when TLS validation fails")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestConfigureLakebase_ChainsExistingBeforeConnect(t *testing.T) {
	cfg, err := pgxpool.ParseConfig("postgres://u@localhost:5432/db?sslmode=require")
	if err != nil {
		t.Fatalf("failed to parse connection string: %v", err)
	}

	// A pre-existing BeforeConnect hook (e.g. set by other configuration). It must
	// still be invoked after Lakebase auth is wired up.
	existingCalled := false
	cfg.BeforeConnect = func(_ context.Context, connConfig *pgx.ConnConfig) error {
		existingCalled = true
		connConfig.RuntimeParams["application_name"] = "existing-hook"
		return nil
	}

	if err := configureLakebase(cfg, &spec.LakebaseSpec{
		Endpoint:     "projects/p/branches/b/endpoints/e",
		Host:         "https://example.cloud.databricks.com",
		ClientID:     "dummy-id",
		ClientSecret: "dummy-secret",
	}); err != nil {
		t.Fatalf("configureLakebase returned error: %v", err)
	}

	// Invoke the composed hook. The Databricks credential call will fail offline,
	// but the existing hook must run (and persist its changes) before that.
	connConfig := cfg.ConnConfig.Copy()
	_ = cfg.BeforeConnect(context.Background(), connConfig)

	if !existingCalled {
		t.Error("expected the pre-existing BeforeConnect hook to be invoked")
	}
	if got := connConfig.RuntimeParams["application_name"]; got != "existing-hook" {
		t.Errorf("expected existing hook's changes to be preserved, application_name = %q", got)
	}
}
