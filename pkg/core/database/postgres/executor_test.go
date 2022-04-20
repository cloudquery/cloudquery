package postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

type mockConn struct {
	row pgx.Row
}

func (m mockConn) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return m.row
}

type mockScanner struct {
	t   *testing.T
	val string
	err error
}

func (m mockScanner) Scan(dst ...interface{}) error {
	if len(dst) != 1 {
		m.t.Fatalf("called with %d args, want exactly one", len(dst))
	}
	ptr, ok := dst[0].(*string)
	if !ok {
		m.t.Fatalf("received %T, expected *string", dst[0])
	}
	*ptr = m.val
	return m.err
}

func Test_doValidatePostgresVersion(t *testing.T) {
	tests := []struct {
		name       string
		q          mockConn
		minVersion string
		wantErr    error
	}{
		{
			"scan error",
			mockConn{row: mockScanner{err: errors.New("scan")}},
			"10.0",
			errors.New("error getting PostgreSQL version: scan"),
		},
		{
			"strange version output",
			mockConn{row: mockScanner{t, "MSSQL", nil}},
			"10.0",
			errors.New("error getting PostgreSQL version: failed to parse version: MSSQL"),
		},
		{
			"unparsable version",
			mockConn{row: mockScanner{t, "PostgreSQL 10.a.1", nil}},
			"10.0",
			errors.New("error getting PostgreSQL version: Malformed version: 10.a.1"),
		},
		{
			"lower than needed",
			mockConn{row: mockScanner{t, "PostgreSQL 9.5 blah blah", nil}},
			"10.0",
			errors.New("unsupported PostgreSQL version: 9.5.0. (should be >= 10.0.0)"),
		},
		{
			"equal",
			mockConn{row: mockScanner{t, "PostgreSQL 10.0 blah blah", nil}},
			"10.0",
			nil,
		},
		{
			"greater than needed",
			mockConn{row: mockScanner{t, "PostgreSQL 12.5 blah blah", nil}},
			"10.0",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := version.Must(version.NewVersion(tt.minVersion))
			err := doValidatePostgresVersion(context.Background(), tt.q, want)
			if (tt.wantErr == nil) != (err == nil) {
				t.Errorf("wantErr is %v, returned error is %v", tt.wantErr, err)
			}
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}
