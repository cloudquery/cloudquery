package plugin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/oracledb/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/cloudquery/plugin-sdk/testdata"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/sijms/go-ora/v2/network"
)

func getTestDB(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("oracle", getTestConnectionString())
	if err != nil {
		return nil, err
	}
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return db, err
}

func getTestConnectionString() string {
	testConn := os.Getenv("CQ_SOURCE_ORACLE_DB_TEST_CONNECTION_STRING")
	if testConn == "" {
		return "oracle://cq:test@localhost:1521/cloudquery"
	}
	return testConn
}

func createTable(ctx context.Context, db *sql.DB, table *schema.Table) error {
	builder := strings.Builder{}
	builder.WriteString("CREATE TABLE ")
	builder.WriteString(client.Identifier(table.Name))
	builder.WriteString(" (\n  ")
	for i, column := range table.Columns {
		builder.WriteString(client.Identifier(column.Name))
		builder.WriteString(" ")
		builder.WriteString(client.SQLType(column.Type))
		if column.CreationOptions.NotNull {
			builder.WriteString(" NOT NULL")
		}
		if column.CreationOptions.Unique {
			builder.WriteString(" UNIQUE")
		}
		if column.CreationOptions.PrimaryKey {
			builder.WriteString(" PRIMARY KEY")
		}
		if i < len(table.Columns)-1 {
			builder.WriteString(",\n  ")
		}
	}
	builder.WriteString("\n)")
	_, err := db.ExecContext(ctx, builder.String())
	return err
}

func insertTable(ctx context.Context, db *sql.DB, table *schema.Table, data schema.CQTypes) error {
	builder := strings.Builder{}
	builder.WriteString("INSERT INTO " + client.Identifier(table.Name))
	builder.WriteString(" (")
	for i, col := range table.Columns {
		builder.WriteString(client.Identifier(col.Name))
		if i < len(table.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(") VALUES (")
	for i := 0; i < len(table.Columns); i++ {
		builder.WriteString(fmt.Sprintf(":%d", i+1))
		if i < len(table.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(")")
	dbData := schema.TransformWithTransformer(&client.Transformer{}, data)
	if _, err := db.ExecContext(ctx, builder.String(), dbData...); err != nil {
		return err
	}
	return nil
}

func isNotExistsError(err error) bool {
	var dbError *network.OracleError
	if errors.As(err, &dbError) {
		return dbError.ErrCode == 942
	}

	return false
}

func TestPlugin(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)
	spec := specs.Source{
		Name:         "test_oracledb_source",
		Path:         "cloudquery/oracledb",
		Version:      "vDevelopment",
		Destinations: []string{"test"},
		Tables:       []string{"test_oracledb_source"},
		Spec: client.Spec{
			ConnectionString: getTestConnectionString(),
		},
	}
	db, err := getTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := testdata.TestSourceTable("test_oracledb_source")
	if _, err := db.ExecContext(ctx, "DROP TABLE \"test_oracledb_source\" purge"); err != nil {
		if !isNotExistsError(err) {
			t.Fatal(err)
		}
	}
	if err := createTable(ctx, db, testTable); err != nil {
		t.Fatal(err)
	}
	data := testdata.GenTestData(testTable)
	if err := insertTable(ctx, db, testTable, data); err != nil {
		t.Fatal(err)
	}

	otherTable := testdata.TestSourceTable("other_oracledb_table")
	if _, err := db.ExecContext(ctx, "DROP TABLE \"other_oracledb_table\" purge"); err != nil {
		if !isNotExistsError(err) {
			t.Fatal(err)
		}
	}
	if err := createTable(ctx, db, otherTable); err != nil {
		t.Fatal(err)
	}
	otherData := testdata.GenTestData(otherTable)
	if err := insertTable(ctx, db, otherTable, otherData); err != nil {
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
	actualStrings := make([]string, len(gotData))
	for i, v := range gotData {
		actualStrings[i] = v.String()
	}
	expectedStrings := make([]string, len(data))
	for i, v := range data {
		expectedStrings[i] = v.String()
	}
	require.Equal(t, expectedStrings, actualStrings)
}
