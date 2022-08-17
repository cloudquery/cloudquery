package postgresql

import (
	"context"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

var createTablesTests = []*schema.Table{
	{
		Name:    "empty_table",
		Columns: nil,
	},
	{
		Name: "simple_table",
		Columns: schema.ColumnList{
			{
				Name: "id",
				Type: schema.TypeBigInt,
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
	if err := c.Configure(ctx,
		specs.DestinationSpec{
			Spec: &PostgreSqlSpec{
				ConnectionString: "postgres://postgres:pass@localhost:5432/postgres",
			},
		},
	); err != nil {
		t.Fatal(err)
	}
	if err := c.Migrate(ctx, "test", "v1.0.0", createTablesTests); err != nil {
		t.Fatal(err)
	}

}
