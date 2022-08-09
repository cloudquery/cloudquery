package destinations

import (
	"context"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/spec"
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
	p := PostgreSqlPlugin{}
	if err := p.Configure(ctx, l, &spec.DestinationSpec{
		Name:    "postgresql",
		Version: "test",
		Spec: &PostgreSqlSpec{
			ConnectionString: "postgres://postgres:pass@localhost:5432/postgres",
		},
	}); err != nil {
		t.Fatal("failed to configure postgresql plugin:", err)
	}
	if err := p.CreateTables(ctx, createTablesTests); err != nil {
		t.Fatal("failed to create tables:", err)
	}
}
