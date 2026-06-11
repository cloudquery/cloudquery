package client

import (
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
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
