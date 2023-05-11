package plugin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/postgresql/client"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/google/uuid"
	pgx_zero_log "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
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

type pgTypeToValue struct {
	typeName string
	value    any
}

func getPgTypesToData() []pgTypeToValue {
	var pgTypesToData = []pgTypeToValue{
		{"bigint", 1},
		{"bigserial", nil},
		{"bit", "1"},
		{"bit(5)", "11111"},
		{"bit varying", "1"},
		{"bit varying(5)", "11111"},
		{"boolean", true},
		{"box", "((1,2),(3,4))"},
		{"bytea", []byte("test")},
		{"character", "a"},
		{"character(5)", "aaaaa"},
		{"character varying", "a"},
		{"character varying(5)", "aaaaa"},
		{"cidr", "10.1.2.3/32"},
		{"circle", "<(1,2),3>"},
		{"date", "1999-01-08"},
		{"double precision", 1.1},
		{"inet", "192.168.0.1/24"},
		{"integer", 1},
		{"interval", "1-2"},
		{"json", `{"a":1}`},
		{"jsonb", `{"a":1}`},
		{"line", "{1,2,3}"},
		{"lseg", "[(1,2),(3,4)]"},
		{"macaddr", "08:00:2b:01:02:03"},
		{"macaddr8", "08:00:2b:01:02:03:04:05"},
		{"money", "$1,000.00"},
		{"path", `[(1,2),(3,4)]`},
		{"point", "(1,2)"},
		{"polygon", `((1,2),(3,4))`},
		{"real", 1.1},
		{"smallint", 1},
		{"smallserial", nil},
		{"serial", nil},
		{"text", "test"},
		{"time without time zone", "04:05:06.789"},
		{"time(3)", "04:05:06.789"},
		{"time(3) without time zone", "04:05:06.789"},
		{"timestamp", "1999-01-08 04:05:06.789"},
		{"timestamp without time zone", "1999-01-08 04:05:06.789"},
		{"timestamp(3)", "1999-01-08 04:05:06.789"},
		{"timestamp(3) without time zone", "1999-01-08 04:05:06.789"},
		{"tsquery", "a & b"},
		{"tsvector", "'a':1 'b':2"},
		{"uuid", uuid.New().String()},
		{"xml", "<a>1</a>"},
	}

	return pgTypesToData
}

func createTestTable(ctx context.Context, conn *pgxpool.Pool, tableName string) error {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE ")
	sb.WriteString(pgx.Identifier{tableName}.Sanitize())
	sb.WriteString(" (")
	columns := getPgTypesToData()
	for i, col := range columns {
		sb.WriteString(pgx.Identifier{col.typeName + "_type"}.Sanitize())
		sb.WriteString(" ")
		sb.WriteString(col.typeName)
		if col.typeName == "uuid" {
			sb.WriteString(" PRIMARY KEY")
		}
		if i < len(columns)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	if _, err := conn.Exec(ctx, sb.String()); err != nil {
		return err
	}
	return nil
}

func insertTestTable(ctx context.Context, conn *pgxpool.Pool, tableName string, columns []pgTypeToValue) error {
	var query = ""
	query += "INSERT INTO " + pgx.Identifier{tableName}.Sanitize() + " ("
	for _, col := range columns {
		if col.value == nil {
			continue
		}
		query += pgx.Identifier{col.typeName + "_type"}.Sanitize() + ", "
	}
	query = query[:len(query)-2] + ") VALUES ("
	dataIndex := 0
	for _, col := range columns {
		if col.value == nil {
			continue
		}
		query += "$" + fmt.Sprintf("%d", dataIndex+1) + ", "
		dataIndex++
	}
	query = query[:len(query)-2] + ")"
	pgData := make([]any, dataIndex)
	i := 0
	for _, col := range columns {
		if col.value == nil {
			continue
		}
		pgData[i] = col.value
		i++
	}
	if _, err := conn.Exec(ctx, query, pgData...); err != nil {
		return err
	}

	return nil
}

func getExpectedData(uuidValue any, serialValue int64) schema.CQTypes {
	cidr := schema.CIDR{}
	_ = cidr.Set("10.1.2.3/32")
	mac1 := schema.Macaddr{}
	_ = mac1.Set("08:00:2b:01:02:03")
	mac2 := schema.Macaddr{}
	_ = mac2.Set("08:00:2b:01:02:03:04:05")
	inet := schema.Inet{}
	_ = inet.Set("192.168.0.1/24")
	timestamp := schema.Timestamptz{}
	_ = timestamp.Set("1999-01-08 04:05:06.789")
	uuidData := schema.UUID{}
	_ = uuidData.Set(uuidValue)
	expectedData := schema.CQTypes{
		&schema.Int8{Int: 1, Status: schema.Present},
		&schema.Int8{Int: serialValue, Status: schema.Present},
		&schema.Text{Str: "1", Status: schema.Present},
		&schema.Text{Str: "11111", Status: schema.Present},
		&schema.Text{Str: "1", Status: schema.Present},
		&schema.Text{Str: "11111", Status: schema.Present},
		&schema.Bool{Bool: true, Status: schema.Present},
		&schema.Text{Str: "(3,4),(1,2)", Status: schema.Present},
		&schema.Bytea{Bytes: []byte("test"), Status: schema.Present},
		&schema.Text{Str: "a", Status: schema.Present},
		&schema.Text{Str: "aaaaa", Status: schema.Present},
		&schema.Text{Str: "a", Status: schema.Present},
		&schema.Text{Str: "aaaaa", Status: schema.Present},
		&cidr,
		&schema.Text{Str: "<(1,2),3>", Status: schema.Present},
		&schema.Text{Str: "1999-01-08 00:00:00 +0000 UTC", Status: schema.Present},
		&schema.Float8{Float: 1.1, Status: schema.Present},
		&inet,
		&schema.Int8{Int: 1, Status: schema.Present},
		&schema.Text{Str: "14 mon 00:00:00.000000", Status: schema.Present},
		&schema.JSON{Bytes: []byte(`{"a":1}`), Status: schema.Present},
		&schema.JSON{Bytes: []byte(`{"a":1}`), Status: schema.Present},
		&schema.Text{Str: "{1,2,3}", Status: schema.Present},
		&schema.Text{Str: "[(1,2),(3,4)]", Status: schema.Present},
		&mac1,
		&mac2,
		&schema.Text{Str: "$1,000.00", Status: schema.Present},
		&schema.Text{Str: "[(1,2),(3,4)]", Status: schema.Present},
		&schema.Text{Str: "(1,2)", Status: schema.Present},
		&schema.Text{Str: "((1,2),(3,4))", Status: schema.Present},
		&schema.Float8{Float: 1.100000023841858, Status: schema.Present},
		&schema.Int8{Int: 1, Status: schema.Present},
		&schema.Int8{Int: serialValue, Status: schema.Present},
		&schema.Int8{Int: serialValue, Status: schema.Present},
		&schema.Text{Str: "test", Status: schema.Present},
		&schema.Text{Str: "04:05:06.789000", Status: schema.Present},
		&schema.Text{Str: "04:05:06.789000", Status: schema.Present},
		&schema.Text{Str: "04:05:06.789000", Status: schema.Present},
		&timestamp,
		&timestamp,
		&timestamp,
		&timestamp,
		&schema.Text{Str: "'a' & 'b'", Status: schema.Present},
		&schema.Text{Str: "'a':1 'b':2", Status: schema.Present},
		&uuidData,
		&schema.Text{Str: "<a>1</a>", Status: schema.Present},
	}
	return expectedData
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
	defer conn.Close()

	testTable := "test_pg_source"
	if _, err := conn.Exec(ctx, "DROP TABLE IF EXISTS test_pg_source"); err != nil {
		t.Fatal(err)
	}
	if err := createTestTable(ctx, conn, testTable); err != nil {
		t.Fatal(err)
	}
	data := getPgTypesToData()
	err = insertTestTable(ctx, conn, testTable, data)
	if err != nil {
		t.Fatal(err)
	}

	otherTable := "other_pg_table"
	if _, err := conn.Exec(ctx, "DROP TABLE IF EXISTS other_pg_table"); err != nil {
		t.Fatal(err)
	}
	if err := createTestTable(ctx, conn, otherTable); err != nil {
		t.Fatal(err)
	}
	err = insertTestTable(ctx, conn, otherTable, getPgTypesToData())
	if err != nil {
		t.Fatal(err)
	}

	// Init the plugin so we can call migrate
	if err := p.Init(ctx, spec); err != nil {
		t.Fatal(err)
	}
	res := make(chan *schema.Resource, 1)
	g := errgroup.Group{}
	g.Go(func() error {
		defer close(res)
		return p.Sync(ctx, res)
	})
	var resource *schema.Resource
	totalResources := 0
	for r := range res {
		resource = r
		totalResources++
	}
	err = g.Wait()
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	if totalResources != 1 {
		t.Fatalf("expected 1 resource, got %d", totalResources)
	}
	gotData := resource.GetValues()
	expectedData := getExpectedData(data[44].value, 1)

	for i, v := range gotData {
		expected := expectedData[i]
		if !reflect.DeepEqual(v, expected) {
			t.Fatalf("expected %v, got %v", expected, v)
		}
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
		// use a reserved keyword as table name to validate escaping
		Tables: []string{"user"},
		Spec: &client.Spec{
			ConnectionString: getTestConnectionString() + "&replication=database",
			PgxLogLevel:      client.LogLevelTrace,
		},
	}
	conn, err := getTestConnection(ctx, l, getTestConnectionString())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if _, err := conn.Exec(ctx, "DROP TABLE IF EXISTS \"user\""); err != nil {
		t.Fatal(err)
	}
	var pgErr *pgconn.PgError
	if _, err := conn.Exec(ctx, "SELECT pg_drop_replication_slot('\"user\"')"); err != nil {
		if !(errors.As(err, &pgErr) && pgErr.Code == "42704") {
			t.Fatal(err)
		}
	}

	testTable := "user"

	if err := createTestTable(ctx, conn, testTable); err != nil {
		t.Fatal(err)
	}
	data := getPgTypesToData()
	err = insertTestTable(ctx, conn, testTable, data)
	if err != nil {
		t.Fatal(err)
	}

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
	data2 := getPgTypesToData()
	time.AfterFunc(2*time.Second, func() {
		err = insertTestTable(ctx, conn, testTable, data2)
		if err != nil {
			t.Fatal(err)
		}
	})

	totalResources := 0
	for r := range res {
		gotData := r.GetValues()
		if totalResources == 0 {
			expectedData := getExpectedData(data[44].value, 1)
			for i, v := range gotData {
				expected := expectedData[i]
				if !reflect.DeepEqual(v, expected) {
					t.Fatalf("expected %v, got %v", expected, v)
				}
			}
		} else {
			expectedData := getExpectedData(data2[44].value, 2)
			for i, v := range gotData {
				expected := expectedData[i]
				if !reflect.DeepEqual(v, expected) {
					t.Fatalf("expected %v, got %v", expected, v)
				}
			}
		}
		totalResources++
	}
	wg.Wait()
	if totalResources != 2 {
		t.Fatalf("expected 2 resource, got %d", totalResources)
	}
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
