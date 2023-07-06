package client

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/julienschmidt/httprouter"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func MockTestHelper(t *testing.T, table *schema.Table, createService func(*httprouter.Router) error) {
	t.Helper()

	mux := httprouter.New()
	ts := httptest.NewUnstartedServer(mux)
	defer ts.Close()

	require.NoError(t, createService(mux))
	ts.Start()

	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))

	snykClient := snyk.NewClient("test-key", snyk.WithBaseURL(ts.URL+"/"))

	c := &Client{
		Client: snykClient,
		logger: l,
		Organizations: []snyk.Organization{
			{ID: "test-org-id", Name: "test-org-name", Group: &snyk.Group{
				ID:   "test-group-id",
				Name: "test-group-name",
			}},
		},
	}

	Transform(schema.Tables{table})
	table.IgnoreInTests = false
	messages, err := sched.SyncAll(context.Background(), c, schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	records := messages.GetInserts().GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}
