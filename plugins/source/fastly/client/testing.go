package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/services"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/fastly/go-fastly/v7/fastly"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	Service *fastly.Service
	Region  string
}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) services.FastlyClient, opts TestOptions) {
	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	c := &Client{
		logger:   l,
		Fastly:   builder(t, ctrl),
		services: []*fastly.Service{opts.Service},
		regions:  []string{opts.Region},
	}
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))
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
