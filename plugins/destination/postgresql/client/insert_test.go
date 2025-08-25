package client

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MockBatchResults struct {
	closeErr error
}

func (m *MockBatchResults) Close() error { return m.closeErr }

func (*MockBatchResults) Exec() (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, nil
}

func (*MockBatchResults) Query() (pgx.Rows, error) { return nil, nil }

func (*MockBatchResults) QueryRow() pgx.Row { return nil }

type MockDBPool struct {
	sendBatchErrs []error
	callCount     int
}

func (*MockDBPool) Acquire(ctx context.Context) (*pgxpool.Conn, error) { return nil, nil }

func (*MockDBPool) Close() {}
func (*MockDBPool) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, nil
}
func (*MockDBPool) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return nil, nil
}
func (*MockDBPool) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row { return nil }
func (m *MockDBPool) SendBatch(ctx context.Context, batch *pgx.Batch) pgx.BatchResults {
	var err error
	if m.callCount < len(m.sendBatchErrs) {
		err = m.sendBatchErrs[m.callCount]
	}
	m.callCount++
	return &MockBatchResults{closeErr: err}
}

func TestClient_flushBatch(t *testing.T) {
	pgErr := &pgconn.PgError{Code: "40P01", Message: "deadlock detected"}
	ctx := context.Background()
	batch := &pgx.Batch{}
	// Add a dummy query so batch.Len() > 0
	batch.Queue("SELECT 1")

	tests := []struct {
		name            string
		sendBatchErrs   []error
		retryOnDeadlock bool
		wantErr         bool
	}{
		{
			name:            "happy path",
			sendBatchErrs:   []error{nil},
			retryOnDeadlock: true,
			wantErr:         false,
		},
		{
			name: "two retries then success",
			sendBatchErrs: []error{
				pgErr,
				pgErr,
				nil,
			},
			retryOnDeadlock: true,
			wantErr:         false,
		},
		{
			name: "six retries, always deadlock, fail",
			sendBatchErrs: []error{
				pgErr,
				pgErr,
				pgErr,
				pgErr,
				pgErr,
				pgErr,
			},
			retryOnDeadlock: true,
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				conn: &MockDBPool{sendBatchErrs: tt.sendBatchErrs},
				spec: &spec.Spec{RetryOnDeadlock: tt.retryOnDeadlock},
			}
			err := client.flushBatch(ctx, batch)
			if (err != nil) != tt.wantErr {
				t.Errorf("flushBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
