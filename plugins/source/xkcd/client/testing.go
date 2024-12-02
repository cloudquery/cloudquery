package client

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/xkcd/internal/xkcd"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
)

func TestHelper(t *testing.T, table *schema.Table, ts *httptest.Server) {
	table.IgnoreInTests = false
	t.Helper()

	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))

	spec := &Spec{}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		t.Fatalf("failed to validate spec: %v", err)
	}

	xkcdClient, err := xkcd.NewClient(xkcd.WithBaseURL(ts.URL), xkcd.WithHTTPClient(ts.Client()))
	if err != nil {
		t.Fatal(err)
	}
	c := New(l, *spec, xkcdClient, nil)

	tables := schema.Tables{table}
	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}
	messages, err := sched.SyncAll(context.Background(), c, tables)
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	plugin.ValidateNoEmptyColumns(t, tables, messages)
}
