package client

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

func filterInserts(msgs message.SyncMessages) message.SyncInserts {
	inserts := []*message.SyncInsert{}
	for _, msg := range msgs {
		if m, ok := msg.(*message.SyncInsert); ok {
			inserts = append(inserts, m)
		}
	}
	return inserts
}

func MockTestHelper(t *testing.T, table *schema.Table, createService func(*httprouter.Router) error) {
	t.Helper()
	table.IgnoreInTests = false
	mux := httprouter.New()
	ts := httptest.NewUnstartedServer(mux)
	defer ts.Close()
	if err := createService(mux); err != nil {
		t.Fatalf("failed to createService: %v", err)
	}
	ts.Start()

	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))
	c, err := Configure(context.Background(), l, &Spec{
		APIKey: "test", Tailnet: "test", EndpointURL: ts.URL,
	})
	if err != nil {
		t.Fatalf("failed to configure: %v", err)
	}
	messages, err := sched.SyncAll(context.Background(), c, schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	records := filterInserts(messages).GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}
