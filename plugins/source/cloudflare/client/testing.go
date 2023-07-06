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

const (
	TestAccountID = "test_account"
	TestZoneID    = "test_zone"
)

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Clients) {
	t.Helper()
	table.IgnoreInTests = false

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))
	clients := builder(t, ctrl)
	c := New(l, clients, clients[TestAccountID], AccountZones{
		TestAccountID: {
			AccountId: TestAccountID,
			Zones:     []string{TestZoneID},
		},
	})

	messages, err := sched.SyncAll(context.Background(), &c, schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	records := messages.GetInserts().GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}
