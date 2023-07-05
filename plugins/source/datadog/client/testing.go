package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct{}

func DatadogMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) DatadogServices, _ TestOptions) {
	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	c := &Client{
		logger:     l,
		DDServices: builder(t, ctrl),
		Accounts:   []Account{{Name: "test", APIKey: "test", AppKey: "test"}},
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

func MockPaginatedResponse[T any](result T) (<-chan datadog.PaginationResult[T], func()) {
	ch := make(chan datadog.PaginationResult[T], 1)
	ch <- datadog.PaginationResult[T]{Item: result}
	close(ch)
	return ch, func() {}
}
