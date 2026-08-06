package client

import (
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestCheckTLS(t *testing.T) {
	cases := []struct {
		connString    string
		wantEncrypted bool
		wantVerified  bool
	}{
		{connString: "postgres://u@localhost:5432/db?sslmode=verify-full", wantEncrypted: true, wantVerified: true},
		// verify-ca skips hostname verification.
		{connString: "postgres://u@localhost:5432/db?sslmode=verify-ca", wantEncrypted: true},
		{connString: "postgres://u@localhost:5432/db?sslmode=require", wantEncrypted: true},
		{connString: "postgres://u@localhost:5432/db?sslmode=prefer"},
		{connString: "postgres://u@localhost:5432/db?sslmode=allow"},
		{connString: "postgres://u@localhost:5432/db?sslmode=disable"},
		{connString: "postgres://u@localhost:5432/db"},
	}

	for _, tc := range cases {
		t.Run(tc.connString, func(t *testing.T) {
			connConfig, err := pgx.ParseConfig(tc.connString)
			if err != nil {
				t.Fatalf("failed to parse connection string: %v", err)
			}
			if got := checkTLS(connConfig, tlsEncrypted); got != tc.wantEncrypted {
				t.Errorf("checkTLS(tlsEncrypted) = %v, want %v", got, tc.wantEncrypted)
			}
			if got := checkTLS(connConfig, tlsVerified); got != tc.wantVerified {
				t.Errorf("checkTLS(tlsVerified) = %v, want %v", got, tc.wantVerified)
			}
		})
	}
}
