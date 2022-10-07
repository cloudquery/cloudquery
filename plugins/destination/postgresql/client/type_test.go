package client

import (
	"context"
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func TestTypes(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, getTestLogger(t), specs.Destination{
		WriteMode: specs.WriteModeOverwriteDeleteStale,
		Spec: &Spec{
			ConnectionString: getTestConnection(),
		},
	})
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}
	c := client.(*Client)
	testTable := getTestTable()
	testTables := []*schema.Table{testTable}
	// check migration logic
	if err := c.Drop(ctx, testTables); err != nil {
		t.Fatalf("failed to drop tables: %v", err)
	}
	if err := c.Migrate(ctx, testTables); err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}
	pgColumns, err := c.getPgTableColumns(ctx, testTable.Name)
	if err != nil {
		t.Fatalf("failed to get table columns: %v", err)
	}
	for _, column := range testTable.Columns {
		pgColumn := pgColumns.getPgColumn(column.Name)
		if pgColumn == nil {
			t.Fatalf("failed to find column %s in table %s", column.Name, testTable.Name)
		}

		if pgColumn.typ != SchemaTypeToPg(column.Type) {
			t.Fatalf("failed to migrate table. column %s pg type is %s but expected %s", column.Name, pgColumn.typ, SchemaTypeToPg(column.Type))
		}
	}
}
