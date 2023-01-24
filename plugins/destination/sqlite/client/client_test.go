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
can't migrate table "table_with_new_pk_column" since adding the new PK column "new_pk_column" is not supported. Try dropping this table
can't migrate table "table_with_pk_addition_existing_column" since making the existing column "id" a PK is not supported. Try dropping this table
can't migrate table "table_with_pk_removal" since removing an existing column "id" as a PK is not supported. Try dropping this table
can't migrate table "table_with_pk_type_change" since changing the type of the PK column "id" from "integer" to "text" is not supported. Try dropping this table
can't migrate table "table_with_non_pk_type_change" since changing the type of column "created_at" from "text" to "timestamp" is not supported. Try dropping this column for this table`,
		},
		{
			name: "should succeed on migrate mode force",
			spec: specs.Destination{Name: "cq_test_migrate", Spec: Spec{ConnectionString: ":memory:"}, MigrateMode: specs.MigrateModeForced},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cqId := schema.CqIDColumn
			cqIdWithPK := cqId
			cqIdWithPK.CreationOptions.PrimaryKey = true
			beforeSchema := schema.Tables{
				{
					Name:    "table_with_new_column",
					Columns: []schema.Column{cqId, {Name: "id", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}}},
				},
				{
					Name: "table_with_new_pk_column",
					// cq_id is needed for initial migrations of tables without a PK
					Columns: []schema.Column{cqIdWithPK},
				},
				{
					Name: "table_with_pk_addition_existing_column",
					// cq_id is needed for initial migrations of tables without a PK
					Columns: []schema.Column{cqIdWithPK, {Name: "id", Type: schema.TypeUUID}},
				},
				{
					Name:    "table_with_pk_removal",
					Columns: []schema.Column{cqId, {Name: "id", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}}},
				},
				{
					Name:    "table_with_pk_type_change",
					Columns: []schema.Column{cqId, {Name: "id", Type: schema.TypeInt, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}}},
				},
				{
					Name: "table_with_non_pk_type_change",
					// cq_id is needed for initial migrations of tables without a PK
					Columns: []schema.Column{cqIdWithPK, {Name: "created_at", Type: schema.TypeString}},
				},
			}

			p := destination.NewPlugin("sqlite", "development", New)
			ctx := context.Background()

			// Init the plugin so we can call migrate
			if err := p.Init(ctx, zerolog.Logger{}, tt.spec); err != nil {
				t.Fatal(err)
			}

			if err := p.Migrate(ctx, schema.Tables{beforeSchema[0]}); err != nil {
				t.Fatal(err)
			}

			// Adding a new column to the table is the only safe migratable change in SQLite
			beforeSchema[0].Columns = append(beforeSchema[0].Columns, schema.Column{Name: "new_column", Type: schema.TypeString})
			if err := p.Migrate(ctx, beforeSchema); err != nil {
				t.Fatal(err)
			}

			afterSchema := schema.Tables{
				{
					Name:    "table_with_new_column",
					Columns: []schema.Column{cqId, {Name: "id", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}}, {Name: "new_column", Type: schema.TypeString}},
				},
				{
					Name:    "table_with_new_pk_column",
					Columns: []schema.Column{cqId, {Name: "new_pk_column", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}}},
				},
				{
					Name:    "table_with_pk_addition_existing_column",
					Columns: []schema.Column{cqId, {Name: "id", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}}},
				},
				{
					Name:    "table_with_pk_removal",
					Columns: []schema.Column{cqIdWithPK, {Name: "id", Type: schema.TypeUUID}},
				},
				{
					Name:    "table_with_pk_type_change",
					Columns: []schema.Column{cqId, {Name: "id", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}}},
				},
				{
					Name:    "table_with_non_pk_type_change",
					Columns: []schema.Column{cqIdWithPK, {Name: "created_at", Type: schema.TypeTimestamp}},
				},
			}

			err := p.Migrate(ctx, afterSchema)
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
