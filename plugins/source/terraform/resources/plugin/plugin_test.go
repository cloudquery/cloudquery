package plugin

import (
	"context"
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestTerraform(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	pth := filepath.Dir(filename)
	p := Terraform()

	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(logger)

	cfg := map[string]string{"path": path.Join(pth, "testdata/terraform.tfstate")}
	cfgBytes, err := json.Marshal(cfg)
	raw := json.RawMessage(cfgBytes)
	require.NoError(t, err)
	spec := &client.Spec{
		Backends: []client.BackendConfigBlock{
			{
				BackendName: "mylocal",
				Type:        "local",
				Config:      &raw,
			},
		},
	}

	specBytes, err := json.Marshal(spec)
	require.NoError(t, err)
	require.NoError(t, p.Init(context.Background(), specBytes, plugin.NewClientOptions{}))

	tables, err := p.Tables(context.Background(), plugin.TableOptions{Tables: []string{"*"}})
	require.NoError(t, err)
	require.NotEmpty(t, tables)

	messages, err := p.SyncAll(context.Background(), plugin.SyncOptions{Tables: []string{"*"}})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}

	for _, table := range tables {
		records := messages.GetInserts().GetRecordsForTable(table)
		emptyColumns := schema.FindEmptyColumns(table, records)
		if len(emptyColumns) > 0 {
			t.Fatalf("empty columns: %v", emptyColumns)
		}
	}
}
