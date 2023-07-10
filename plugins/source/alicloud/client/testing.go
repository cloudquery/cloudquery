package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct{}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, _ TestOptions) {
	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	spec := Spec{
		Accounts: []AccountSpec{
			{
				Name:      "test-account",
				Regions:   []string{"cn-hangzhou"},
				AccessKey: "test-access-key",
				SecretKey: "test-secret-key",
			},
		},
		BillHistoryMonths: 0,
		Concurrency:       0,
		DeterministicCQID: false,
	}
	schedulerClient, err := New(l, spec)
	if err != nil {
		t.Fatal(err)
	}
	schedulerClient.(*Client).updateServices(builder(t, ctrl))
	tables := schema.Tables{table}
	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}
	sc := scheduler.NewScheduler(scheduler.WithLogger(l))
	messages, err := sc.SyncAll(context.Background(), schedulerClient, tables)
	if err != nil {
		t.Fatal(err)
	}
	inserts := messages.GetInserts()
	records := inserts.GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}
