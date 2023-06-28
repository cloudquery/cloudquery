package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client/services"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	StartTime time.Time
}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T) services.HackernewsClient, opts TestOptions) {
	// version := "vDev"
	table.IgnoreInTests = false
	t.Helper()
	//ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	schedulerClient := &Client{
		logger:     l,
		HackerNews: builder(t),
		// Backend:    mockBackend, TODO(v4)
		maxRetries: 0,
		backoff:    1 * time.Millisecond,
	}
	tables := schema.Tables{table}
	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}
	sc := scheduler.NewScheduler(schedulerClient, scheduler.WithLogger(l))
	messages, err := sc.SyncAll(context.Background(), tables)
	if err != nil {
		t.Fatal(err)
	}
	inserts := messages.InsertMessage()
	records := inserts.GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
	//newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	//	startTime := ""
	//	if !opts.StartTime.IsZero() {
	//		startTime = opts.StartTime.Format(time.RFC3339)
	//	}
	//	return &Client{
	//		logger:     l,
	//		Backend:    opts.Backend,
	//		HackerNews: builder(t, ctrl),
	//		Spec:       Spec{ItemConcurrency: 10, StartTime: startTime},
	//	}, nil
	//}
	//p := source.NewPlugin(
	//	table.Name,
	//	version,
	//	[]*schema.Table{
	//		table,
	//	},
	//	newTestExecutionClient)
	//p.SetLogger(l)
	//source.TestPluginSync(t, p, specs.Source{
	//	Name:         "dev",
	//	Path:         "cloudquery/dev",
	//	Version:      version,
	//	Tables:       []string{table.Name},
	//	Destinations: []string{"mock-destination"},
	//})
}
