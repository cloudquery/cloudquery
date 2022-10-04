package client

import (
	"context"
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func TestDrop(t *testing.T) {
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

	if err := c.Drop(ctx, testTables); err != nil {
		t.Fatalf("failed to drop tables second time: %v", err)
	}

	if err := c.Migrate(ctx, testTables); err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}

	isExist, err := c.isTableExistSQL(ctx, testTable.Name)
	if err != nil {
		t.Fatalf("failed to check if table exists: %v", err)
	}
	if !isExist {
		t.Fatalf("failed to migrate table. table %s doesn't exist", testTable.Name)
	}

	if err := c.Drop(ctx, testTables); err != nil {
		t.Fatalf("failed to drop tables after migration: %v", err)
	}

	isExist, err = c.isTableExistSQL(ctx, testTable.Name)
	if err != nil {
		t.Fatalf("failed to check if table exists: %v", err)
	}
	if isExist {
		t.Fatalf("failed to migrate table. table %s stil exists", testTable.Name)
	}
}
