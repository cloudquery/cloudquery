package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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

func MockPaginatedResponse[T any](result T) (<-chan datadog.PaginationResult[T], func()) {
	ch := make(chan datadog.PaginationResult[T], 1)
	ch <- datadog.PaginationResult[T]{Item: result}
	close(ch)
	return ch, func() {}
}
