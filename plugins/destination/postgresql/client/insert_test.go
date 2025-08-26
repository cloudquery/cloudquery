package client

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
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
		retryOnDeadlock int64
		wantErr         bool
	}{
		{
			name:            "happy path",
			sendBatchErrs:   []error{nil},
			retryOnDeadlock: 5,
			wantErr:         false,
		},
		{
			name: "two retries then success",
			sendBatchErrs: []error{
				pgErr,
				pgErr,
				nil,
			},
			retryOnDeadlock: 5,
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
			retryOnDeadlock: 5,
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

func TestConcurrentSyncsAgainstSameTable(t *testing.T) {
	const syncConcurrency = 10 // Lowered from 100 to 10 to work well with other tests
	const rounds = 99
	ctx := context.Background()
	group, _ := errgroup.WithContext(ctx)
	randomUUIDString := uuid.New().String()
	tableName := "k8s_core_namespaces_" + randomUUIDString

	table := &schema.Table{
		Name: tableName,
		Columns: []schema.Column{
			{Name: "id", Type: arrow.BinaryTypes.String, NotNull: true},
			{Name: "name", Type: arrow.BinaryTypes.String, PrimaryKey: true},
			schema.CqSyncTimeColumn,
		},
	}
	// Create the table
	migratePlugin := plugin.NewPlugin("postgresql",
		internalPlugin.Version,
		New,
		plugin.WithJSONSchema(spec.JSONSchema),
	)
	s := &spec.Spec{ConnectionString: getTestConnection(), BatchSize: 1, RetryOnDeadlock: 5}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	err = migratePlugin.Init(ctx, b, plugin.NewClientOptions{})
	require.NoError(t, err)
	migrateContext := context.Background()
	if err := migratePlugin.WriteAll(migrateContext, []message.WriteMessage{&message.WriteMigrateTable{Table: table}}); err != nil {
		t.Fatal(fmt.Errorf("failed to create table: %w", err))
	}

	for range syncConcurrency {
		group.Go(func() error {
			// Simulate a sync job against the same table
			syncContext := context.Background()
			p := plugin.NewPlugin("postgresql",
				internalPlugin.Version,
				New,
				plugin.WithJSONSchema(spec.JSONSchema),
			)
			p.SetLogger(zerolog.New(zerolog.NewTestWriter(t)).Level(zerolog.WarnLevel))
			s := &spec.Spec{ConnectionString: getTestConnection()}
			b, err := json.Marshal(s)
			require.NoError(t, err)
			err = p.Init(syncContext, b, plugin.NewClientOptions{})
			require.NoError(t, err)

			for range rounds {
				jobIndexAsString := fmt.Sprintf("%02d", 1)
				randomUUIDStringWithLastCharacterReplaced := randomUUIDString[:len(randomUUIDString)-len(jobIndexAsString)] + jobIndexAsString
				bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
				bldr.Field(0).(*array.StringBuilder).Append(uuid.MustParse(randomUUIDStringWithLastCharacterReplaced).String())
				bldr.Field(1).(*array.StringBuilder).Append("source")
				bldr.Field(2).(*array.TimestampBuilder).Append(arrow.Timestamp(time.Now().UnixMicro()))
				record := bldr.NewRecord()

				if err := p.WriteAll(syncContext, []message.WriteMessage{&message.WriteInsert{
					Record: record,
				}}); err != nil {
					t.Fatal(fmt.Errorf("failed to insert record: %w", err))
				}
			}

			return nil
		})
	}

	require.NoError(t, group.Wait())

	ch := make(chan arrow.Record)
	go func() {
		defer close(ch)
		err = migratePlugin.Read(ctx, table, ch)
	}()

	numRows := 0
	for record := range ch {
		numRows += int(record.NumRows())
	}

	require.Equal(t, 1, numRows)
	require.NoError(t, err)
}
