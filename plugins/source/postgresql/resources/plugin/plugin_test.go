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

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/postgresql/client"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/scalar"
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

type testCase struct {
	typeName string
	value    any
	expect   scalar.Scalar
}

func getTestCases(serialValue int64) []testCase {
	cidr := scalar.Inet{}
	err := cidr.Set("10.1.2.3/32")
	if err != nil {
		panic(err)
	}
	mac1 := scalar.Mac{}
	err = mac1.Set("08:00:2b:01:02:03")
	if err != nil {
		panic(err)
	}
	mac2 := scalar.Mac{}
	err = mac2.Set("08:00:2b:01:02:03:04:05")
	if err != nil {
		panic(err)
	}
	inet := scalar.Inet{}
	err = inet.Set("192.168.0.1/24")
	if err != nil {
		panic(err)
	}
	timeMicrosecond := scalar.Time{
		Unit: arrow.Microsecond,
	}
	err = timeMicrosecond.Set("04:05:06.789000")
	if err != nil {
		panic(err)
	}
	timeMillisecond := scalar.Time{
		Int: scalar.Int{
			BitWidth: 32,
		},
		Unit: arrow.Millisecond,
	}
	err = timeMillisecond.Set("04:05:06.789")
	if err != nil {
		panic(err)
	}
	timestamp := scalar.Timestamp{}
	err = timestamp.Set("1999-01-08 04:05:06.789")
	if err != nil {
		panic(err)
	}
	timestampMillisecond := scalar.Timestamp{
		Type: &arrow.TimestampType{
			Unit:     arrow.Millisecond,
			TimeZone: "UTC",
		},
	}
	err = timestampMillisecond.Set("1999-01-08 04:05:06.789")
	if err != nil {
		panic(err)
	}
	uuidData := scalar.UUID{}
	err = uuidData.Set(uuid.New())
	if err != nil {
		panic(err)
	}
	return []testCase{
		{"int", 1, &scalar.Int{Value: 1, Valid: true, BitWidth: 32}},
		{"bigint", 1, &scalar.Int{Value: 1, Valid: true, BitWidth: 64}},
		{"bigserial", nil, &scalar.Int{Value: serialValue, Valid: true, BitWidth: 64}},
		{"bit", "1", &scalar.String{Value: "1", Valid: true}},
		{"bit(5)", "11111", &scalar.String{Value: "11111", Valid: true}},
		{"bit varying", "1", &scalar.String{Value: "1", Valid: true}},
		{"bit varying(5)", "11111", &scalar.String{Value: "11111", Valid: true}},
		{"boolean", true, &scalar.Bool{Value: true, Valid: true}},
		{"box", "((1,2),(3,4))", &scalar.String{Value: "(3,4),(1,2)", Valid: true}},
		{"bytea", []byte("test"), &scalar.Binary{Value: []byte("test"), Valid: true}},
		{"character", "a", &scalar.String{Value: "a", Valid: true}},
		{"character(5)", "aaaaa", &scalar.String{Value: "aaaaa", Valid: true}},
		{"character varying", "a", &scalar.String{Value: "a", Valid: true}},
		{"character varying(5)", "aaaaa", &scalar.String{Value: "aaaaa", Valid: true}},
		{"cidr", "10.1.2.3/32", &cidr},
		{"circle", "<(1,2),3>", &scalar.String{Value: "<(1,2),3>", Valid: true}},
		{"date", "1999-01-08", &scalar.Date32{Value: 10599, Valid: true}},
		{"double precision", 1.1, &scalar.Float{Value: 1.1, Valid: true, BitWidth: 64}},
		{"inet", "192.168.0.1/24", &inet},
		{"integer", 1, &scalar.Int{Value: 1, Valid: true, BitWidth: 32}},
		{"interval", "1-2", &scalar.String{Value: "14 mon 00:00:00.000000", Valid: true}},
		{"json", `{"a":1}`, &scalar.JSON{Value: []byte(`{"a":1}`), Valid: true}},
		{"jsonb", `{"a":1}`, &scalar.JSON{Value: []byte(`{"a":1}`), Valid: true}},
		{"line", "{1,2,3}", &scalar.String{Value: "{1,2,3}", Valid: true}},
		{"lseg", "[(1,2),(3,4)]", &scalar.String{Value: "[(1,2),(3,4)]", Valid: true}},
		{"macaddr", "08:00:2b:01:02:03", &mac1},
		{"macaddr8", "08:00:2b:01:02:03:04:05", &mac2},
		{"money", "$1,000.00", &scalar.String{Value: "$1,000.00", Valid: true}},
		{"path", `[(1,2),(3,4)]`, &scalar.String{Value: "[(1,2),(3,4)]", Valid: true}},
		{"point", "(1,2)", &scalar.String{Value: "(1,2)", Valid: true}},
		{"polygon", `((1,2),(3,4))`, &scalar.String{Value: "((1,2),(3,4))", Valid: true}},
		{"real", 1.1, &scalar.Float{Value: 1.100000023841858, Valid: true, BitWidth: 32}},
		{"smallint", 1, &scalar.Int{Value: 1, Valid: true, BitWidth: 16}},
		{"smallserial", nil, &scalar.Int{Value: serialValue, Valid: true, BitWidth: 16}},
		{"serial", nil, &scalar.Int{Value: serialValue, Valid: true, BitWidth: 32}},
		{"text", "test", &scalar.String{Value: "test", Valid: true}},
		{"time without time zone", "04:05:06.789", &timeMicrosecond},
		{"time(3)", "04:05:06.789", &timeMillisecond},
		{"time(3) without time zone", "04:05:06.789", &timeMillisecond},
		{"timestamp", "1999-01-08 04:05:06.789", &timestamp},
		{"timestamp without time zone", "1999-01-08 04:05:06.789", &timestamp},
		{"timestamp(3)", "1999-01-08 04:05:06.789", &timestampMillisecond},
		{"timestamp(3) without time zone", "1999-01-08 04:05:06.789", &timestampMillisecond},
		{"tsquery", "a & b", &scalar.String{Value: "'a' & 'b'", Valid: true}},
		{"tsvector", "'a':1 'b':2", &scalar.String{Value: "'a':1 'b':2", Valid: true}},
		{"uuid", &uuidData, &uuidData},
		{"xml", "<a>1</a>", &scalar.String{Value: "<a>1</a>", Valid: true}},
	}
}

func createTestTable(ctx context.Context, conn *pgxpool.Pool, tableName string) error {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE ")
	sb.WriteString(pgx.Identifier{tableName}.Sanitize())
	sb.WriteString(" (")
	columns := getTestCases(0)
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

func insertTestTable(ctx context.Context, conn *pgxpool.Pool, tableName string, testCases []testCase) error {
	var query = ""
	query += "INSERT INTO " + pgx.Identifier{tableName}.Sanitize() + " ("
	for _, col := range testCases {
		if col.value == nil {
			continue
		}
		query += pgx.Identifier{col.typeName + "_type"}.Sanitize() + ", "
	}
	query = query[:len(query)-2] + ") VALUES ("
	dataIndex := 0
	for _, col := range testCases {
		if col.value == nil {
			continue
		}
		query += "$" + fmt.Sprintf("%d", dataIndex+1) + ", "
		dataIndex++
	}
	query = query[:len(query)-2] + ")"
	pgData := make([]any, dataIndex)
	i := 0
	for _, col := range testCases {
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
	data := getTestCases(1)
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
	err = insertTestTable(ctx, conn, otherTable, getTestCases(2))
	if err != nil {
		t.Fatal(err)
	}

	// Init the plugin so we can call migrate
	if err := p.Init(ctx, spec); err != nil {
		t.Fatal(err)
	}
	res := make(chan *schema.Resource, 1)
	g := errgroup.Group{}
	syncTime := time.Now()
	g.Go(func() error {
		defer close(res)
		return p.Sync(ctx, syncTime, res)
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

	for i, got := range gotData {
		expected := data[i].expect
		if !got.Equal(expected) {
			t.Fatalf("type %v: expected %v, got %v", data[i].typeName, expected, got)
		}
	}
}

func TestPluginCDC(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.WarnLevel).With().Timestamp().Logger()
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
	if _, err := conn.Exec(ctx, "SELECT pg_drop_replication_slot('test_pg_source')"); err != nil {
		if !(errors.As(err, &pgErr) && pgErr.Code == "42704") {
			t.Fatal(err)
		}
	}

	testTable := "user"

	if err := createTestTable(ctx, conn, testTable); err != nil {
		t.Fatal(err)
	}
	data := getTestCases(1)
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

	syncTime := time.Now()
	go func() {
		defer wg.Done()
		defer close(res)
		syncErr = p.Sync(syncCtx, syncTime, res)
	}()
	data2 := getTestCases(2)
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
			for i, got := range gotData {
				expected := data[i].expect
				if !got.Equal(expected) {
					t.Fatalf("type %v: expected %v, got %v", data[i].typeName, expected, got)
				}
			}
		} else {
			for i, got := range gotData {
				expected := data2[i].expect
				if !got.Equal(expected) {
					t.Fatalf("type %v: expected %v, got %v", data[i].typeName, expected, got)
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
