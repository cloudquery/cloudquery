package postgresql

import (
	"context"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog"
)

var createTablesTests = []*schema.Table{
	{
		Name:    "empty_table",
		Columns: nil,
	},
	{
		Name: "simple_table",
		Options: schema.TableCreationOptions{
			PrimaryKeys: []string{"id"},
		},
		Columns: schema.ColumnList{
			{
				Name: "id",
				Type: schema.TypeInt,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
		},
	},
}

func TestPostgreSqlCreateTables(t *testing.T) {
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	ctx := context.Background()
	c := NewClient(l)
	if err := c.Initialize(ctx,
		specs.Destination{
			Spec: &PostgreSqlSpec{
				ConnectionString: "postgres://postgres:pass@localhost:5432/postgres",
				PgxLogLevel:      pgx.LogLevelInfo,
			},
		},
	); err != nil {
		t.Fatal(err)
	}
	if err := c.Migrate(ctx, createTablesTests); err != nil {
		t.Fatal(err)
	}

}
