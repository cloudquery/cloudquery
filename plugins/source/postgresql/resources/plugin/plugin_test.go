package plugin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/postgresql/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/cloudquery/plugin-sdk/testdata"
	pgx_zero_log "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

func getTestConnection(ctx context.Context, logger zerolog.Logger, connectionString string) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string %w", err)
	}
	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return nil
	}
	pgxConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgx_zero_log.NewLogger(logger),
		LogLevel: tracelog.LogLevelTrace,
	}
	// maybe expose this to the user?
	pgxConfig.ConnConfig.RuntimeParams["timezone"] = "UTC"
	conn, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgresql: %w", err)
	}
	return conn, nil
}

func getTestConnectionString() string {
	testConn := os.Getenv("CQ_SOURCE_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	return testConn
}

func create_test_table(ctx context.Context, c *client.Client, table *schema.Table) error {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE ")
	sb.WriteString(table.Name)
	sb.WriteString(" (")
	for i, col := range table.Columns {
		sb.WriteString(col.Name)
		sb.WriteString(" ")
		sb.WriteString(c.SchemaTypeToPg(col.Type))
		if col.CreationOptions.PrimaryKey {
			sb.WriteString(" PRIMARY KEY")
		}
		if i < len(table.Columns)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	if _, err := c.Conn.Exec(ctx, sb.String()); err != nil {
		return err
	}
	return nil
}

func insert_test_table(ctx context.Context, conn *pgxpool.Pool, table *schema.Table, data schema.CQTypes) error {
	var sb strings.Builder
	sb.WriteString("INSERT INTO ")
	sb.WriteString(table.Name)
	sb.WriteString(" (")
	for i, col := range table.Columns {
		sb.WriteString(col.Name)
		if i < len(table.Columns)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(") VALUES (")
	for i := range table.Columns {
		sb.WriteString("$")
		sb.WriteString(fmt.Sprintf("%d", i+1))
		if i < len(table.Columns)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	pgData := schema.TransformWithTransformer(&client.Transformer{}, data)
	if _, err := conn.Exec(ctx, sb.String(), pgData...); err != nil {
		return err
	}
	return nil
}

func TestPlugin(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)
	spec := specs.Source{
		Name:         "test_pg_source",
		Path:         "cloudquery/postgresql",
		Version:      "vdevelopment",
		Destinations: []string{"test"},
		Tables:       []string{"test_pg_source"},
		Spec: client.Spec{
			ConnectionString: getTestConnectionString(),
			PgxLogLevel:      client.LogLevelTrace,
		},
	}
	conn, err := getTestConnection(ctx, l, spec.Spec.(client.Spec).ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	metaClient, err := client.Configure(ctx, l, spec, source.Options{})
	if err != nil {
		t.Fatal(err)
	}
	testClient := metaClient.(*client.Client)

	testTable := testdata.TestSourceTable("test_pg_source")
	if _, err := testClient.Conn.Exec(ctx, "DROP TABLE IF EXISTS test_pg_source"); err != nil {
		t.Fatal(err)
	}
	if err := create_test_table(ctx, testClient, testTable); err != nil {
		t.Fatal(err)
	}
	data := testdata.GenTestData(testTable)
	if err := insert_test_table(ctx, conn, testTable, data); err != nil {
		t.Fatal(err)
	}

	// Init the plugin so we can call migrate
	if err := p.Init(ctx, spec); err != nil {
		t.Fatal(err)
	}
	res := make(chan *schema.Resource, 10)
	if err := p.Sync(ctx, res); err != nil {
		t.Fatal(err)
	}
	close(res)
	var resource *schema.Resource
	totalResources := 0
	for r := range res {
		resource = r
		totalResources++
	}
	if totalResources != 1 {
		t.Fatalf("expected 1 resource, got %d", totalResources)
	}
	gotData := resource.GetValues()
	if diff := data.Diff(gotData); diff != "" {
		t.Fatalf("got unexpected data: %s", diff)
	}
}

func TestPluginCDC(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)
	spec := specs.Source{
		Name:         "test_pg_source",
		Path:         "cloudquery/postgresql",
		Version:      "vdevelopment",
		Destinations: []string{"test"},
		Tables:       []string{"test_pg_source"},
		Spec: &client.Spec{
			ConnectionString: getTestConnectionString() + "&replication=database",
			PgxLogLevel:      client.LogLevelTrace,
		},
	}
	conn, err := getTestConnection(ctx, l, getTestConnectionString())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := conn.Exec(ctx, "DROP TABLE IF EXISTS test_pg_source"); err != nil {
		t.Fatal(err)
	}
	var pgErr *pgconn.PgError
	if _, err := conn.Exec(ctx, "SELECT pg_drop_replication_slot('test_pg_source')"); err != nil {
		if !(errors.As(err, &pgErr) && pgErr.Code == "42704") {
			t.Fatal(err)
		}
	}

	testTable := testdata.TestSourceTable("test_pg_source")

	metaClient, err := client.Configure(ctx, l, spec, source.Options{})
	if err != nil {
		t.Fatal(err)
	}
	testClient := metaClient.(*client.Client)
	if err := create_test_table(ctx, testClient, testTable); err != nil {
		t.Fatal(err)
	}
	data := testdata.GenTestData(testTable)
	if err := insert_test_table(ctx, conn, testTable, data); err != nil {
		t.Fatal(err)
	}
	data2 := testdata.GenTestData(testTable)

	spec.Spec.(*client.Spec).CDC = true
	// Init the plugin so we can call migrate
	if err := p.Init(ctx, spec); err != nil {
		t.Fatal(err)
	}
	res := make(chan *schema.Resource, 10)
	var wg sync.WaitGroup
	var syncErr error
	syncCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(res)
		syncErr = p.Sync(syncCtx, res)
	}()
	time.AfterFunc(2*time.Second, func() {
		if err := insert_test_table(ctx, conn, testTable, data2); err != nil {
			t.Fatal(err)
		}
	})
	totalResources := 0
	for r := range res {
		gotData := r.GetValues()
		if totalResources == 0 {
			if diff := data.Diff(gotData); diff != "" {
				t.Fatalf("got unexpected data: %s", diff)
			}
		} else {
			if diff := data2.Diff(gotData); diff != "" {
				t.Fatalf("got unexpected data2: %s", diff)
			}
		}
		totalResources++
	}
	if totalResources != 2 {
		t.Fatalf("expected 2 resource, got %d", totalResources)
	}
	wg.Wait()
	if syncErr != nil && !IsContextDeadlineExceeded(syncErr) {
		t.Fatal(syncErr)
	}
}

func IsContextDeadlineExceeded(err error) bool {
	var deadlineExceeded bool
	for err != nil {
		if err == context.DeadlineExceeded {
			deadlineExceeded = true
			break
		}
		err = errors.Unwrap(err)
	}
	return deadlineExceeded
}
