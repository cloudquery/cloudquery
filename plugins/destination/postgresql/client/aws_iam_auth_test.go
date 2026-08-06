package client

import (
	"context"
	"net/url"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Signing is local, so pinning credentials and pointing the AWS config files at a
// missing path keeps these tests hermetic.
func setStaticAWSCredentials(t *testing.T) {
	t.Helper()
	missing := filepath.Join(t.TempDir(), "does-not-exist")
	t.Setenv("AWS_CONFIG_FILE", missing)
	t.Setenv("AWS_SHARED_CREDENTIALS_FILE", missing)
	t.Setenv("AWS_PROFILE", "")
	t.Setenv("AWS_ACCESS_KEY_ID", "AKIAIOSFODNN7EXAMPLE")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")
	t.Setenv("AWS_SESSION_TOKEN", "")
	t.Setenv("AWS_REGION", "")
}

func TestConfigureAWSIAMAuth_SignsRDSToken(t *testing.T) {
	setStaticAWSCredentials(t)

	cases := []struct {
		name         string
		connString   string
		specEndpoint string
		wantHostPort string
		wantUser     string
	}{
		{
			name:         "signs the connection string host and port",
			connString:   "postgres://cq_user@mydb.123456789012.us-east-1.rds.amazonaws.com:5432/db?sslmode=require",
			wantHostPort: "mydb.123456789012.us-east-1.rds.amazonaws.com:5432",
			wantUser:     "cq_user",
		},
		{
			name:         "signs a non-default port",
			connString:   "postgres://cq_user@mydb.123456789012.us-east-1.rds.amazonaws.com:5433/db?sslmode=require",
			wantHostPort: "mydb.123456789012.us-east-1.rds.amazonaws.com:5433",
			wantUser:     "cq_user",
		},
		{
			name:         "endpoint override replaces host and port",
			connString:   "postgres://cq_user@localhost:15432/db?sslmode=require",
			specEndpoint: "mydb.123456789012.us-east-1.rds.amazonaws.com:5432",
			wantHostPort: "mydb.123456789012.us-east-1.rds.amazonaws.com:5432",
			wantUser:     "cq_user",
		},
		{
			name:         "endpoint override without a port keeps the connection string port",
			connString:   "postgres://cq_user@localhost:5433/db?sslmode=require",
			specEndpoint: "mydb.123456789012.us-east-1.rds.amazonaws.com",
			wantHostPort: "mydb.123456789012.us-east-1.rds.amazonaws.com:5433",
			wantUser:     "cq_user",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := pgxpool.ParseConfig(tc.connString)
			if err != nil {
				t.Fatalf("failed to parse connection string: %v", err)
			}

			err = configureAWSIAMAuth(context.Background(), cfg, &spec.AWSIAMAuthSpec{
				Region:   "us-east-1",
				Endpoint: tc.specEndpoint,
			})
			if err != nil {
				t.Fatalf("configureAWSIAMAuth returned error: %v", err)
			}
			if cfg.BeforeConnect == nil {
				t.Fatal("expected BeforeConnect callback to be set")
			}

			connConfig := cfg.ConnConfig.Copy()
			if err := cfg.BeforeConnect(context.Background(), connConfig); err != nil {
				t.Fatalf("BeforeConnect returned error: %v", err)
			}

			// `host:port/?Action=connect&DBUser=...&X-Amz-Signature=...`
			token := connConfig.Password
			hostPort, rawQuery, found := strings.Cut(token, "?")
			if !found {
				t.Fatalf("token is not a signed request: %q", token)
			}
			if hostPort = strings.TrimSuffix(hostPort, "/"); hostPort != tc.wantHostPort {
				t.Errorf("token signed for %q, want %q", hostPort, tc.wantHostPort)
			}
			query, err := url.ParseQuery(rawQuery)
			if err != nil {
				t.Fatalf("failed to parse the token query: %v", err)
			}
			if got := query.Get("DBUser"); got != tc.wantUser {
				t.Errorf("DBUser = %q, want %q", got, tc.wantUser)
			}
			if got := query.Get("Action"); got != "connect" {
				t.Errorf("Action = %q, want %q", got, "connect")
			}
			if query.Get("X-Amz-Signature") == "" {
				t.Error("expected the token to be signed, X-Amz-Signature is empty")
			}
			if creds := query.Get("X-Amz-Credential"); !strings.Contains(creds, "us-east-1/rds-db") {
				t.Errorf("X-Amz-Credential = %q, want it to be scoped to us-east-1/rds-db", creds)
			}
		})
	}
}

func TestConfigureAWSIAMAuth_SignsFreshTokenPerConnection(t *testing.T) {
	setStaticAWSCredentials(t)

	cfg, err := pgxpool.ParseConfig("postgres://cq_user@mydb.123456789012.us-east-1.rds.amazonaws.com:5432/db?sslmode=require")
	if err != nil {
		t.Fatalf("failed to parse connection string: %v", err)
	}
	if err := configureAWSIAMAuth(context.Background(), cfg, &spec.AWSIAMAuthSpec{Region: "us-east-1"}); err != nil {
		t.Fatalf("configureAWSIAMAuth returned error: %v", err)
	}

	first := cfg.ConnConfig.Copy()
	if err := cfg.BeforeConnect(context.Background(), first); err != nil {
		t.Fatalf("BeforeConnect returned error: %v", err)
	}
	second := cfg.ConnConfig.Copy()
	if err := cfg.BeforeConnect(context.Background(), second); err != nil {
		t.Fatalf("BeforeConnect returned error: %v", err)
	}

	if first.Password == "" || second.Password == "" {
		t.Fatal("expected BeforeConnect to set a token as the connection password")
	}
	if cfg.ConnConfig.Password != "" {
		t.Error("expected the pool's base config password to be left untouched")
	}
}

// Sessions outlive their token, so the pool must not get Lakebase's lifetime cap.
func TestConfigureAWSIAMAuth_KeepsConnectionLifetime(t *testing.T) {
	setStaticAWSCredentials(t)

	cfg, err := pgxpool.ParseConfig("postgres://cq_user@mydb.123456789012.us-east-1.rds.amazonaws.com:5432/db?sslmode=require")
	if err != nil {
		t.Fatalf("failed to parse connection string: %v", err)
	}
	want := cfg.MaxConnLifetime

	if err := configureAWSIAMAuth(context.Background(), cfg, &spec.AWSIAMAuthSpec{Region: "us-east-1"}); err != nil {
		t.Fatalf("configureAWSIAMAuth returned error: %v", err)
	}
	if cfg.MaxConnLifetime != want {
		t.Errorf("MaxConnLifetime = %v, want %v", cfg.MaxConnLifetime, want)
	}
}

func TestConfigureAWSIAMAuth_RequiresTLS(t *testing.T) {
	setStaticAWSCredentials(t)

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
			err = configureAWSIAMAuth(context.Background(), cfg, &spec.AWSIAMAuthSpec{Region: "us-east-1"})
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

func TestConfigureAWSIAMAuth_UnsupportedService(t *testing.T) {
	setStaticAWSCredentials(t)

	cfg, err := pgxpool.ParseConfig("postgres://u@localhost:5432/db?sslmode=require")
	if err != nil {
		t.Fatalf("failed to parse connection string: %v", err)
	}

	err = configureAWSIAMAuth(context.Background(), cfg, &spec.AWSIAMAuthSpec{Service: "dsql", Region: "us-east-1"})
	if err == nil {
		t.Fatal("expected an error for an unsupported service, got nil")
	}
	if !strings.Contains(err.Error(), "dsql") {
		t.Errorf("error should name the unsupported service, got: %v", err)
	}
}

func TestConfigureAWSIAMAuth_InvalidEndpoint(t *testing.T) {
	setStaticAWSCredentials(t)

	cases := []string{
		"mydb.123456789012.us-east-1.rds.amazonaws.com:not-a-port",
		"mydb.123456789012.us-east-1.rds.amazonaws.com:5432:5432",
		"mydb.123456789012.us-east-1.rds.amazonaws.com:99999",
	}

	for _, endpoint := range cases {
		t.Run(endpoint, func(t *testing.T) {
			cfg, err := pgxpool.ParseConfig("postgres://u@localhost:5432/db?sslmode=require")
			if err != nil {
				t.Fatalf("failed to parse connection string: %v", err)
			}
			err = configureAWSIAMAuth(context.Background(), cfg, &spec.AWSIAMAuthSpec{
				Region:   "us-east-1",
				Endpoint: endpoint,
			})
			if err == nil {
				t.Error("expected an error for a malformed endpoint, got nil")
			}
		})
	}
}

func TestConfigureAWSIAMAuth_UnresolvedRegion(t *testing.T) {
	setStaticAWSCredentials(t)

	cfg, err := pgxpool.ParseConfig("postgres://u@localhost:5432/db?sslmode=require")
	if err != nil {
		t.Fatalf("failed to parse connection string: %v", err)
	}

	err = configureAWSIAMAuth(context.Background(), cfg, &spec.AWSIAMAuthSpec{})
	if err == nil {
		t.Fatal("expected an error when the aws region cannot be resolved, got nil")
	}
	if !strings.Contains(err.Error(), "region") {
		t.Errorf("error should mention the unresolved region, got: %v", err)
	}
}

func TestConfigureAWSIAMAuth_ChainsExistingBeforeConnect(t *testing.T) {
	setStaticAWSCredentials(t)

	cfg, err := pgxpool.ParseConfig("postgres://cq_user@mydb.123456789012.us-east-1.rds.amazonaws.com:5432/db?sslmode=require")
	if err != nil {
		t.Fatalf("failed to parse connection string: %v", err)
	}

	existingCalled := false
	cfg.BeforeConnect = func(_ context.Context, connConfig *pgx.ConnConfig) error {
		existingCalled = true
		connConfig.RuntimeParams["application_name"] = "existing-hook"
		return nil
	}

	if err := configureAWSIAMAuth(context.Background(), cfg, &spec.AWSIAMAuthSpec{Region: "us-east-1"}); err != nil {
		t.Fatalf("configureAWSIAMAuth returned error: %v", err)
	}

	connConfig := cfg.ConnConfig.Copy()
	if err := cfg.BeforeConnect(context.Background(), connConfig); err != nil {
		t.Fatalf("BeforeConnect returned error: %v", err)
	}

	if !existingCalled {
		t.Error("expected the pre-existing BeforeConnect hook to be invoked")
	}
	if got := connConfig.RuntimeParams["application_name"]; got != "existing-hook" {
		t.Errorf("expected existing hook's changes to be preserved, application_name = %q", got)
	}
	if connConfig.Password == "" {
		t.Error("expected the token to be set as the connection password")
	}
}

func TestParseAWSIAMAuthEndpoint(t *testing.T) {
	cases := []struct {
		endpoint string
		wantHost string
		wantPort uint16
		wantErr  bool
	}{
		{endpoint: "", wantHost: "", wantPort: 0},
		{endpoint: "mydb.us-east-1.rds.amazonaws.com", wantHost: "mydb.us-east-1.rds.amazonaws.com"},
		{endpoint: "mydb.us-east-1.rds.amazonaws.com:5432", wantHost: "mydb.us-east-1.rds.amazonaws.com", wantPort: 5432},
		{endpoint: "abc.dsql.us-east-1.on.aws", wantHost: "abc.dsql.us-east-1.on.aws"},
		{endpoint: "[::1]:5432", wantHost: "::1", wantPort: 5432},
		{endpoint: "mydb:0", wantErr: true},
		{endpoint: "mydb:-1", wantErr: true},
		{endpoint: "mydb:65536", wantErr: true},
		{endpoint: "mydb:5432:5432", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.endpoint, func(t *testing.T) {
			host, port, err := parseAWSIAMAuthEndpoint(tc.endpoint)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected an error for %q, got host %q port %d", tc.endpoint, host, port)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if host != tc.wantHost {
				t.Errorf("host = %q, want %q", host, tc.wantHost)
			}
			if port != tc.wantPort {
				t.Errorf("port = %d, want %d", port, tc.wantPort)
			}
		})
	}
}
