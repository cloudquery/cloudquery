package plugin

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/postgresql/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/cloudquery/plugin-sdk/testdata"
	"github.com/rs/zerolog"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_SOURCE_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5433/postgres?sslmode=disable"
	}
	return testConn
}

func create_test_table(ctx context.Context, c *client.Client, table *schema.Table) error {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE test_pg_source (")
	for i, col := range table.Columns {
		sb.WriteString(col.Name)
		sb.WriteString(" ")
		sb.WriteString(SchemaTypeToPg10(col.Type))
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

func insert_test_table(ctx context.Context, c *client.Client, table *schema.Table, data schema.CQTypes) error {
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
	if _, err := c.Conn.Exec(ctx, sb.String(), pgData...); err != nil {
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
			ConnectionString: getTestConnection(),
			PgxLogLevel:      client.LogLevelTrace,
		},
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
	if err := insert_test_table(ctx, testClient, testTable, data); err != nil {
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
		fmt.Println("got: ", gotData)
		t.Fatalf("got unexpected data: %s", diff)
	}
}

func SchemaTypeToPg10(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "boolean"
	case schema.TypeInt:
		return "bigint"
	case schema.TypeFloat:
		return "double precision"
	case schema.TypeUUID:
		return "uuid"
	case schema.TypeString:
		return "text"
	case schema.TypeByteArray:
		return "bytea"
	case schema.TypeStringArray:
		return "text[]"
	case schema.TypeTimestamp:
		return "timestamp without time zone"
	case schema.TypeJSON:
		return "jsonb"
	case schema.TypeUUIDArray:
		return "uuid[]"
	case schema.TypeCIDR:
		return "cidr"
	case schema.TypeCIDRArray:
		return "cidr[]"
	case schema.TypeMacAddr:
		return "macaddr"
	case schema.TypeMacAddrArray:
		return "macaddr[]"
	case schema.TypeInet:
		return "inet"
	case schema.TypeInetArray:
		return "inet[]"
	case schema.TypeIntArray:
		return "bigint[]"
	default:
		panic("unknown type " + t.String())
	}
}
