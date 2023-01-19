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
	tests := []struct {
		name    string
		spec    specs.Destination
		wantErr string
	}{
		{
			name: "should fail on migrate mode safe",
			spec: specs.Destination{Name: "cq_test_migrate", Spec: Spec{ConnectionString: ":memory:"}},
			wantErr: `failed to migrate schema:
can't migrate table "table_1" since adding the new PK column "new_pk_column" is not supported. Try dropping this table
can't migrate table "table_2" since changing the type of column "age" from "integer" to "text" is not supported. Try dropping this column for this table
can't migrate table "table_2" since changing the type of column "created_at" from "text" to "timestamp" is not supported. Try dropping this column for this table

To force a migration add "migrate_mode: forced" to your destination spec`,
		},
		{
			name: "should succeed on migrate mode force",
			spec: specs.Destination{Name: "cq_test_migrate", Spec: Spec{ConnectionString: ":memory:"}, MigrateMode: specs.MigrateModeForced},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table := schema.Table{
				Name: "table_1",
				Columns: []schema.Column{
					{Name: "id", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}},
					{Name: "name", Type: schema.TypeString},
					{Name: "age", Type: schema.TypeInt},
					{Name: "created_at", Type: schema.TypeString},
				},
			}
			p := destination.NewPlugin("sqlite", "development", New)
			ctx := context.Background()

			// Init the plugin so we can call migrate
			if err := p.Init(ctx, zerolog.Logger{}, tt.spec); err != nil {
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
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
