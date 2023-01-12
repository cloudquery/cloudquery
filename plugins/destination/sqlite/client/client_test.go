package client

import (
	"context"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
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

func TestMigrateErrors(t *testing.T) {
	table := schema.Table{
		Name: "table_1",
		Columns: []schema.Column{
			{Name: "id", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}},
			{Name: "name", Type: schema.TypeString},
			{Name: "age", Type: schema.TypeInt},
			{Name: "created_at", Type: schema.TypeString},
		},
	}
	p := destination.NewPlugin("postgresql", "development", New)
	ctx := context.Background()

	spec := Spec{
		ConnectionString: ":memory:",
	}

	// Init the plugin so we can call migrate
	if err := p.Init(ctx, zerolog.Logger{}, specs.Destination{Name: "cq_test_migrate", Spec: spec}); err != nil {
		t.Fatal(err)
	}

	if err := p.Migrate(ctx, []*schema.Table{&table}); err != nil {
		t.Fatal(err)
	}

	tableWithMigratableChange := schema.Table{
		Name: "table_1",
		Columns: []schema.Column{
			{Name: "id", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}},
			{Name: "name", Type: schema.TypeString},
			{Name: "age", Type: schema.TypeInt},
			{Name: "created_at", Type: schema.TypeString},
			{Name: "new_column", Type: schema.TypeString},
		},
	}
	newTable := schema.Table{
		Name: "table_2",
		Columns: []schema.Column{
			{Name: "id", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}},
			{Name: "name", Type: schema.TypeString},
			{Name: "age", Type: schema.TypeInt},
			{Name: "created_at", Type: schema.TypeString},
		},
	}

	if err := p.Migrate(ctx, []*schema.Table{&tableWithMigratableChange, &newTable}); err != nil {
		t.Fatal(err)
	}

	tableWithNonTableDropNeeded := schema.Table{
		Name: "table_1",
		Columns: []schema.Column{
			{Name: "id", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}},
			{Name: "name", Type: schema.TypeString},
			{Name: "age", Type: schema.TypeInt},
			{Name: "new_pk_column", Type: schema.TypeString, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}},
			{Name: "created_at", Type: schema.TypeTimestamp},
		},
	}

	tableWithColumnsDropNeeded := schema.Table{
		Name: "table_2",
		Columns: []schema.Column{
			{Name: "id", Type: schema.TypeUUID},
			{Name: "name", Type: schema.TypeString, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}},
			{Name: "age", Type: schema.TypeString},
			{Name: "created_at", Type: schema.TypeTimestamp},
		},
	}

	err := p.Migrate(ctx, []*schema.Table{&tableWithNonTableDropNeeded, &tableWithColumnsDropNeeded})
	expectedError := `failed to migrate schema:
	can't migrate table "table_1" since adding the new PK column "new_pk_column" is not supported. Try dropping this table
	can't migrate table "table_2" since changing the type of column "age" from "integer" to "text" is not supported. Try dropping this column for this table
	can't migrate table "table_2" since changing the type of column "created_at" from "text" to "timestamp" is not supported. Try dropping this column for this table
	`
	require.Errorf(t, err, expectedError)
}
