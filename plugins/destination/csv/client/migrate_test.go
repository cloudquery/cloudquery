package client

import (
	"context"
	"testing"

	"github.com/cloudquery/plugin-sdk/specs"
)

func TestMigrate(t *testing.T) {
	ctx := context.Background()
	tmpDir := t.TempDir()
	client, err := New(ctx, getTestLogger(t), specs.Destination{
		WriteMode: specs.WriteModeAppend,
		Spec: Spec{
			Directory: tmpDir,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	client.Migrate(ctx, nil)
	
}
