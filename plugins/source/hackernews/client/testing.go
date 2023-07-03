package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client/services"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	StartTime time.Time
	Backend   state.Client
}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) services.HackernewsClient, opts TestOptions) {
	table.IgnoreInTests = false
	t.Helper()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	startTimeStr := ""
	if !opts.StartTime.IsZero() {
		startTimeStr = opts.StartTime.Format(time.RFC3339)
	}
	schedulerClient := &Client{
		logger:     l,
		HackerNews: builder(t, gomock.NewController(t)),
		Backend:    opts.Backend,
		maxRetries: 0,
		backoff:    1 * time.Millisecond,
		Spec: Spec{
			ItemConcurrency: 10,
			StartTime:       startTimeStr,
		},
	}
	tables := schema.Tables{table}
	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}
	sc := scheduler.NewScheduler(scheduler.WithLogger(l))
	messages, err := sc.SyncAll(context.Background(), schedulerClient, tables)
	if err != nil {
		t.Fatal(err)
	}
	inserts := messages.InsertMessage()
	records := inserts.GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}
