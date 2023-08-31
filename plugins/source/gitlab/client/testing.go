package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/xanzy/go-gitlab"
)

type TestOptions struct{}

func GitlabMockTestHelper(t *testing.T, table *schema.Table, createService func(*httprouter.Router) error, options TestOptions) {
	t.Helper()
	table.IgnoreInTests = false
	mux := httprouter.New()
	ts := httptest.NewUnstartedServer(mux)
	defer ts.Close()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	err := createService(mux)
	if err != nil {
		t.Fatal(err)
	}
	ts.Start()
	client, err := gitlab.NewClient("",
		gitlab.WithBaseURL(ts.URL),
		// Disable backoff to speed up tests that expect errors.
		gitlab.WithCustomBackoff(func(_, _ time.Duration, _ int, _ *http.Response) time.Duration {
			return 0
		}),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	c := &Client{
		logger:  l,
		Gitlab:  client,
		BaseURL: ts.URL,
	}

	sched := scheduler.NewScheduler(scheduler.WithLogger(l))
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
