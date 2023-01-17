package client

import (
	"context"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

func TestPlugin(t *testing.T) {
	p := destination.NewPlugin("sqlite", "development", New)

	destination.PluginTestSuiteRunner(t, p,
		Spec{
			ConnectionString: ":memory:",
		},
		destination.PluginTestSuiteTests{})
}

func TestPluginMigrateMultiplePKs(t *testing.T) {
	table := schema.Table{
		Name: "table_1",
		Columns: []schema.Column{
			{Name: "id", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}},
			{Name: "name", Type: schema.TypeString, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}},
			{Name: "age", Type: schema.TypeInt},
		},
	}
	p := destination.NewPlugin("sqlite", "development", New)
	ctx := context.Background()

	spec := Spec{
		ConnectionString: ":memory:",
	}

	// Init the plugin so we can call migrate
	if err := p.Init(ctx, zerolog.Logger{}, specs.Destination{Name: "cq_test_migrate_multiple_pks", Spec: spec}); err != nil {
		t.Fatal(err)
	}

	if err := p.Migrate(ctx, []*schema.Table{&table}); err != nil {
		t.Fatal(err)
	}

	if err := p.Migrate(ctx, []*schema.Table{&table}); err != nil {
		t.Fatal(err)
	}
}
