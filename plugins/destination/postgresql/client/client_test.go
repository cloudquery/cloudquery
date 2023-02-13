package client

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/cloudquery/plugin-sdk/testdata"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	return testConn
}

func TestPgPlugin(t *testing.T) {
	p := destination.NewPlugin("postgresql", "development", New)
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			ConnectionString: getTestConnection(),
			PgxLogLevel:      LogLevelTrace,
		},
		destination.PluginTestSuiteTests{})
}

func TestPgPluginPrimaryKeyRename(t *testing.T) {
	tableName := fmt.Sprintf("cq_test_pk_rename_%03d", rand.Intn(100))
	tableWithStalePk := testdata.TestTable(tableName)
	// We simulate that a primary column was renamed
	tableWithStalePk.Columns = append(tableWithStalePk.Columns, schema.Column{Name: "stale_pk_1", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}})
	tableWithStalePk.Columns = append(tableWithStalePk.Columns, schema.Column{Name: "stale_pk_2", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}})
	p := destination.NewPlugin("postgresql", "development", New)
	ctx := context.Background()

	spec := Spec{
		ConnectionString: getTestConnection(),
		PgxLogLevel:      LogLevelTrace,
	}

	// Init the plugin so we can call migrate
	if err := p.Init(ctx, zerolog.Logger{}, specs.Destination{Name: "stale_pk", Spec: spec}); err != nil {
		t.Fatal(err)
	}

	// Call migrate so the `stale_pk_1` and `stale_pk_2` columns are created
	if err := p.Migrate(ctx, []*schema.Table{tableWithStalePk}); err != nil {
		t.Fatal(err)
	}

	// Migrate the table again without the `stale_pk_1` and `stale_pk_2` columns
	table := testdata.TestTable(tableName)
	err := p.Migrate(ctx, []*schema.Table{table})
	expected := `the following primary keys were removed from the schema ["stale_pk_1" "stale_pk_2"] for table "%s".
You can migrate the table manually by running:
-------------------------------------------------------------------------------------------
alter table "%s" drop constraint if exists "%s_cqpk";
alter table "%s" alter column "stale_pk_1" drop not null;
alter table "%s" alter column "stale_pk_2" drop not null;
-------------------------------------------------------------------------------------------`
	require.ErrorContains(t, err, fmt.Sprintf(expected, tableName, tableName, tableName, tableName, tableName))
}
