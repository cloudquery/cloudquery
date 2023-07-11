package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestHelper(t *testing.T, table *schema.Table, createService func(t *testing.T, ctrl *gomock.Controller) Services, options TestOptions) {
	t.Helper()
	table.IgnoreInTests = false
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))

	spec := &Spec{
		Token: "test-token",
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		t.Fatalf("failed to validate spec: %v", err)
	}

	ctrl := gomock.NewController(t)
	services := createService(t, ctrl)

	c := &Client{
		logger:       l,
		SpacesRegion: firstSpacesRegion,
		Services:     &services,
	}

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
