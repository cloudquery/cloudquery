package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

func getTestLogger(t *testing.T) zerolog.Logger {
	t.Helper()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	return zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}

func TestInitialize(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, getTestLogger(t), specs.Destination{
		Spec: Spec{
			ConnectionString: ":memory:",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if client == nil {
		t.Fatal("client is nil")
	}
	if err := client.Close(ctx); err != nil {
		t.Fatal(err)
	}
	err = client.Close(ctx)
	if err == nil {
		t.Fatal("expected error when closing a closed client second time")
	}

	if err.Error() != "client already closed or not initialized" {
		t.Fatal("expected error when closing a closed client second time")
	}
}

func TestPgPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugins.NewDestinationPlugin("postgresql", "development", New)

	if err := plugins.DestinationPluginTestHelper(ctx, p, getTestLogger(t), specs.Destination{
		WriteMode: specs.WriteModeAppend,
		Spec: Spec{
			ConnectionString: ":memory:",
		},
	}); err != nil {
		t.Fatal(err)
	}
}
